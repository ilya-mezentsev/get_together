#!/usr/bin/env bash

function prepareFolders() {
  mkdir $1/backend/test_report
}

function prepareFiles() {
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

envVars=(
  "ENV_VARS_WERE_SET=1"
  "PROJECT_ROOT=${rootFolder}"
  "REPORT_FOLDER=${rootFolder}/backend/test_report"
  "GOPATH=${rootFolder}/backend"
  "FRONTEND_DIR=${rootFolder}/frontend"
)

prepareFolders ${rootFolder}
prepareFiles ${rootFolder}
installAngularDeps ${rootFolder}

for envVar in ${envVars[@]}; do
  echo ${envVar} >> ${rootFolder}/.env
done
