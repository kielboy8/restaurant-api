#!/bin/bash

set -ex

go build -o bin/service service/*.go
