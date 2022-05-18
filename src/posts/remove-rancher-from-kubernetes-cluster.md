---
type: Blog
title: How to remove Rancher from a Kubernetes cluster
subheading: Rancher is a platform for managing Kubernetes clusters and workloads. This short post covers a recent experience trying to remove Rancher from a cluster. This approach can be used to either remove Rancher itself, or for cleaning up a cluster that was imported into Rancher.
authors:
- jlarfors
tags:
- Kubernetes
- Rancher
date: 2022-05-19
image: "/blogs/remove-rancher-from-kubernetes-cluster.png"
featured: false

---

Rancher is a platform for managing Kubernetes clusters and workloads. This short post covers a recent experience trying to remove Rancher from a cluster and will hopefully help others in a similar situation. This approach can be used to either remove Rancher itself, or for cleaning up a cluster that was imported into Rancher.

## Removing Rancher (according to docs)

There’s a section under the FAQ part of the Rancher docs that provides a very short version of [how to remove Rancher](https://rancher.com/docs/rancher/v2.6/en/faq/removing-rancher/#what-if-i-don-t-want-rancher-anymore).

It’s worth noting that simply deleting the Helm release (or Kubernetes manifests) is not sufficient for removing Rancher, as Rancher creates a whole bunch of Custom Resource Definitions (CRDs), instances of those CRDs and multiple namespaces. These DO NOT get removed when deleting the Rancher Helm release.

So, there’s a [System Tools](https://github.com/rancher/system-tools) available on GitHub which is supposed to clean up the leftovers from Rancher. However, at the time of writing, the CI pipeline on `master` branch is broken and the last commit was on 15th April 2019... So my confidence in this tool was not off to a great start.

The actual command you want to run is:

```bash
# NOTE: I would not allow system-tools to point to your user-level kubeconfig.
# I would create a specific kubeconfig containing only the target cluster configs
# because I want to be sure that I don't start deleting stuff in the wrong cluster
system-tools remove --kubeconfig <cluster-specific-kubeconfig>
```

This might produce a bunch of warnings like this:

```bash
WARN[0019] Can't build dynamic client for [apiservices]: the server could not find the requested resource
```

Apparently that’s all ok because I’ve had this utility work and successfully clean my cluster. However, this command has failed, and has left a whole bunch of resources in a `Terminating` state because of finalizers that referred to resource types that no longer existed...

## Cleaning up after system-tools

If the `system-tools remove` sub-command completed but you have a bunch of namespaces and resources in the `Terminating` status, you might get a bit disappointed and now want to go through each of these manually to find that they are not being deleted because of [finalizers](https://kubernetes.io/docs/concepts/overview/working-with-objects/finalizers/). In my case, the finalizers that they referred to were Custom Resource Definitions that either never existed, or were somehow deleted and no longer existed (I never found out the truth...). Regardless, the finalizer will prevent the resources from being deleted. So let’s find out the left over resource types and instances of those, and then how to remove the finalizers that are blocking the delete from happening.

### Listing resource types

To list any resource types left over by Rancher, I  ran:

```bash
# List cluster-level resource types
kubectl api-resources --verbs=list -o name --namespaced=false | grep "cattle.io"

# List namespaced resource types
kubectl api-resources --verbs=list -o name --namespaced=true | grep "cattle.io"
```

I found that all resource types created by Rancher had the suffix `cattle.io`.

### Listing resources

To list resource instances left over by Rancher, I ran:

```bash
# List cluster-level resources
kubectl api-resources --verbs=list -o name --namespaced=false | grep "cattle.io" | xargs -n 1 kubectl get --show-kind

# List namespaced resources
kubectl api-resources --verbs=list -o name --namespaced=true | grep "cattle.io" | xargs -n 1 kubectl get --show-kind --all-namespaces
```

### Removing finalizers

In my case, there were only two unique subdomains of finalizers that were causing issues:

1. `<possible-prefix.controller.cattle.io/<some-resource>` 
2. `<possible-prefix>.wrangler.cattle.io/<some-resource>`

Neither `controller.cattle.io` or `wrangler.cattle.io` existed as resource types (nor any sub-types), and these finalizers were attached to many resources in multiple namespaces, and also some cluster-level resources.

To remove these, I wrote a simple bash script (I did not find a clean way of making a `kubectl` one-liner with the fully-resolved resource type name... And I did not particularly enjoy doing this, so I made a bash script).

Note: I have added the `--dry-run=client` argument to the `kubectl patch` command just in case someone copy+pastes this to avoid any harm. To actually do something, remove the `--dry-run` argument altogether.

```bash
# Get the resource types for both namespace and cluster-level
NS_TYPES=$(kubectl api-resources --verbs=list -o name --namespaced=true | grep "cattle.io")
CLUSTER_TYPES=$(kubectl api-resources --verbs=list -o name --namespaced=false | grep "cattle.io")

echo "Removing finalizers from namespaced Rancher resources"
for type in $NS_TYPES; do
    echo "Removing finalizers for $type"
    kubectl get $type --all-namespaces -o custom-columns='NAMESPACE:.metadata.namespace','NAME:.metadata.name' --no-headers | awk '{print $1 " " $2}' | xargs -L1 bash -c "kubectl patch --dry-run=client -n \$0 $type/\$1 --type=merge -p \$(kubectl get -n \$0 $type/\$1 -o json | jq -Mcr '.metadata.finalizers // [] | {metadata:{finalizers:map(select(. | (contains(\"controller.cattle.io/\") or contains(\"wrangler.cattle.io/\")) | not ))}}')"
done

echo "Removing finalizers from cluster Rancher resources"
for type in $CLUSTER_TYPES; do
    echo "Removing finalizers for $type"
    kubectl get $type -o name --show-kind --no-headers | awk '{print $1 }' | xargs -L1 bash -c "kubectl patch --dry-run=client \$0 --type=merge -p \$(kubectl get \$0 -o json | jq -Mcr '.metadata.finalizers // [] | {metadata:{finalizers:map(select(. | (contains(\"controller.cattle.io/\") or contains(\"wrangler.cattle.io/\")) | not ))}}')"
done
```

I also had some of these finalizers on namespaces, but I did not find a clean way to automatically identify which namespaces were created by Rancher.

```bash
# Some example Rancher namespaces. Replace with your own.
NAMESPACES="local c-xxx p-yyy"
echo "Removing finalizers from namespaces"
for ns in $NAMESPACES; do
    echo "Removing finalizers from namespace $ns"
    kubectl patch namespace/\$0 --type=merge -p \$(kubectl get namespace/\$0 -o json | jq -Mcr '.metadata.finalizers // [] | {metadata:{finalizers:map(select(. | (contains(\"controller.cattle.io/\") or contains(\"wrangler.cattle.io/\")) | not ))}}')
done
```

Once these finalizers are removed, the Kubernetes control plane was able to successfully finish the cleanup and remove any trace that Rancher had ever existed. So I guess the `system-tools` kinda did work, just with a gentle nudge.

#### Only removing the faulty finalizers

As a general practice, removing finalizers is a bad idea. They are often there for a reason, but when they refer to resource types that no longer exist, I think we can remove them. However, we don’t want to remove ALL the finalizers from a resource; only those which we can consider faulty.

If you want to test the `jq` query here is the nested `kubectl` command that pipes the JSON output to `jq`.

```bash
kubectl get <example-resource> -o json | jq -Mcr '.metadata.finalizers // [] | {metadata:{finalizers:map(select(. | (contains("controller.cattle.io/") or contains("wrangler.cattle.io/")) | not ))}}'
```

Check that this only removes the finalizers you want to remove.

## Summary

Rancher can be a pain to remove because it creates so many CRDs, resources and namespaces in your cluster *at runtime*. This means your deployment tool has no record of those resources being created and cannot clean them up for you. Hopefully for you, the `system-tools remove` subcommand will work and purge your cluster of Rancher, but if not, I hope this post has given you some tips for removing Rancher. Personally, I spent too much time doing this and that’s why I decided to write this post :)

If you have any questions, feedback or want help with managing Kubernetes please leave us a comment or [get in touch here!](https://verifa.io/contact/)