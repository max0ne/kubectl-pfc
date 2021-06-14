# kubectl-pfc

_kubectl port-forward-curl_

A kubectl plugin to turn `kubectl port-forward` and `curl localhost:...` into 1 command.

## Usage

This is 1 command to:

1. Instantiate a `kubectl port-forward`
2. Issue curl to your localhost against the port listening by port-forward, with stdout/stderr piped
3. Close that `port-forward`

Example:

1. This makes a port-forward against pod `backend-0` of port :80. It then queries the `/api` path of this service:
    ```sh
    # k pfc <args to normal port-forward command> -- curl <args to curl command>
    kubectl pfc backend-0 :80 -- curl localhost/api
    ```

2. This makes a port-forward against deployment `kube-state-metrics` in namespace kube-system of port :8080. It then queries the `/metrics` path of this service:
    ```sh
    # k pfc <args to normal port-forward command> -- curl <args to curl command>
    kubectl pfc -n kube-system deploy/kube-state-metrics :8080 -- curl localhost/metrics
    ```

3. This makes a port-forward against service `nginx` in namespace ingress of port :443. It then issues an **HTTPS** call against this service, using `api.my-site.com` as HTTP host header and TLS SNI:
    ```sh
    # k pfc <args to normal port-forward command> -- curl <args to curl command>
    kubectl pfc -n ingress svc/nginx :443 -- curl https://api.my-site.com
    ```

## Install

1. Install Go, have `$GOPATH/bin` as part of your $PATH

2.
    ```sh
    go install github.com/max0ne/kubectl-pfc
    ```

3. `kubectl` should automatically register `pfc` as a plugin, you can start curling
