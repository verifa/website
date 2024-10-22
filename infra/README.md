# Website Deployment

Website infra is built on Cloud Run with a global load balancer in-front and Cloud CDN enabled.
This directory contains all the Lingon/OpenTofu code to configure the production, pre-production and PR preview environment infrastructure.

The actual deployment of the site is decoupled from the infra for convenience, it does not matter for the infrastructure which container is running inside the Cloud Run service(s).
Changes to the image of the deployed Cloud Run services are ignored, could add more fields in future if needed.

## Usage

```bash
# Login to GCP (fetches Application Default Credentials for use by Tofu).
gcloud auth application-default login

# Bootstrap infra, creating OIDC trust for GitHub Actions to assume a
# Google Cloud IAM service account.
go run ./cmd/bootstrap/main.go

# Deploy LB and prod and staging environments.
go run ./cmd/base/main.go

# Deploy Preview environments.
go run ./cmd/preview/main.go

```

THen run:

```bash

```
