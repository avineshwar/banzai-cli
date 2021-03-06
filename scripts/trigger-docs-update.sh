#!/bin/bash

set -euf

PROJECT_SLUG='gh/banzaicloud/banzaicloud.github.io'
RELEASE_TAG="$1"

function main() {
  curl \
    -u "$CIRCLE_TOKEN:" \
    -X POST \
    --header "Content-Type: application/json" \
    -d "{
            \"branch\": \"gh-pages\",
            \"parameters\": {
                \"remote-trigger\": true,
                \"cli\": \"banzai-cli\",
                \"cli-release-tag\": \"$RELEASE_TAG\",
                \"cli-base-path\": \"/docs/pipeline/cli/reference/\"
            }
        }" "https://circleci.com/api/v2/project/$PROJECT_SLUG/pipeline"
}

main "$@"
