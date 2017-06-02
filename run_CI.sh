#!/bin/bash
shopt -s globstar
rootdir=`pwd`

for dir in **/ ; do
  cd $dir
  exec 5>&1
  out=$(go test -v -coverprofile=cov_part.out | tee >(cat - >&5))
  if [ $? -eq 1 ] ; then
    if [ $(cat $out | grep -o 'no buildable Go source files') == "" ] ; then
      echo "Tests failed! Exiting..." ; exit 1
    fi
  fi
  cd $rootdir
done

find . -name cov_part.out | xargs cat > cov.out
# make sure we do not run the ruby gem which is first in $PATH :(
node /usr/local/lib/node_modules/codeclimate-test-reporter/bin/codeclimate.js < cov.out