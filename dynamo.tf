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
}