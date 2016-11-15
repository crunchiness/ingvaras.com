#!/usr/bin/env bash

# Allows to run script from anywhere
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"

# Deploy to GAE
/opt/sdk/google-appengine-go/appcfg.py -A original-storm-93323 update ./
