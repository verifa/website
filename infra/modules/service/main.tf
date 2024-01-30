resource "google_cloud_run_service" "service" {
  name     = "${var.service_name_prefix}-service"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"

        ports {
          container_port = var.port
        }

        dynamic "env" {
          for_each = var.env
          content {
            name  = env.value.name
            value = env.value.value
          }
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/minScale" = tostring(var.min_scale)
        "autoscaling.knative.dev/maxScale" = tostring(var.max_scale)
      }
    }
  }

  lifecycle {
    ignore_changes = [
      # ignore changes to image, because we don't deploy with Tofu
      template[0].spec[0].containers["image"],
      template[0].metadata[0].annotations["client.knative.dev/user-image"],
      template[0].metadata[0].annotations["run.googleapis.com/client-name"],
      template[0].metadata[0].annotations["run.googleapis.com/client-version"],
    ]
  }
}

// allows unauthenticated access to service
resource "google_cloud_run_service_iam_member" "member" {
  location = google_cloud_run_service.service.location
  project  = google_cloud_run_service.service.project
  service  = google_cloud_run_service.service.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_compute_region_network_endpoint_group" "cloudrun_neg" {
  name                  = "${var.service_name_prefix}-neg"
  network_endpoint_type = "SERVERLESS"
  region                = var.region
  cloud_run {
    service = google_cloud_run_service.service.name
  }
}

resource "google_compute_backend_service" "default" {
  name = "${var.service_name_prefix}-backend"

  protocol    = "HTTP"
  port_name   = "http"
  timeout_sec = 30

  enable_cdn = true

  cdn_policy {
    cache_mode  = "CACHE_ALL_STATIC" // this is the default mode
    client_ttl  = 3600
    default_ttl = 3600
    max_ttl     = 3600 * 24

    signed_url_cache_max_age_sec = 7200 // mandatory

  }

  backend {
    group = google_compute_region_network_endpoint_group.cloudrun_neg.id
  }
}

