#!/usr/bin/env bash
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

setFoldersWithTests
