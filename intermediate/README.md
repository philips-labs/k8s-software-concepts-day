# Intermediate exercise

This exercise has two parts:-
    1. Define and Mount a Persistent Volume
    2. Sidecar Containers

To follow this exercise use following oreilly sandbox.

- [Kubernetes sandbox](https://learning.oreilly.com/scenarios/kubernetes-sandbox/9781492062820/)

## Bootstrap sandbox

Launch a cluster:

```bash
launch.sh
```

Clone the exercise contents:

```bash
git clone https://github.com/philips-labs/k8s-software-concepts-day.git
```

## Define and Mount a Persistent Volume

Navigate to the volume directory:

```bash
cd k8s-software-concepts-day/intermediate/volume
```

### Exercises

#### 1. Create a Persistent Volume

In this exercise, you will have to create a persistent volume.

1. Deploy the PersistentVolume from the given yaml (pv.yaml) in this folder.
2. Make sure volume status should be "Available" in the cluster.

#### 2. Create a Persistent Volume Claim

In this exercise we are creating a PersistentVolumeClaim that requests the PersistentVolume from the previous step.

1. Deploy PersistentVolumeClaim from the given yaml(pvc.yaml) in this folder.
2. Ensure that the PersistentVolumeClaim is properly bound(Status should be Bound) after its creation. Fix the issues if not bound properly.

#### 3. Mounting the Persistent Volume from a Pod

In this exercise we are creating a new Pod that requests the PersistentVolume from the previous step.

1. Deploy nginx pod from the given yaml(nginx.yaml) in this folder.
2. Make sure PersistentVolume mount to the pod.
3. This can be checked by kubectl describe pod <pod_name>
4. Ensure pod should also be in "Running" State. Wait until the Pod enters the Running status and then check its events with the describe command.

## Deploy a pod with a sidecar container

Navigate to the volume directory:

```bash
cd ../sidecar
```

### Exercises

#### 1. Deploy a sidecar-webapp deployment 

In this exercise we deploy a sidecar-webapp from yaml (sidecar.yaml)

#### 2. Send curl request to service

Use curl to send request to the service. Troubleshoot any issues that you may have while doing so.

#### 3. Make containers in the same pod exchange the information

When curl request from the previous exercise finally starts working, the curl output will show nginx response with *403 Forbidden* in the body. Troubleshoot the response by making containers inside the pod exchange the information. The final curl output shall contain lines like

```text
Tue Apr 19 19:24:40 UTC 2022 Hi I am from a Sidecar container
```
