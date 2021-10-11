#!/bin/bash

set +e
set -v

CWD=$PWD

cd $CWD/resource/bucket && sls deploy --verbose || error=true
cd $CWD/service && ./deploy.sh
cd $CWD/worker && ./deploy.sh