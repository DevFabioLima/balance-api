output "instance_public_ip" {
  description = "Public IP address of the API server"
  value       = aws_instance.api.public_ip
}

output "api_base_url" {
  description = "Base URL to call the API"
  value       = "http://${aws_instance.api.public_ip}:${var.port}"
}
