#!/usr/bin/env bash

usage() {
  echo "Usage: $0 [-vh]" 1>&2
  echo -e "\t-v:\tVerbose"
  echo -e "\t-h:\tShow usage"
  exit 1
}

## Download dep for dep resolution
if [ -z "$(which dep)" ]; then
  echo Installing dep...
  go get -u github.com/golang/dep/cmd/dep
  echo Done.
fi

while getopts ":vh" o; do
    case "${o}" in
        v)
            GOFLAGS="-x"
            ;;
        h)
            usage
            ;;
        *)
            echo Unknown flag: ${OPTARG}
            usage
            ;;
    esac
done
shift $((OPTIND-1))

## update deps
VENDOR=vendor/
if [ -d $VENDOR ]; then
  rm -rf $VENDOR
fi

dep ensure -v

## Download bindata for db migrations
if [ -z "$(which go-bindata)" ]; then
  echo Installing go-bindata...
  go get -u github.com/jteeuwen/go-bindata/...
  echo Done.
fi


CGO_ENABLED=0 GOOS=linux go build ${GOFLAGS} -a \
    -installsuffix cgo \
    github.com/minesweeper/cmd/minesweeper
