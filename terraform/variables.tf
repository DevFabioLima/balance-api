variable "aws_region" {
  type        = string
  description = "AWS region"
  default     = "us-east-1"
}

variable "project_name" {
  type        = string
  description = "Project name prefix for resources"
  default     = "eth-balance-api"
}

variable "instance_type" {
  type        = string
  description = "EC2 instance type"
  default     = "t3.micro"
}

variable "ssh_allowed_cidr" {
  type        = string
  description = "CIDR allowed to SSH into EC2"
}

variable "api_allowed_cidr" {
  type        = string
  description = "CIDR allowed to access API port"
  default     = "0.0.0.0/0"
}

variable "public_key_path" {
  type        = string
  description = "Path to SSH public key"
}

variable "docker_image" {
  type        = string
  description = "Docker image including tag (example: user/repo:latest)"
}

variable "infura_url" {
  type        = string
  description = "Infura URL used by the API"
}

variable "port" {
  type        = number
  description = "Application container port"
  default     = 8080
}
