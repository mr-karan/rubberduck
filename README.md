# rubberduck

![alt text](http://rubberduckvba.com/Content/Images/ducky-color.png)

## What is Rubberduck

Rubberduck is a lightweight container (Alpine based) which installs a set of utilities which are handy while debugging issues in a container environment.
It also ships with an small `Golang` binary which runs a bare-bones API server.

## ⚡️⚡️ Quick Deployment

`kubectl run -i --tty --rm debug --image=mrkaran/rubberduck:latest --restart=Never -- sh`

Spawns a pod in your namespace, gives access to all required tools to inspect networking issues and on `exit` it dies. This image is meant to be used for debugging purposes like these only.

### Contents of the image

List of utilities installed in the container:

    - iftop
    - drill
    - strace
    - curl
    - iperf
    - nmap
    - conntrack-tools
    - bash
    - jq
    - vim
    - nano
    - tree
    - ca-certificates
    - bind-tools

## Deployment on K8s cluster

Run `kubectl apply -f k8s/`

### What it does

- Creates a Namespace named `rubberduck`.
- Creates a Deployment resource, which runs `3` replicas of `rubberduck` pod.
- Exposes port `8080` on `rubberduck` container port using the service type `NodePort` inside the cluster.
