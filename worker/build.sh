#!/bin/bash

set -ex

go build -o bin/etlWorker etlWorker/*.go
