package infra

import (
	"github.com/golingon/lingon/pkg/terra"
	"github.com/golingon/terraproviders/google/5.25.0/google_iam_workload_identity_pool"
	"github.com/golingon/terraproviders/google/5.25.0/google_iam_workload_identity_pool_provider"
	"github.com/golingon/terraproviders/google/5.25.0/google_project_iam_member"
	"github.com/golingon/terraproviders/google/5.25.0/google_service_account"
	"github.com/golingon/terraproviders/google/5.25.0/google_service_account_iam_member"
)

type GitHubOIDC struct {
	ProjectIAMMember
	ServiceAccount *google_service_account.Resource

	WorkloadIDPool         *google_iam_workload_identity_pool.Resource
	WorkloadIDPoolProvider *google_iam_workload_identity_pool_provider.Resource
	WorkloadIDMember       *google_service_account_iam_member.Resource
}

type GitHubOIDCOpts struct {
	Project string
}

func NewGitHubOIDC(opts GitHubOIDCOpts) GitHubOIDC {
	serviceAccount := google_service_account.Resource{
		Name: "gha",
		Args: google_service_account.Args{
			Project:   terra.String(opts.Project),
			AccountId: terra.String("verifa-website-gha"),
		},
	}
	iamMember := NewProjectIAMMember(
		opts.Project,
		serviceAccount.Attributes().Email(),
	)

	workloadIDPool := google_iam_workload_identity_pool.Resource{
		Name: "main",
		Args: google_iam_workload_identity_pool.Args{
			Project:                terra.String(opts.Project),
			WorkloadIdentityPoolId: terra.String("verifa-website-pool"),
			DisplayName:            terra.String(""),
			Description: terra.String(
				"Workload Identity Pool managed by Terraform",
			),
			Disabled: terra.Bool(false),
		},
	}
	workloadIDPoolProvider := google_iam_workload_identity_pool_provider.Resource{
		Name: "main",
		Args: google_iam_workload_identity_pool_provider.Args{
			Project: terra.String(opts.Project),
			WorkloadIdentityPoolId: workloadIDPool.Attributes().
				WorkloadIdentityPoolId(),
			WorkloadIdentityPoolProviderId: terra.String(
				"verifa-website-gh-provider",
			),
			DisplayName: terra.String(""),
			Description: terra.String(
				"Workload Identity Pool Provider managed by Terraform",
			),
			AttributeMapping: terra.MapString(map[string]string{
				"google.subject":       "assertion.sub",
				"attribute.actor":      "assertion.actor",
				"attribute.aud":        "assertion.aud",
				"attribute.repository": "assertion.repository",
			}),
			Oidc: &google_iam_workload_identity_pool_provider.Oidc{
				IssuerUri: terra.String(
					"https://token.actions.githubusercontent.com",
				),
			},
		},
	}
	workloadIDMember := google_service_account_iam_member.Resource{
		Name: "main",
		Args: google_service_account_iam_member.Args{
			ServiceAccountId: serviceAccount.Attributes().Id(),
			Role:             terra.String("roles/iam.workloadIdentityUser"),
			Member: terra.StringFormat(
				"principalSet://iam.googleapis.com/${%s}/attribute.repository/verifa/website",
				workloadIDPool.Attributes().Name(),
			),
		},
	}

	return GitHubOIDC{
		ProjectIAMMember:       iamMember,
		ServiceAccount:         &serviceAccount,
		WorkloadIDPool:         &workloadIDPool,
		WorkloadIDPoolProvider: &workloadIDPoolProvider,
		WorkloadIDMember:       &workloadIDMember,
	}
}

func NewProjectIAMMember(
	project string,
	serviceAccountEmail terra.StringValue,
) ProjectIAMMember {
	member := terra.StringFormat(
		"serviceAccount:${%s}",
		serviceAccountEmail,
	)
	return ProjectIAMMember{
		CloudRunAdmin: &google_project_iam_member.Resource{
			Name: "gha",
			Args: google_project_iam_member.Args{
				Project: terra.String(project),
				Role:    terra.String("roles/run.admin"),
				Member:  member,
			},
		},
		LBAdmin: &google_project_iam_member.Resource{
			Name: "gha_lb_admin",
			Args: google_project_iam_member.Args{
				Project: terra.String(project),
				Role:    terra.String("roles/compute.loadBalancerAdmin"),
				Member:  member,
			},
		},
		StorageObjectUser: &google_project_iam_member.Resource{
			Name: "gha_storage_object_user",
			Args: google_project_iam_member.Args{
				Project: terra.String(project),
				Role:    terra.String("roles/storage.objectUser"),
				Member:  member,
			},
		},
		IAMWorkloadIdentityUser: &google_project_iam_member.Resource{
			Name: "gha_iam_workload_identity_user",
			Args: google_project_iam_member.Args{
				Project: terra.String(project),
				Role:    terra.String("roles/iam.workloadIdentityUser"),
				Member:  member,
			},
		},
		IAMServiceAccountUser: &google_project_iam_member.Resource{
			Name: "gha_iam_service_account_user",
			Args: google_project_iam_member.Args{
				Project: terra.String(project),
				Role:    terra.String("roles/iam.serviceAccountUser"),
				Member:  member,
			},
		},
		ArtifactRegistryWriter: &google_project_iam_member.Resource{
			Name: "gha_artifactregistry",
			Args: google_project_iam_member.Args{
				Project: terra.String(project),
				Role:    terra.String("roles/artifactregistry.writer"),
				Member:  member,
			},
		},
	}
}

type ProjectIAMMember struct {
	CloudRunAdmin           *google_project_iam_member.Resource
	LBAdmin                 *google_project_iam_member.Resource
	StorageObjectUser       *google_project_iam_member.Resource
	IAMWorkloadIdentityUser *google_project_iam_member.Resource
	IAMServiceAccountUser   *google_project_iam_member.Resource
	ArtifactRegistryWriter  *google_project_iam_member.Resource
}
