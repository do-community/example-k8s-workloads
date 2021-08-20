# Show Node Name DaemonSet 
This is a Kubernetes manifest that creates a DaemonSet. Each pod that the DaemonSet launches runs a busybox container that prints off a message showing the name of the node that is hosting the pod and an ASCII cat. 

Here's what you should see when you look at the logs for one of the pods: 

```
this daemon is running on node <node_name>

here's a cat =^..^=
``` 

The node name is set as an environment variable and the value is set using the [Kubernetes Downward API](https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/#the-downward-api). 

## Run on a Kubernetes Cluster 
```bash
git clone https://github.com/do-community/example-k8s-workloads/ && \
cd daemonset && \
kubectl apply -f ds.yaml \
```


## See the Logs 
To see the message, first find all the nodename pods 

`kubectl get pods` 

You will see a list of pods that start with the prefix `nodename-` and then have their own unique has. 

To see the logs for a pod, copy the pod name and run 

`kubectl logs <pod_name>`

You should see the message above with the node name that your pod is running on. 

## Delete the DaemonSet 

`kubectl delete ds nodename`

