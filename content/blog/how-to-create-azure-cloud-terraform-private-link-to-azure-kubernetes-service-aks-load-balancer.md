---
type: Blog
title: How to create an Azure Private Link to a Load Balancer in AKS with Terraform
subheading: In this blog we'll look at how to build an Azure Kubernetes Service (AKS)
  cluster with a private Load Balancer.
authors:
- alarfors
tags:
- Azure
- Terraform
- Kubernetes
date: 2021-11-02
image: "/blog/2021-11-03/azure-cloud-partners-01.svg"
featured: true
---
**Infrastructure-as-Code (IaC) is a great way of managing your infrastructure, increasing maintainability and traceability, among other things. Terraform is a very popular IaC tool, in part because it supports so many different platforms and tools with its large community of Providers. If you're working with Azure Cloud, for example, you will probably use the** [**Azure Cloud Terraform Provider**](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs) **(azurerm - for it uses the Azure Resource Manager API).**

You may need to use more than one provider in order to deploy your entire system. In this blog, we are creating an Azure Kubernetes Service (AKS) Cluster and then deploying a series of Kubernetes Resources in it. The aforementioned azurerm provider can manage Azure Cloud Resources, but its ability to deploy resources into our cluster is limited. Therefore, we are also using the [Kubernetes Terraform Provider](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs).

In our use case, we are building an AKS cluster with security requirements, so it should not have any Public IP addresses. This means the Load Balancers we create will have internal IP addresses and should only be accessible from certain Virtual Networks in our Azure Cloud subscription. We are achieving this connection via Azure Private Link Service, which allows a Private Endpoint to route traffic to a Load Balancer.

## The problem

So, to the crux of our problem: our Load Balancers are managed by the Kubernetes provider and once they are created we need to use the azurerm Terraform provider to create the Private Link Services. In order to create a Private Link Service to a Load Balancer, we need to know the ID of the Frontend IP Configuration. This is where things get a little bit tricky because the Azure LoadBalancer and Frontend IP Configurations are created by our Kubernetes cluster when we deploy our Kubernetes Service of type LoadBalancer. The ordering therefore goes:

1. \[azurerm\] Create the AKS cluster
2. \[kubernetes\] Deploy Kubernetes Service of type LoadBalancer
3. \[azurerm\] Fetch the ID of the Azure Load Balancer's Front IP Configuration created indirectly in step 2

**Side note:** an alternative way of creating this setup would be to create the Azure Load Balancers in Azure directly (instead of through a Kubernetes service). We could then rely on our single azurerm provider and if you want to use that method, you can stop reading this blog here. The downside of this approach is that we must also create the backend configuration, routing rules, frontend IP configuration etc. that the Azure Kubernetes Service so nicely takes care of for us if we create the Load Balancer in Kubernetes. When we declare our Load Balancer in AKS, we provide the Selector and Ports configurations and AKS takes care of the rest of the required setup.

\[As a final note on this topic, the application being deployed here is a COTS application which provides the Kubernetes service of type LoadBalancer as part of the application package, so it is being used as part of retaining the application support provided by the retailer.\]

**Back to our problem:** Using the Kubernetes Terraform provider to deploy a service of type LoadBalancer means that the [attribute reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service#attributes-1) will be missing Azure-specific resource IDs (and other Azure-specific values) that you might need when creating related objects, such as Private Links. An Azure Cloud Private Link requires a frontend IP configuration ID, for example, and the Kubernetes provider resource "kubernetes_service" exports no such attribute.

## The solution

We will fetch the load balancer as an azurerm provider "azurerm_lb" data object from which we can get the frontend ip configuration id.

### Example

Given that we are setting up two environments, "prod" and "test", in a single cluster, let's create a load balancer for each ...

```hcl
variable "environments" {
  description = "List of environments"
  type = list
  default = ["prod", "test"]
}

...

# Create a loadbalancer for each environment
resource "kubernetes_service" "loadbalancer" {
  for_each = toset( var.environments )
  metadata {
    name        = "${each.key}-loadbalancer"
    ...
  }
  spec {
    selector = {
      ...
    }
    port {
      ...
    }
    type = "LoadBalancer"
  }
}
```

Now we need to create a Private Link Service to each Load Balancer. This is where the magic happens (and it's pretty ugly magic at that):

```hcl
resource "azurerm_private_link_service" "example" {
  for_each = kubernetes_service.loadbalancer

  name                = "pl-${each.value.metadata.0.name}"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location


  load_balancer_frontend_ip_configuration_ids = [
    data.azurerm_lb.example.frontend_ip_configuration[
      index(        data.azurerm_lb.example.frontend_ip_configuration.*.private_ip_address,
        each.value.status.0.load_balancer.0.ingress.0.ip
      )
    ].id
  ]
  ...
}
```

Let's break down what's going on here. We have two representations of our load balancers:

1. `kubernetes_service.loadbalancer`- is the Kubernetes Provider created resource (A Kubernetes Service resource of kind LoadBalancer).
2. `data.azurerm_lb.example` - is a data object we have fetched after the Kubernetes Provider has created the load balancers (a standard Azure Cloud Load Balancer).

Creating our `azurerm_private_link_service.example` resource requires the `load_balancer_frontend_ip_configuration_ids` argument. Only `data.azurerm_lb.example` has this value, and we are iterating over `kubernetes_service.loadbalancer`. So we need to match the frontend ip configuration in the `data.azurerm_lb.example` to the current `kubernetes_service.loadbalancer`.

We achieve this using the index( ... ) function in Terraform:

```hcl
index(
  data.azurerm_lb.example.frontend_ip_configuration.*.private_ip_address,
  each.value.status.0.load_balancer.0.ingress.0.ip
)
```

The call to index( ... ) above will return the index of the item which matches the condition. The condition is that the `private_ip_address` of the `data.azurerm_lb.example.frontend_ip_configuration` matches the ip of the current `kubernetes_service.loadbalancer` being iterated over. Note how we must get the latter value from the `status` attribute.

Once we have our index value, we fetch the correct map and retrieve the "id" field from there:

```hcl
load_balancer_frontend_ip_configuration_ids = [
    data.azurerm_lb.example.frontend_ip_configuration[
      index( ... )
    ].id
  ]
```

A bit of a messy solution to a simple problem, but such is sometimes the nature of working with multiple providers in Terraform.

### Full example

```hcl
variables.tf

variable "environments" {

description = "List of environments"

type = list

default = ["prod", "test"]

}

kubernetes.tf

# Create the cluster
resource "azurerm_kubernetes_cluster" "example" {
  name                = "aks-example"
  ...
}

# Configure the kubernetes provider
provider "kubernetes" {
  host = azurerm_kubernetes_cluster.example.kube_config.0.host

  client_certificate     = base64decode( azurerm_kubernetes_cluster.example.kube_config.0.client_certificate)
  client_key             = base64decode( azurerm_kubernetes_cluster.example.kube_config.0.client_key)
  cluster_ca_certificate = base64decode( azurerm_kubernetes_cluster.example.kube_config.0.cluster_ca_certificate)
}


# Create a loadbalancer for each environment
resource "kubernetes_service" "loadbalancer" {
  for_each = toset( var.environments )
  metadata {
    name        = "${each.key}-loadbalancer"
    ...
  }
  spec {
    selector = {
      ...
    }
    port {
      ...
    }
    type = "LoadBalancer"
  }
}

# Get the Kubernetes LB as an azurerm object
data "azurerm_lb" "example" {
  name                = "kubernetes"
  resource_group_name = azurerm_kubernetes_cluster.example.node_resource_group
}


# Create a private link service for each environment
resource "azurerm_private_link_service" "example" {
  for_each = kubernetes_service.loadbalancer

  name                = "pl-${each.value.metadata.0.name}"
  ...


  load_balancer_frontend_ip_configuration_ids = [
    data.azurerm_lb.example.frontend_ip_configuration[
      index(        data.azurerm_lb.example.frontend_ip_configuration.*.private_ip_address,
        each.value.status.0.load_balancer.0.ingress.0.ip
      )
    ].id
  ]

  nat_ip_configuration {
    ...
  }
}
```

## Useful Links

[Azurerm Terraform provider](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs)

[Kubernetes Terraform provider](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs)
