#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

cd ${GOPATH}

let linesCount=$(cat main.go | wc -l)

cd src/

for dir in $(ls)
do
  if [[ ${dir} != github.com ]]; then
    cd ${dir}
    let linesCount=linesCount+$(find . -name '*.go' -type f -print0 | xargs -0 cat | wc -l)
    cd ../
  fi
done

echo ${linesCount}
