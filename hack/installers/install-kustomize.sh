#!/bin/bash
set -eux -o pipefail

export VER=3.1.0
[ -e $DOWNLOADS/kustomize_${VER} ] || curl -sLf -o $DOWNLOADS/kustomize_${VER} https://github.com/kubernetes-sigs/kustomize/releases/download/v${VER}/kustomize_${VER}_linux_amd64
sudo cp $DOWNLOADS/kustomize_${VER} $BIN/kustomize
sudo chmod +x $BIN/kustomize
kustomize version
