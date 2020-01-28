#!/usr/bin/env bash

function prepareFolders() {
  mkdir $1/backend/test_report
}

function prepareFiles() {
  rm $1/.env 2>/dev/null
  touch $1/.env
}

function installAngularDeps() {
  cd $1/frontend && npm install
}

rootFolder=$1
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
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
  ['CONN_STR']="\"host=localhost port=5432 user=gt_admin password=root dbname=gt_db sslmode=disable\""
  ['CODER_KEY']="123456789012345678901234"
  ['SHORT_MODE']="1"
)

prepareFolders ${rootFolder}
prepareFiles ${rootFolder}
installAngularDeps ${rootFolder}

for envKey in "${!env[@]}"
do
  echo "${envKey}=${env[$envKey]}" >> ${rootFolder}/.env
done
