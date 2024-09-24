#!/bin/bash


for dockerfile in build/*.Dockerfile; do
  exeName=$(basename "$dockerfile" .Dockerfile)
  echo "Building $exeName"

  imageName="$CI_REGISTRY"/"$CI_PROJECT_PATH"/"$exeName"

  docker build -f "$dockerfile" -t "$imageName" . \
        -t "$imageName":"$CI_COMMIT_SHA"
  docker push "$imageName"
  docker push "$imageName":"$CI_COMMIT_SHA"
done
