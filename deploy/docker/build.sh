#!/bin/bash
set -e

REGISTRY=${REGISTRY:-gomall}
TAG=${TAG:-latest}

echo "=== 构建所有 Docker 镜像 ==="

echo "构建 product..."
docker build -f Dockerfile.product -t ${REGISTRY}/product:${TAG} ../..

echo "构建 frontend..."
docker build -f Dockerfile.frontend -t ${REGISTRY}/frontend:${TAG} ../..

echo "构建 cart..."
docker build -f Dockerfile.cart -t ${REGISTRY}/cart:${TAG} ../..

echo "构建 order..."
docker build -f Dockerfile.order -t ${REGISTRY}/order:${TAG} ../..

echo "构建 checkout..."
docker build -f Dockerfile.checkout -t ${REGISTRY}/checkout:${TAG} ../..

echo "构建 payment..."
docker build -f Dockerfile.payment -t ${REGISTRY}/payment:${TAG} ../..

echo "构建 email..."
docker build -f Dockerfile.email -t ${REGISTRY}/email:${TAG} ../..

echo "构建 user..."
docker build -f Dockerfile.user -t ${REGISTRY}/user:${TAG} ../..

echo "=== 构建完成 ==="
echo "查看镜像: docker images | grep ${REGISTRY}"
