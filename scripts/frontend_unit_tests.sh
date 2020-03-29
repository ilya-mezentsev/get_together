#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

export RUN_FRONTEND_COMMAND="cd ${CONTAINER_FRONTEND_SRC} && npm run test"
docker-compose up -d --build frontend

docker wait "${COMPOSE_PROJECT_NAME}"_frontend_1
docker logs "${COMPOSE_PROJECT_NAME}"_frontend_1

docker-compose down
