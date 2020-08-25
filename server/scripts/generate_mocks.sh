#!/bin/bash


docker run -v "$PWD":/src -w /src -u $(id -u):$(id -g) vektra/mockery --all --keeptree
