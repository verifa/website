terraform {
  backend "gcs" {
    bucket = "verifa-website-tfstate"
    prefix = "tofu/prod"
  }
}
