#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi
# we need to check that REPORT_FOLDER is subdirectory for GOPATH
if [[ ${REPORT_FOLDER}/ != ${GOPATH}* ]]; then
  echo 'go tests report folder should be in GOPATH'
  exit 1
fi

source ${PROJECT_ROOT}/scripts/_set_tests_folders.sh
rm -rf ${REPORT_FOLDER}/*
for dir in "${folders[@]}"
do
  reportFileName=$(echo -n ${dir} | md5sum | awk '{print $1}')
  cd ${GOPATH}/src/${dir} && go test -coverprofile=${REPORT_FOLDER}/${reportFileName}.out
  if [[ $1 = html ]]; then # open reports in browser
    go tool cover -html=${REPORT_FOLDER}/${reportFileName}.out
  fi
done
