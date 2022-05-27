#!/usr/bin/env bash
res=$(golint ./...)
if [[ -n "$res" ]]; then
    echo "Lint failed"
    echo "$res"
    exit 1
fi
