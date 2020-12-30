resource "aws_dynamodb_table" "stored_items" {
  name             = "stored_items"
  billing_mode     = "PAY_PER_REQUEST"
  stream_enabled   = true
  stream_view_type = "NEW_IMAGE"
  hash_key         = "PK"
  range_key        = "SK"

  attribute {
    name = "PK"
    type = "S"
  }
  attribute {
    name = "SK"
    type = "S"
  }
  attribute {
    name = "GSIPK"
    type = "S"
  }
  attribute {
    name = "GSISK"
    type = "S"
  }

  global_secondary_index {
    name            = "GSI"
    hash_key        = "GSIPK"
    range_key       = "GSISK"
    projection_type = "ALL"
  }
}
