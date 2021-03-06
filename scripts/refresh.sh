#!/bin/bash
# script to update project deps in gvt manifest
#

go get github.com/cespare/deplist
rm  -rf vendor/
for i in $(go list -f '{{.ImportPath}}' ./... | xargs -n 1 deplist | grep -v 'pubsub-cleaner' | grep -v vendor | sort -u); do
  gvt fetch  $i
done
