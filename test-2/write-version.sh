#!/bin/ash

version=$1
commitsha=$2
description=$3

jq  --arg ver "$version"  \
 --arg commit "$commitsha"  \
  --arg desc "$description"  \
'."anz-version-test"[0]."version"=$ver |  ."anz-version-test"[0]."lastcommitsha"=$commit | ."anz-version-test"[0]."description"=$desc'   resource/version.json \
> version.json.tmp && cp version.json.tmp resource/version.json && rm -rf version.json.tmp