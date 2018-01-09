#!/bin/sh

goreleaser --skip-publish
if [ $? -ne 0 ]; then
  echo "No git tag, just snapshot"
  goreleaser --skip-publish --snapshot
fi
