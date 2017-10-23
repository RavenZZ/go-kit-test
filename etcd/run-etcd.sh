#!/bin/bash

export NODE1=10.10.100.188

REGISTRY=quay.io/coreos/etcd
#REGISTRY=gcr.io/etcd-development/etcd
DATA_DIR=/usr/local/golang/src/ravenzz/go-kit-test/etcd/data

docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  --volume=${DATA_DIR}:/etcd-data \
  --name etcd ${REGISTRY}:latest \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${NODE1}:2380 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://${NODE1}:2379 \
  --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://${NODE1}:2380



