#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

source ${PROJECT_ROOT}/scripts/_set_tests_folders.sh

TESTING_COMMAND='go test'
for dir in "${folders[@]}"
do
  TESTING_COMMAND="${TESTING_COMMAND} $(echo ${dir} | sed 's/\.\///g')/... "
done
TESTING_COMMAND="${TESTING_COMMAND} -cover"

cd ${PROJECT_ROOT}

export TESTING_COMMAND=${TESTING_COMMAND}
docker-compose -f docker/docker-compose.test.yaml up -d

docker wait docker_go_api_1
docker logs docker_go_api_1

docker-compose -f docker/docker-compose.test.yaml down
