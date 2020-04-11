#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

cd "$PROJECT_ROOT"/backend && go fmt ./...
cd "$PROJECT_ROOT"/frontend && npm run lint -- --fix=true

cd "$PROJECT_ROOT" && git add . && git commit -m "$1" && git push
