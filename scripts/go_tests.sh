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

folders=()
function setFoldersWithTests() {
  cd ${GOPATH}/src

  for dir in $(find . -type d)
  do
    # should skip github libs
    if [[ ${dir} == *github* ]]; then
      continue
    fi
    if tests=$(find ${GOPATH}/src/${dir} -maxdepth 1 -name '*_test.go'); [[ ${tests} != "" ]]; then
      folders+=(${dir})
    fi
  done
}

rm -rf ${REPORT_FOLDER}/*
setFoldersWithTests
for dir in "${folders[@]}"
do
  reportFileName=$(echo -n ${dir} | md5sum | awk '{print $1}')
  cd ${GOPATH}/src/${dir} && TESTING_MODE=1 go test -coverprofile=${REPORT_FOLDER}/${reportFileName}.out
  if [[ $1 = html ]]; then # open reports in browser
    go tool cover -html=${REPORT_FOLDER}/${reportFileName}.out
  fi
done
