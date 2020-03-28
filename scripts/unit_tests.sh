#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

source "${PROJECT_ROOT}"/scripts/_set_tests_folders.sh

TESTING_COMMAND='go test'
for dir in "${folders[@]}"
do
  TESTING_COMMAND="${TESTING_COMMAND} $(echo "${dir}" | sed 's/\.\///g')/... "
done
TESTING_COMMAND="${TESTING_COMMAND} -cover -p 1"

cd "${PROJECT_ROOT}" || exit

export RUN_GO_COMMAND=${TESTING_COMMAND}
docker-compose up -d --build api

docker wait "${COMPOSE_PROJECT_NAME}"_api_1
docker logs "${COMPOSE_PROJECT_NAME}"_api_1

docker-compose down
