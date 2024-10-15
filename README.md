# hello-node

<p align="left">
 <a href="https://hub.docker.com/r/levaitamas/hello-node" alt="Docker pulls">
  <img src="https://img.shields.io/docker/pulls/levaitamas/hello-node" /></a>

A minimal webserver that returns the hostname. Useful for teaching and debugging K8s networking.

Available as:
- [docker.io/levaitamas/hello-node](https://hub.docker.com/r/levaitamas/hello-node)

## Usage examples

1. Create a deployment and a LoadBalancer service:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: hello
  name: hello
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hello
  template:
    metadata:
      labels:
        app: hello
    spec:
      containers:
      - name: hello
        image: docker.io/levaitamas/hello-node
---
apiVersion: v1
kind: Service
metadata:
  name: hello
spec:
  type: LoadBalancer
  ports:
  - port: 80
    targetPort: 8080
  selector:
    app: hello
```

2. Check connectivity:
```
$ curl $(kubectl get service hello -o jsonpath='{.status.loadBalancer.ingress[0].ip}'):80
Hello World from hello-6bcc85747b-mxwl4!
```

### K8s container networking
A short course on Kubernetes container networking:  http://lendulet.tmit.bme.hu/~levai/k8s/kubernetes_intro_aws.html

## Acknowledgments

Original idea and first implementation from [rg0now](https://github.com/rg0now).

## License

Licensed under [GPLv3+](LICENSE).
