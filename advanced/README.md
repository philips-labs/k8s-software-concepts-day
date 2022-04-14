# Advanced exercise

## http-echo operator

To follow this exercise use following oreilly sandbox.

[Kubernetes sandbox](https://learning.oreilly.com/scenarios/kubernetes-sandbox/9781492062820/)

### Steps

#### Bootstrap sandbox

```bash
git clone https://github.com/philips-labs/k8s-software-concepts-day.git
cd k8s-software-concepts-day/advanced
./install-operator.sh
source ./install-go.sh
source ./install-registry.sh
```

#### Install the operator

```bash
cd http-echo-operator
make generate
make manifests
export IMG=$REGISTRY/http-echo-operator:v0.0.1
make docker-build
make deploy
kubectl get services,deployments,pods --namespace http-echo-operator-system
```
