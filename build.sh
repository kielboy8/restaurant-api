#!/bin/bash

set -e
set -v


go get ./...

CWD=`pwd`

cd $CWD/service && ./build.sh

cd $CWD/worker && ./build.sh
