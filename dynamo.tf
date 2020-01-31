provider "aws" {
  region = "us-east-1"
}

resource "aws_dynamodb_table" "priority-queue-table" {
  name = "PriorityQueue"
  hash_key = "id"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "queued"
    type = "N"
  }

  attribute {
    name = "priority_timestamp"
    type = "S"
  }

  global_secondary_index {
    name = "queued-priority_timestamp-index"
    hash_key = "queued"
    range_key = "priority_timestamp"
    projection_type = "ALL"
  }
}
