#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

cd ${FRONTEND_DIR}/src && find . -name '*' -type f -print0 | xargs -0 cat | wc -l
