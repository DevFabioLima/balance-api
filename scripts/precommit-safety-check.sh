#!/usr/bin/env bash
set -euo pipefail

echo "==> Pre-commit safety check"
echo

if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  echo "-- git status --"
  git status --short
  echo
else
  echo "Not inside a git repository. Skipping git status check."
  echo
fi

echo "-- potential secret scan --"
PATTERN='(api[_-]?key|secret|token|password|private[_-]?key|aws_access_key_id|aws_secret_access_key|xox[baprs]-|ghp_[a-zA-Z0-9]{20,})'
if command -v rg >/dev/null 2>&1; then
  rg -n -i \
    "${PATTERN}" \
    . \
    -g '!.git' \
    -g '!.env' \
    -g '!terraform/terraform.tfvars' || true
else
  grep -R -n -E -i \
    "${PATTERN}" \
    . \
    --exclude-dir=.git \
    --exclude=.env \
    --exclude=terraform.tfvars || true
fi
echo

echo "-- files that should not be committed --"
if [ -f ".env" ]; then
  echo "Found .env (OK if ignored, do not commit)."
fi
if [ -f "terraform/terraform.tfvars" ]; then
  echo "Found terraform/terraform.tfvars (OK if ignored, do not commit)."
fi

echo
echo "Pre-commit safety check complete."
