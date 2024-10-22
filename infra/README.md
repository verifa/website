# Website Deployment

Website infra is built on Cloud Run with a global load balancer in-front and Cloud CDN enabled.
This directory contains all the OpenTofu code to configure the production and pre-production infrastructure.

The actual deployment of the site is decoupled from the infra for convenience, it does not matter for the infrastructure what is exactly running inside the Cloud Run service(s).
Changes to the image of the deployed Cloud Run services are ignored, could add more fields in future if needed.

## Usage

Login to GCP (fetches Application Default Credentials for use by Tofu):

```bash
gcloud auth application-default login
```

Then your usual OpenTofu usage:

```bash
tofu plan
tofu apply
```
