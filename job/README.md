# Generate UUID
This is a Kubernetes manifest that creates a Job. The Job creates a pod running an Ubuntu container, and the generates a list of test emails with UUIDS. The pod will run for 60 seconds and then terminate. 

The original commands are from the DigitalOcean Tutorial [How To Generate Universally Unique Identifiers (UUIDs) with uuidgen](https://www.digitalocean.com/community/tutorials/workflow-command-line-basics-generating-uuids). 



Here's the kind of output you should see when you look at the logs for the pod: 

```
Output
826119d2-f590-4fa3-ba7e-0717869d40b1@mailinator.com
795fec1a-76fe-4fed-8a06-ed517c1a5e7d@mailinator.com
14a502ad-0aa9-40e5-a46f-5806264b5316@mailinator.com
c6c2a588-7cce-4675-a490-0101d7bcc614@mailinator.com
7346c15b-0c92-44c4-a854-5de18c0c202d@mailinator.com
c67a535a-e28d-43b1-b553-c203bc22a821@mailinator.com
76d22d18-0f09-405d-9903-eb44ec93b605@mailinator.com
2b631756-21e6-4d95-873b-3245797f9028@mailinator.com
aab686e8-540e-43e9-9e24-ca04fbf4d414@mailinator.com
a577e9c9-0ad1-4934-b5f1-17b68938fff8@mailinator.com
``` 

## Run on a Kubernetes Cluster
Clone this repository and change into the daemonset directory
```bash
git clone https://github.com/do-community/example-k8s-workloads/ && \
cd example-k8s-workloads/job
```
Apply the Job manifest

`kubectl apply -f job.yaml`

## See the Logs 
To see the message, first find all the nodename pods 

`kubectl get pods` 

You will see one po that start with the prefix `uuidgen-` followed by a hash.  

To see the logs for a pod, copy the pod name and run 

`kubectl logs <pod_name>`

You should see the message above with the node name that your pod is running on. 

## Delete the Job 

`kubectl delete job uuidgen`

