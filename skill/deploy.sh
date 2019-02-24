#!/bin/sh

rm -rf deploy
mkdir deploy
GOOS=linux go build -o ./deploy/babyappLambda
pushd deploy
zip babyappLambda.zip babyappLambda
popd