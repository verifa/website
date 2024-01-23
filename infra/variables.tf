// defaults are for production usage
variable "project" {
  type    = string
  default = "verifa-website"
}

variable "region" {
  type    = string
  default = "europe-north1"
}

variable "zone" {
  type    = string
  default = "europe-north1-a"
}

variable "stack_name" {
  type    = string
  default = "website-prod"
}

variable "certmanager_certificate_name" {
  type    = string
  default = "projects/verifa-website/locations/global/certificates/verifa-website"
}
