#!/bin/sh

THIS_PATH="$(dirname "$(realpath "$0")")"
TMP_DIR=$THIS_PATH/../.tmp

CONFIG_REL_DIR="configuration"

# Reset OUT_DIR
rm -rf $OUT_DIR
mkdir -p $OUT_DIR

# Lambda artifact
createLambdaArtifact() {
  LAMBDA_ARTIFACT=$1
  LAMBDA_FN_SRC_PATH=$2
  
  rm -f $TMP_DIR/$LAMBDA_ARTIFACT
  GOOS=linux GOARCH=amd64 go build -o $TMP_DIR/$LAMBDA_ARTIFACT $THIS_PATH/../src/$LAMBDA_FN_SRC_PATH
  zip -qjr $OUT_DIR/$LAMBDA_ARTIFACT.zip $TMP_DIR/$LAMBDA_ARTIFACT
  # Add config folder
  zip -gr $OUT_DIR/$LAMBDA_ARTIFACT.zip $CONFIG_REL_DIR
}

createLambdaArtifact hello_world_lambda hello.go
createLambdaArtifact watch_twitter handlers/watchTwitter/watchTwitter.go
createLambdaArtifact ddb_stream_consumer handlers/ddbStreamConsumer/ddbStreamConsumer.go
createLambdaArtifact main_ddb_stream_consumer handlers/mainDdbStreamConsumer/mainDdbStreamConsumer.go
createLambdaArtifact healthcheck_website handlers/healthcheckWebsite/healthcheckWebsite.go
