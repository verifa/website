variable "lb_name_prefix" {
  description = "Unique prefix for the load balancer"
  type        = string
  default     = "lb-default"
}

variable "backend_services" {
  description = "List of backend services to use"
  type = list(object(
    {
      id      = string
      domains = list(string)
    }
  ))
}

variable "certificate_name" {
  description = "Name of the certificate to use"
  type        = string
}

variable "global_ip_address" {
  description = "Global IP address to use"
  type        = string
}

variable "default_service" {
  description = "Default service to use when host does not match any rules"
  type        = string
}

variable "project" {
  type    = string
  default = "verifa-website"
}
