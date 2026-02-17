# Ethereum Balance API (Go + DDD)

Small Go API that reads Ethereum account balances from Infura and returns ETH values.

## Features

- `GET /address/balance/{ethAddress}`: returns current balance for an Ethereum address
- `GET /requests/history`: returns recent balance request history (demo feature)
- DDD-inspired structure (`domain`, `application`, `ports`, `infrastructure`, `interfaces`)
- Dockerized service
- Terraform for AWS EC2 deployment
- GitHub Actions CI/CD pipeline

## API examples

### Get balance

```bash
curl "http://localhost:8080/address/balance/0xc94770007dda54cF92009BFF0dE90c06F603a09f"
```

Response:

```json
{
  "balance": "0.0001365"
}
```

### Get request history

```bash
curl "http://localhost:8080/requests/history"
```

## Local development

### Prerequisites

- Go 1.22+
- Docker (optional for container run)

### Run locally

1. Copy env file:

```bash
cp .env.example .env
```

2. Edit `.env` and set your Infura URL:

```bash
INFURA_URL=https://mainnet.infura.io/v3/<your-project-id>
```

3. Run:

```bash
set -a && source .env && set +a
go run ./cmd/api
```

### Run tests

```bash
go test ./...
```

## Docker

Build:

```bash
docker build -t go-api:local .
```

Run:

```bash
docker run --rm -p 8080:8080 \
  -e PORT=8080 \
  -e INFURA_URL="https://mainnet.infura.io/v3/<your-project-id>" \
  go-api:local
```

## Terraform deployment (AWS EC2)

Files are in `terraform/`.

1. Create vars file:

```bash
cd terraform
cp terraform.tfvars.example terraform.tfvars
```

2. Fill `terraform.tfvars`:

- `ssh_allowed_cidr`: your public IP in CIDR format
- `public_key_path`: path to your SSH public key
- `docker_image`: Docker Hub image/tag
- `infura_url`: Infura endpoint

3. Deploy:

```bash
terraform init
terraform plan
terraform apply
```

4. Use output:

```bash
terraform output api_base_url
```

## GitHub Actions CI/CD

Workflow: `.github/workflows/ci-cd.yml`

On push to `main`:
- run tests
- build and push Docker image to Docker Hub
- SSH into EC2 and restart container with the new image

Configure these GitHub secrets:
- `DOCKERHUB_USERNAME`
- `DOCKERHUB_TOKEN`
- `EC2_HOST`
- `EC2_USER`
- `EC2_SSH_PRIVATE_KEY`
- `INFURA_URL`

## Pre-commit safety check

Run this before each commit/push to catch likely secrets:

```bash
./scripts/precommit-safety-check.sh
```

Do not commit:
- `.env`
- `terraform/terraform.tfvars`
- `terraform/*.tfstate`

## DDD mapping

- `internal/domain`: enterprise rules/value objects (`address`, `balance`)
- `internal/application`: use cases (`get_balance_usecase`, `get_request_history_usecase`)
- `internal/ports`: interfaces between application and adapters
- `internal/infrastructure`: adapters (Infura client, in-memory history repository)
- `internal/interfaces/http`: HTTP handlers and routing

## Live demo script

1. Call health endpoint:
   - `GET /healthz`
2. Call balance endpoint:
   - `GET /address/balance/{address}`
3. Show history endpoint:
   - `GET /requests/history`
4. Live feature extension idea:
   - add query params `?limit=10` to history endpoint and deploy through pipeline

## Follow-up questions (interview prep)

### How can we make deployment secure?

- Keep secrets in GitHub Secrets / AWS Secrets Manager (no hardcoded keys).
- Restrict security group ingress by CIDR and use HTTPS termination (ALB + ACM).
- Use least-privilege IAM role for EC2.
- Add WAF/rate limiting at edge layer.

### Is this HA? How improve HA?

- Current EC2 single instance is not HA.
- Improve with ALB + Auto Scaling Group across multiple AZs.
- Externalize state and use health checks with automatic replacement.

### How to deploy 100+ times/day?

- Trunk-based development + short-lived branches.
- Fast automated tests and static checks.
- Immutable image tags and progressive rollout with rollback.

### How to scale to thousands of customers/minute?

- Horizontal scale with ALB + autoscaling.
- Add request throttling and caching for repeated addresses.
- Use async work for heavy non-critical tasks.

### How to monitor and alert on outage?

- Structured app logs + CloudWatch metrics.
- Alerts on error rate, latency, and health check failures.
- Integrate alerts with SNS -> Slack/PagerDuty.
