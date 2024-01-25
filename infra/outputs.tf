output "ip_address" {
  value = google_compute_global_address.this.address
}

output "cloud_run_staging" {
  value = module.staging-service.service_name
}

output "cloud_run_prod" {
  value = module.prod-service.service_name
}
