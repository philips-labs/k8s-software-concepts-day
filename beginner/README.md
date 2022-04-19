# Beginner exercise
## nginx Deployment

To follow this exercise use following oreilly sandbox.

- [Kubernetes sandbox](https://learning.oreilly.com/scenarios/kubernetes-sandbox/9781492062820/)

## Bootstrap sandbox

```bash
git clone https://github.com/philips-labs/k8s-software-concepts-day.git
cd k8s-software-concepts-day/beginner
```
### Excercises

#### 1. Deploy NGINX deployment

In this exercise, you will have to fix the CPU/Memory request & limit.

1. Deploy the nginx deployment from the given yaml in this folder.
2. Fix the deployment issues.
3. Make sure pod status should be "Running" in the cluster.


#### 2. Deploy NGINX service

In this exercise we are exposing application via service. We should be able to access application via service


1. Deploy nginx service
2. curl request to the service
3. Figure out the external nodeport
   1. To get your sandbox url click the url as shown in the picture.
      1. First click the `+`
      2. Then choose `Select port to view on Host 1`
   2. Now fill in the port number --  Do you see the server response and what does it show ?

![](../sandbox-url.png)
