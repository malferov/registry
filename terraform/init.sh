#!/bin/bash
set -e
jinja2 backend.hcl -D be_token=$be_token > backend.hcl.rendered
terraform init -backend-config=backend.hcl.rendered
