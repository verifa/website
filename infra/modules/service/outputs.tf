output "service_url" {
  value = google_cloud_run_service.service.status[0].url
}

output "network_endpoint_group_id" {
  value = google_compute_region_network_endpoint_group.cloudrun_neg.id
}

output "compute_backend_service_id" {
  value = google_compute_backend_service.default.id
}

output "service_id" {
  value = google_cloud_run_service.service.id
}

output "service_name" {
  value = google_cloud_run_service.service.name
}
