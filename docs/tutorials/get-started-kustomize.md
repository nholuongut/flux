# How to bootstrap Flux using Kustomize

This guide shows you how to use Kustomize to bootstrap Flux on a Kubernetes cluster.

## Prerequisites

You will need to have Kubernetes set up. For a quick local test,
you can use `minikube` or `kubeadm`. Any other Kubernetes setup
will work as well though.

### A note on GKE with RBAC enabled

If working on e.g. GKE with RBAC enabled, you will need to add a cluster role binding:

```sh
kubectl create clusterrolebinding "cluster-admin-$(whoami)" \
    --clusterrole=cluster-admin \
    --user="$(gcloud config get-value core/account)"
```

## Prepare Flux installation 

First you'll need a git repository to store your cluster desired state.
In our example we are going to use [`nholuongut/flux-get-started`](https://github.com/nholuongut/flux-get-started).
If you want to use that too, be sure to create a fork of it on GitHub.

Create a directory and add the `flux` namespace definition to it:

```sh
mkdir nholuongut

cat > nholuongut/namespace.yaml <<EOF
apiVersion: v1
kind: Namespace
metadata:
  name: flux
EOF
```

Create a kustomization file and use the Flux deploy YAMLs as base:

```sh
cat > nholuongut/kustomization.yaml <<EOF
namespace: flux
resources:
  - namespace.yaml
bases:
 - github.com/nholuongut/flux//deploy
patchesStrategicMerge:
  - patch.yaml
EOF
```

> **Note:** If you want to install a specific Flux release,
> add the version number to the base URL:
> `github.com/nholuongut/flux//deploy?ref=v1.14.1`

Create a patch file for Flux deployment and set the `--git-url`
parameter to point to the config repository
(replace `YOURUSER` with your GitHub username):

```sh
export GHUSER="YOURUSER"
cat > nholuongut/patch.yaml <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: flux
spec:
  template:
    spec:
      containers:
        - name: flux
          args:
            - --manifest-generation=true
            - --memcached-hostname=memcached.flux
            - --memcached-service=
            - --ssh-keygen-dir=/var/fluxd/keygen
            - --git-branch=master
            - --git-path=namespaces,workloads
            - --git-user=${GHUSER}
            - --git-email=${GHUSER}@users.noreply.github.com
            - --git-url=git@github.com:${GHUSER}/flux-get-started
EOF
```

We set `--git-path=namespaces,workloads` to exclude Helm manifests.
If you want to get started with Helm, please refer to the
["Get started with Flux using Helm"](get-started-helm.md) tutorial.

## Install Flux with Kustomize

In the next step, deploy Flux to the cluster (you'll need kubectl **1.14** or newer):

```sh
kubectl apply -k nholuongut
```

Wait for Flux to start:

```sh
kubectl -n flux rollout status deployment/flux
```

## Setup GitHub write access

At startup Flux generates a SSH key and logs the public key. Find
the SSH public key by installing [fluxctl](../references/fluxctl.md) and
running:

```sh
fluxctl identity --k8s-fwd-ns flux
```

In order to sync your cluster state with git you need to copy the
public key and create a deploy key with write access on your GitHub
repository.

Open GitHub, navigate to your fork, go to **Setting > Deploy keys**,
click on **Add deploy key**, give it a `Title`, check **Allow write
access**, paste the Flux public key and click **Add key**. See the
[GitHub docs](https://developer.github.com/v3/guides/managing-deploy-keys/#deploy-keys)
for more info on how to manage deploy keys.

## Committing a small change

In this example we'll be making a configuration change to a web application
and display a different message in the UI.

Replace `YOURUSER` in
`https://github.com/YOURUSER/flux-get-started/blob/master/workloads/podinfo-dep.yaml`
with your GitHub ID), open the URL in your browser, edit the file,
change the `PODINFO_UI_MESSAGE` env var to `Welcome to Flux` and commit the file.

By default, Flux git pull frequency is set to 5 minutes.
You can tell Flux to sync the changes immediately with:

```sh
fluxctl sync --k8s-fwd-ns flux
```

## Confirm the change landed

To access our webservice and check out its welcome message, simply
run:

```sh
kubectl -n demo port-forward deployment/podinfo 9898:9898 &
curl localhost:9898
```

Notice the updated `message` value in the JSON reply.

## Next steps

Try out [flux-kustomize-example](https://github.com/nholuongut/flux-kustomize-example)
for using Flux with Kustomize to manage
a staging and production clusters while minimizing duplicated declarations.

Try out [nholuongut/multi-tenancy](https://github.com/nholuongut/multi-tenancy)
for using Flux with Kustomize to manage a multi-tenant cluster.
