#!/usr/bin/env bash

protoc *.proto --go_out=plugins=grpc:.

echo "Success "