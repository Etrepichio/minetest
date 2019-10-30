#!/usr/bin/env bash

cd ../../
./build.sh
cd docker/dev
./docker.sh up --build