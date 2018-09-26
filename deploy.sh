#!/usr/bin/env bash

# Allows to run script from anywhere
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"

# Deploy to GAE
gcloud app deploy --project ingvaras app.yaml
