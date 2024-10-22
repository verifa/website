package infra

import (
	"fmt"

	"github.com/golingon/lingon/pkg/terra"
	"github.com/golingon/terraproviders/google/5.25.0/google_cloud_run_service"
	"github.com/golingon/terraproviders/google/5.25.0/google_cloud_run_service_iam_member"
	"github.com/golingon/terraproviders/google/5.25.0/google_compute_backend_bucket"
	"github.com/golingon/terraproviders/google/5.25.0/google_compute_backend_service"
	"github.com/golingon/terraproviders/google/5.25.0/google_compute_region_network_endpoint_group"
	"github.com/golingon/terraproviders/google/5.25.0/google_storage_bucket"
	"github.com/golingon/terraproviders/google/5.25.0/google_storage_bucket_iam_member"
	"github.com/golingon/terraproviders/google/5.25.0/google_storage_bucket_object"

	_ "embed"
)

//go:embed 404.html
var notFoundHTML string

var S = terra.String

type Service struct {
	BackendService       *google_compute_backend_service.Resource
	CloudRunService      *google_cloud_run_service.Resource
	IAMMember            *google_cloud_run_service_iam_member.Resource
	NetworkEndpointGroup *google_compute_region_network_endpoint_group.Resource
}

type ServiceOpts struct {
	Name     string
	MinScale int
	MaxScale int
}

func NewService(opts ServiceOpts) Service {
	cloudRun := google_cloud_run_service.Resource{
		Name: "service",
		Args: google_cloud_run_service.Args{
			Name:     S(opts.Name + "-service"),
			Location: S("europe-north1"),

			Template: &google_cloud_run_service.Template{
				// Spec and container is required when creating a CloudRun
				// service. Use a placeholder image and ignore changes to it in
				// Terraform. The "real" image will be set dynamically outside
				// of Terraform.
				Spec: &google_cloud_run_service.TemplateSpec{
					Containers: []google_cloud_run_service.TemplateSpecContainers{
						{
							Image: S(
								"us-docker.pkg.dev/cloudrun/container/placeholder",
							),
						},
					},
				},
				Metadata: &google_cloud_run_service.TemplateMetadata{
					Annotations: terra.MapString(map[string]string{
						"run.googleapis.com/startup-cpu-boost": "true",
						"autoscaling.knative.dev/minScale": fmt.Sprintf(
							"%d",
							opts.MinScale,
						),
						"autoscaling.knative.dev/maxScale": fmt.Sprintf(
							"%d",
							opts.MaxScale,
						),
					}),
				},
			},
		},
	}
	attrTemplate := cloudRun.Attributes().Template().Index(0)
	annotationsKey := attrTemplate.Metadata().
		Index(0).
		Annotations().Key
	cloudRun.Lifecycle = &terra.Lifecycle{
		IgnoreChanges: terra.IgnoreChanges(
			// Ignore changes to placeholder container.
			attrTemplate.Spec().Index(0).Containers().Index(0),
			// Ignore annotations that are added dynamically outside of
			// Terraform.
			annotationsKey("run.googleapis.com/client-name"),
			annotationsKey("run.googleapis.com/client-version"),
			annotationsKey("client.knative.dev/user-image"),
		),
	}

	// Allow unauthenticated access to service.
	iamMember := google_cloud_run_service_iam_member.Resource{
		Name: "member",
		Args: google_cloud_run_service_iam_member.Args{
			Location: cloudRun.Attributes().Location(),
			Project:  cloudRun.Attributes().Project(),
			Service:  cloudRun.Attributes().Name(),
			Role:     S("roles/run.invoker"),
			Member:   S("allUsers"),
		},
	}

	networkEndpointGroup := google_compute_region_network_endpoint_group.Resource{
		Name: "cloudrun_neg",
		Args: google_compute_region_network_endpoint_group.Args{
			Name:                S(opts.Name + "-neg"),
			NetworkEndpointType: S("SERVERLESS"),
			Region:              S("europe-north1"),
			CloudRun: &google_compute_region_network_endpoint_group.CloudRun{
				Service: cloudRun.Attributes().Name(),
			},
		},
	}

	backendService := google_compute_backend_service.Resource{
		Name: "default",
		Args: google_compute_backend_service.Args{
			Name:       S(opts.Name + "-backend"),
			Protocol:   S("HTTP"),
			PortName:   S("http"),
			TimeoutSec: terra.Number(30),
			EnableCdn:  terra.Bool(true),
			CdnPolicy: &google_compute_backend_service.CdnPolicy{
				CacheMode:               S("CACHE_ALL_STATIC"),
				ClientTtl:               terra.Number(3600),
				DefaultTtl:              terra.Number(3600),
				MaxTtl:                  terra.Number(3600 * 24),
				SignedUrlCacheMaxAgeSec: terra.Number(7200),
			},
			Backend: []google_compute_backend_service.Backend{
				{
					Group: networkEndpointGroup.Attributes().Id(),
				},
			},
		},
	}

	return Service{
		BackendService:       &backendService,
		CloudRunService:      &cloudRun,
		IAMMember:            &iamMember,
		NetworkEndpointGroup: &networkEndpointGroup,
	}
}

type DefaultBackend struct {
	Bucket       *google_storage_bucket.Resource
	NotFoundPage *google_storage_bucket_object.Resource
	Backend      *google_compute_backend_bucket.Resource
	IAMMember    *google_storage_bucket_iam_member.Resource
}

func NewDefaultBackend() DefaultBackend {
	bucket := google_storage_bucket.Resource{
		Name: "default",
		Args: google_storage_bucket.Args{
			Name:                     terra.String("verifa-website-404"),
			StorageClass:             terra.String("STANDARD"),
			UniformBucketLevelAccess: terra.Bool(false),
			Location:                 terra.String("europe-north1"),
			Website: &google_storage_bucket.Website{
				MainPageSuffix: terra.String("404.html"),
				NotFoundPage:   terra.String("404.html"),
			},
		},
	}
	notFoundPage := google_storage_bucket_object.Resource{
		Name: "not_found_page",
		Args: google_storage_bucket_object.Args{
			Bucket:      bucket.Attributes().Id(),
			Name:        terra.String("404.html"),
			Content:     terra.String(notFoundHTML),
			ContentType: terra.String("text/html"),
		},
	}
	backend := google_compute_backend_bucket.Resource{
		Name: "default",
		Args: google_compute_backend_bucket.Args{
			Name:        terra.String("verifa-website-default"),
			Description: terra.String("Default backend service for verifa.io"),
			BucketName:  bucket.Attributes().Name(),
			EnableCdn:   terra.Bool(false),
		},
	}
	iamMember := google_storage_bucket_iam_member.Resource{
		Name: "default",
		Args: google_storage_bucket_iam_member.Args{
			Bucket: bucket.Attributes().Name(),
			Role:   terra.String("roles/storage.objectViewer"),
			Member: terra.String("allUsers"),
		},
	}
	return DefaultBackend{
		Bucket:       &bucket,
		NotFoundPage: &notFoundPage,
		Backend:      &backend,
		IAMMember:    &iamMember,
	}
}
