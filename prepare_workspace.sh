#!/usr/bin/env bash

function prepareFolders() {
  mkdir "$1"/backend/test_report 2>/dev/null
}

function prepareFiles() {
  rm "$1"/.env 2>/dev/null
  touch "$1"/.env
}

function installGolangDeps() {
  cd "$1"/backend && GOPATH="$1"/backend go get -v -d ./...
}

function installAngularDeps() {
  cd "$1"/frontend && npm install
}

function buildAngularApp() {
  cd "$1"/frontend && npm run build
}

rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  echo 'usage bash prepare_workspace.sh ROOT_FOLDER'
  exit 1
fi

declare -A env=(
  ['DB_USER']="gt_admin"
  ['DB_PASSWORD']="root"
  ['DB_NAME']="gt_db"
  ['ENV_VARS_WERE_SET']="1"
  ['PROJECT_ROOT']="${rootFolder}"
  ['REPORT_FOLDER']="${rootFolder}/backend/test_report"
  ['GOPATH']="${rootFolder}/backend"
  ['FRONTEND_DIR']="${rootFolder}/frontend"
  ['STATIC_DIR']="${rootFolder}/frontend/dist/get-together"
  ['CONN_STR']="\"host=localhost port=5432 user=gt_admin password=root dbname=gt_db sslmode=disable\""
  ['CODER_KEY']="123456789012345678901234"
  ['SHORT_MODE']="1"
  ['COMPOSE_PROJECT_NAME']='gt'
  ['API_PORT']="8080"
)

prepareFolders "${rootFolder}"
prepareFiles "${rootFolder}"
installGolangDeps "${rootFolder}"
installAngularDeps "${rootFolder}"
buildAngularApp "${rootFolder}"

for envKey in "${!env[@]}"
do
  echo "${envKey}=${env[$envKey]}" >> "${rootFolder}"/.env
done
