#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

cd ${PROJECT_ROOT}
cd frontend/ && npm run build && cd ../ && cd backend/ && go build -o main && cd ../ && echo 'building done'
