THIS_PATH="$(dirname "$(realpath "$0")")"
TMP_DIR=$THIS_PATH/../.tmp


# Reset OUT_DIR
rm -rf $OUT_DIR
mkdir -p $OUT_DIR

# Lambda artifact
rm -f $TMP_DIR/lambda_artifact
GOOS=linux go build -o $TMP_DIR/lambda_artifact $THIS_PATH/../src/hello.go
zip -qjr $OUT_DIR/lambda_artifact.zip $TMP_DIR/lambda_artifact