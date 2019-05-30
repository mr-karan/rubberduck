# rubberduck

![alt text](http://rubberduckvba.com/Content/Images/ducky-color.png)

## What is Rubberduck

Rubberduck is a lightweight container (Alpine based) which installs a set of utilities which are handy while debugging issues in a container environment.
It also ships with an small `Golang` binary which runs a bare-bones API server.

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