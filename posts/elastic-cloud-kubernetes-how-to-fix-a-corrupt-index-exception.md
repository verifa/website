---
type: Blog
title: 'Elastic Cloud on Kubernetes: How to fix a corrupt index exception'
subheading: In this article we’re going to see how to fix a corrupt index which was not automatically recovered.
authors:
- avijayan
tags:
- Elastic
- Kubernetes
- Cloud
date: 2022-11-21
image: "/static/blog/elastic-cloud-kubernetes-how-to-fix-a-corrupt-index-exception.png"
featured: true

---

This article assumes the reader has good knowledge about the Elastic Stack. Also, this article tries to discuss a very specific problem happening in Elastic Stack environments.

## Introduction

To solve corruption problems, Elasticsearch database ships with a binary `elasticsearch-shard` which can remove the corrupt parts in a shard copy of a Lucene index or Translog. This tool comes in handy when the corrupt shard cannot be recovered automatically or restored from a backup. This can cause data loss. From the official documentation, *this approach should be chosen as the last resort when restoring from backups is not an option*. Also, sometimes, it’s hard to restore system indices from Elasticsearch’s snapshot backup.

In order to execute this command in an Elasticsearch node, the Elasticsearch service should not be running. This is to prevent any read/write actions while the shard is being repaired. This is complicated in a container-based environment like Kubernetes because the container exits if the entry point process is stopped, in this case, the Elasticsearch process.

However, Elastic Cloud on Kubernetes comes with a way to suspend containers with the help of Kubernetes annotations. When annotated with [`eck.k8s.elastic.co/suspend`](http://eck.k8s.elastic.co/suspend) the container restarts into one of init containers, namely `elastic-internal-suspend` with all the necessary disks mounted. After connecting to that internal container, one can have access to the disk and can run commands.

In this article we are going to see how to fix a corrupt index which was not automatically recovered.

## Event

One morning one of our Elastic Stack cluster’s health had degraded and was offline. Upon checking, it was found that, one of the indices was corrupt and was throwing the following exception: `failed shard on node [abc]: shard failure, reason [already closed by tragic event on the index writer], failure org.apache.lucene.index.CorruptIndexException: codec footer mismatch (file truncated?): actual footer=0 vs expected footer=-1071082520`. This error can also be seen from the output of `allocations API`.

## Troubleshooting

### Detach the node from the cluster

Due to the index being corrupt it should be in an `unassigned` state. We need to find the index which is corrupt and find the node(s) which was hosting the index previously from the Elasticsearch’s node logs. After finding the node, make sure to mark the cluster level shard allocation settings `cluster.routing.allocation.enable` to `primaries` instead of `all` . This informs the cluster to wait and not rearrange shards when a data node is offline. Now add the Kubernetes annotation to the data node using `kubectl`.

E.g. if your ECK cluster is named: `myeck` and the data node (pod) is `myeck-es-node-1` then perform this:

```bash
kubectl annotate es myeck eck.k8s.elastic.co/suspend=myeck-es-node-1
```

### Remove the corrupted parts and reroute the index

Before we start, remember, this can cause data loss as the tool will **remove** the corrupted parts which can have critical data. Also, try to stop all the sources from writing to the index. E.g. If a Kibana index is corrupt, then stop all instances of Kibana.

Connect into the container with `kubectl exec` and execute the command `elasticsearch-shard remove-corrupted-data` with index and shard id as parameter

E.g. if the index is a Kibana system index like `.kibana_8.3.0` and to recover the first shard `0` , run:

```bash
bin/elasticsearch-shard remove-corrupted-data --index '.kibana_8.3.0' --shard-id 0
```

The tool opens the index and tries to fix it the corrupted parts of the index after getting the permission from the user.

Finally the tool requests the user to `reroute` the index to allocate a primary shard on a node. Sometimes, its needed to get the whole cluster up to `reroute` the index. In that case, you need to run the `reroute` after the node it attached back to the cluster.

### Attach the node back to the cluster

Now that the corruption is fixed, it’s time to start the Elasticsearch node and enable the shard allocations settings  to `all` . The node can be restarted by removing the annotations with the command

```bash
kubectl annotate es myeck eck.k8s.elastic.co/suspend-
```

 When the node gets back into the running state, ensure to mark the cluster level shard allocation settings `cluster.routing.allocation.enable` to `all` instead of `primaries`.

## Conclusion

Not many of us are aware of these great tools under the `$ELASTIC_HOME/bin` directory. This is just one use-case which help us to solve a complex file corruption problem which is hard to solve without these tools. Also, in cloud environments where we have limited access to hardware, solving problems does not get easier. This sets high expectations for the future to solve harder problems on more complex and limited-access environments.

## References

1. [Elasticsearch shard command](https://www.elastic.co/guide/en/elasticsearch/reference/current/shard-tool.html)
2. [Cluster level shard allocation settings](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-cluster.html#cluster-shard-allocation-settings)
3. [Suspend an ECK node](https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-troubleshooting-methods.html#k8s-suspend-elasticsearch)
