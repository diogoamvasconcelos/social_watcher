THIS_PATH="$(dirname "$(realpath "$0")")"

OUT_DIR=$THIS_PATH/../.build
TF_DIR=$THIS_PATH/../tf
REGION=eu-north-1

env OUT_DIR=$OUT_DIR \
  TF_VAR_out_dir=$OUT_DIR \
  TF_DIR=$TF_DIR \
  TF_VAR_tf_dir=$TF_DIR \
  AWS_REGION=$REGION \
  TF_VAR_aws_region=$REGION \
  $@
