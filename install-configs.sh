#!/bin/sh

if [ -n "$CI" ] && [ "$CI" = "true" ]; then
  CONFIG_ROOT=/root;
else
  CONFIG_ROOT=~;
fi
rm -rf $CONFIG_ROOT/.helper
mkdir -p $CONFIG_ROOT/.helper
cp -r configs/* $CONFIG_ROOT/.helper
echo "helper cli configs are installed to $CONFIG_ROOT/.helper"
