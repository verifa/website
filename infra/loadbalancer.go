package infra

import (
	"github.com/golingon/lingon/pkg/terra"
	_ "github.com/golingon/lingon/pkg/terragen"
	"github.com/golingon/terraproviders/google/5.25.0/google_certificate_manager_certificate_map"
	"github.com/golingon/terraproviders/google/5.25.0/google_certificate_manager_certificate_map_entry"
	"github.com/golingon/terraproviders/google/5.25.0/google_compute_global_address"
	"github.com/golingon/terraproviders/google/5.25.0/google_compute_global_forwarding_rule"
	"github.com/golingon/terraproviders/google/5.25.0/google_compute_target_http_proxy"
	"github.com/golingon/terraproviders/google/5.25.0/google_compute_target_https_proxy"
	"github.com/golingon/terraproviders/google/5.25.0/google_compute_url_map"
)

type LoadBalancer struct {
	GlobalAddress        *google_compute_global_address.Resource
	URLMap               *google_compute_url_map.Resource
	TargetHTTPSProxy     *google_compute_target_https_proxy.Resource
	CertManagerMap       *google_certificate_manager_certificate_map.Resource
	CertManagerMapEntry  *google_certificate_manager_certificate_map_entry.Resource
	GlobalForwardingRule *google_compute_global_forwarding_rule.Resource

	HTTPSRedirect
}

type LoadBalancerOpts struct {
	// Project         string
	DefaultService string
	// BackendServices []BackendService
}

type BackendService struct {
	ID      string
	Domains []string
}

func NewLoadBalancer(opts LoadBalancerOpts) LoadBalancer {
	globalAddress := google_compute_global_address.Resource{
		Name: "this",
		Args: google_compute_global_address.Args{
			Name: terra.String("website-prod-address"),
		},
	}
	// hostRules := make(
	// 	[]google_compute_url_map.HostRule,
	// 	len(opts.BackendServices),
	// )
	// pathMatchers := make(
	// 	[]google_compute_url_map.PathMatcher,
	// 	len(opts.BackendServices),
	// )
	// for i, backendService := range opts.BackendServices {
	// 	hostRules[i] = google_compute_url_map.HostRule{
	// 		Hosts: terra.SetString(backendService.Domains...),
	// 		PathMatcher: terra.String(
	// 			fmt.Sprintf("pathmatcher-%d", i),
	// 		),
	// 	}
	// 	pathMatchers[i] = google_compute_url_map.PathMatcher{
	// 		Name: terra.String(fmt.Sprintf("pathmatcher-%d", i)),
	// 		DefaultService: terra.String(
	// 			backendService.ID,
	// 		),
	// 	}
	// }
	urlMap := google_compute_url_map.Resource{
		Name: "default",
		Args: google_compute_url_map.Args{
			Name:           terra.String("lb-prod-urlmap"),
			DefaultService: terra.String(opts.DefaultService),
		},
	}
	urlMap.Lifecycle = &terra.Lifecycle{
		IgnoreChanges: terra.IgnoreChanges(
			urlMap.Attributes().HostRule(),
			urlMap.Attributes().PathMatcher(),
		),
	}

	certMap := google_certificate_manager_certificate_map.Resource{
		Name: "certificate_map",
		Args: google_certificate_manager_certificate_map.Args{
			Name: terra.String("lb-prod-cert-map"),
		},
	}

	httpsProxy := google_compute_target_https_proxy.Resource{
		Name: "default",
		Args: google_compute_target_https_proxy.Args{
			Name:   terra.String("lb-prod-https-proxy"),
			UrlMap: urlMap.Attributes().Id(),
			CertificateMap: terra.StringFormat(
				"//certificatemanager.googleapis.com/${%s}",
				certMap.Attributes().Id(),
			),
		},
	}

	certMapEntry := google_certificate_manager_certificate_map_entry.Resource{
		Name: "default",
		Args: google_certificate_manager_certificate_map_entry.Args{
			Name: terra.String("lb-prod-cert-map-entry"),
			Map:  certMap.Attributes().Name(),
			Certificates: terra.ListString(
				"projects/verifa-website/locations/global/certificates/verifa-website",
			),
			Matcher: terra.String("PRIMARY"),
		},
	}

	forwardingRule := google_compute_global_forwarding_rule.Resource{
		Name: "default",
		// Provider: google-beta
		Args: google_compute_global_forwarding_rule.Args{
			// Project:   terra.String(opts.Project),
			Name:      terra.String("lb-prod-lb"),
			Target:    httpsProxy.Attributes().Id(),
			PortRange: terra.String("443"),
			IpAddress: globalAddress.Attributes().Address(),
		},
	}
	return LoadBalancer{
		GlobalAddress:        &globalAddress,
		URLMap:               &urlMap,
		TargetHTTPSProxy:     &httpsProxy,
		CertManagerMap:       &certMap,
		CertManagerMapEntry:  &certMapEntry,
		GlobalForwardingRule: &forwardingRule,

		HTTPSRedirect: *NewHTTPSRedirect(opts, globalAddress.Attributes().Address()),
	}
}

type HTTPSRedirect struct {
	URLMap         *google_compute_url_map.Resource
	HTTPProxy      *google_compute_target_http_proxy.Resource
	ForwardingRule *google_compute_global_forwarding_rule.Resource
}

func NewHTTPSRedirect(
	opts LoadBalancerOpts,
	ipAddress terra.StringValue,
) *HTTPSRedirect {
	urlMap := google_compute_url_map.Resource{
		Name: "https_redirect",
		Args: google_compute_url_map.Args{
			// Project: terra.String(opts.Project),
			Name: terra.String("lb-prod-https-redirect"),
			DefaultUrlRedirect: &google_compute_url_map.DefaultUrlRedirect{
				HttpsRedirect:        terra.Bool(true),
				RedirectResponseCode: terra.String("MOVED_PERMANENTLY_DEFAULT"),
				StripQuery:           terra.Bool(false),
			},
		},
	}

	httpProxy := google_compute_target_http_proxy.Resource{
		Name: "https_redirect",
		Args: google_compute_target_http_proxy.Args{
			Name:   terra.String("lb-prod-http-proxy"),
			UrlMap: urlMap.Attributes().Id(),
		},
	}

	forwardingRule := google_compute_global_forwarding_rule.Resource{
		Name: "https_redirect",
		Args: google_compute_global_forwarding_rule.Args{
			Name:      terra.String("lb-prod-lb-http"),
			Target:    httpProxy.Attributes().Id(),
			PortRange: terra.String("80"),
			IpAddress: ipAddress,
		},
	}

	return &HTTPSRedirect{
		URLMap:         &urlMap,
		HTTPProxy:      &httpProxy,
		ForwardingRule: &forwardingRule,
	}
}
