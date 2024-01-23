// this code uses a mix of the regular google provider and the beta one
// why? this is what google does too: https://github.com/terraform-google-modules/terraform-google-lb-http/blob/v10.1.0/main.tf
// going without beta seems to be a problem with Certificate Manager certs

resource "google_compute_url_map" "default" {
  provider = google-beta
  project  = var.project

  name = "${var.lb_name_prefix}-urlmap"

  default_service = var.default_service

  dynamic "host_rule" {
    for_each = var.backend_services
    content {
      hosts        = host_rule.value.domains
      path_matcher = "pathmatcher${host_rule.key}"
    }
  }

  dynamic "path_matcher" {
    for_each = var.backend_services
    content {
      name            = "pathmatcher${path_matcher.key}"
      default_service = path_matcher.value.id
    }
  }
}

resource "google_compute_target_https_proxy" "default" {
  name = "${var.lb_name_prefix}-https-proxy"

  url_map = google_compute_url_map.default.id

  certificate_map = "//certificatemanager.googleapis.com/${google_certificate_manager_certificate_map.certificate_map.id}"
}

resource "google_certificate_manager_certificate_map" "certificate_map" {
  name = "${var.lb_name_prefix}-cert-map"
  labels = {
    "terraform" : true,
    "acc-test" : true,
  }
}

resource "google_certificate_manager_certificate_map_entry" "default" {
  name         = "${var.lb_name_prefix}-cert-map-entry"
  map          = google_certificate_manager_certificate_map.certificate_map.name
  certificates = [var.certificate_name]
  matcher      = "PRIMARY"
}

resource "google_compute_global_forwarding_rule" "default" {
  provider = google-beta
  project  = var.project

  name = "${var.lb_name_prefix}-lb"

  target     = google_compute_target_https_proxy.default.id
  port_range = "443"
  ip_address = var.global_ip_address
}

// HTTP to HTTPS redirect
resource "google_compute_url_map" "https_redirect" {
  provider = google-beta
  project  = var.project

  name = "${var.lb_name_prefix}-https-redirect"

  default_url_redirect {
    https_redirect         = true
    redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
    strip_query            = false
  }
}

resource "google_compute_target_http_proxy" "https_redirect" {
  name    = "${var.lb_name_prefix}-http-proxy"
  url_map = google_compute_url_map.https_redirect.id
}

resource "google_compute_global_forwarding_rule" "https_redirect" {
  provider = google-beta
  project  = var.project

  name = "${var.lb_name_prefix}-lb-http"

  target     = google_compute_target_http_proxy.https_redirect.id
  port_range = "80"
  ip_address = var.global_ip_address
}

output "loadbalancer_ip" {
  value = google_compute_global_forwarding_rule.default.ip_address
}
