# Configure GitHub OIDC in Google Cloud to allow keyless access to GCP from GHA.

resource "google_service_account" "gha" {
  project    = var.project
  account_id = "verifa-website-gha"
}

resource "google_project_iam_member" "gha" {
  project = var.project
  role    = "roles/run.admin"
  member  = "serviceAccount:${google_service_account.gha.email}"
}

resource "google_project_iam_member" "gha_artifactregistry" {
  project = var.project
  role    = "roles/artifactregistry.writer"
  member  = "serviceAccount:${google_service_account.gha.email}"
}

module "gh_oidc" {
  source      = "terraform-google-modules/github-actions-runners/google//modules/gh-oidc"
  version     = "3.1.2"
  project_id  = var.project
  pool_id     = "verifa-website-pool"
  provider_id = "verifa-website-gh-provider"
  sa_mapping = {
    (google_service_account.gha.account_id) = {
      sa_name   = google_service_account.gha.name
      attribute = "attribute.repository/verifa/website"
    }
  }
}
