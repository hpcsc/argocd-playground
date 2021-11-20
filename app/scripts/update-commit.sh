#!/bin/bash

set -euo pipefail

SCRIPT_DIR=$(cd $(dirname $0); pwd)
VERSION_FILE=${SCRIPT_DIR}/../version.json
echo "$(jq --arg commit "${COMMIT}" \
            '.commit |= $commit' \
            ${VERSION_FILE})" \
    > ${VERSION_FILE}

echo "Updated version in ${VERSION_FILE} to ${COMMIT}"
