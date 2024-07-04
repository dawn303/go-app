#!/usr/bin/env bash

env_file="$1"
template_file="$2"

GOAPP_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

source "${GOAPP_ROOT}/scripts/logging.sh"

if [ $# -ne 2 ];then
    onex::log::error "Usage: gen-config.sh manifests/env.local configs/markets.tmpl.yaml"
    exit 1
fi

source "${env_file}"

declare -A envs

set +u
for env in $(sed -n 's/^[^#].*${\(.*\)}.*/\1/p' ${template_file})
do
    if [ -z "$(eval echo \$${env})" ];then
        onex::log::error "environment variable '${env}' not set"
        missing=true
    fi
done

if [ "${missing}" ];then
    onex::log::error 'You may run `source manifests/env.local` to set these environment'
    exit 1
fi

eval "cat << EOF
$(cat ${template_file})
EOF"
