#!/bin/bash
set -e

REGISTRY=${REGISTRY:-gomall}
TAG=${TAG:-latest}

echo "=== 推送所有 Docker 镜像到 Registry ==="

for svc in product frontend cart order checkout payment email user; do
    echo "推送 ${REGISTRY}/${svc}:${TAG}..."
    docker push ${REGISTRY}/${svc}:${TAG}
done

echo "=== 推送完成 ==="
