variable "service_name_prefix" {
  description = "Unique prefix for the service"
  type        = string
  default     = "service"
}

variable "region" {
  description = "Region for the service"
  type        = string
  default     = "europe-north1"
}

variable "env" {
  description = "Environment variables to pass to the service"
  type        = list(map(string))
  default     = [{}]
}

variable "min_scale" {
  description = "Minimum number of instances to run"
  type        = number
  default     = 0

}

variable "max_scale" {
  description = "Maximum number of instances to run"
  type        = number
  default     = 10
}
