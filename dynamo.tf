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
    name = "priority"
    type = "N"
  }

  attribute {
    name = "enqueued_time"
    type = "S"
  }

  global_secondary_index {
    name = "priority-enqueued_time-index"
    hash_key = "priority"
    range_key = "enqueued_time"
    projection_type = "ALL"
  }
}