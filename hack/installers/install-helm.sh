#!/bin/bash
set -eux -o pipefail

[ -e $DOWNLOADS/helm.tar.gz ] || curl -sLf -o $DOWNLOADS/helm.tar.gz https://storage.googleapis.com/kubernetes-helm/helm-v2.13.1-linux-amd64.tar.gz
tar -C /tmp/ -xf $DOWNLOADS/helm.tar.gz
sudo cp /tmp/linux-amd64/helm $BIN/helm
helm version --client
helm init --client-only
