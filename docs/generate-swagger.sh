#!/bin/bash

DIRS=$(find ../ -type f -name "*.go" -exec dirname {} \; | sort -u | paste -sd "," -)
swag init -g ../cmd/main.go -d ../cmd,$DIRS -o ../docs