#!/usr/bin/env bash

GOAPP_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
ENV_FILE=$1
TEMPLATE_FILE=$2
TARGET_FILE=$(basename $TEMPLATE_FILE .tmpl.yaml)
bash ${GOAPP_ROOT}/scripts/gen-config.sh ${ENV_FILE} ${TEMPLATE_FILE} > ${GOAPP_ROOT}/configs/${TARGET_FILE}.yaml
