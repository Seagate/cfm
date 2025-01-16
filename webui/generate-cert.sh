# Copyright (c) 2025 Seagate Technology LLC and/or its Affiliates

#!/bin/bash

# Define certificate file paths
KEY_FILE="incoming/key.pem"
CERT_FILE="incoming/cert.pem"

# Create the directory if it doesn't exist
mkdir -p incoming

# Generate self-signed certificate
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $KEY_FILE -out $CERT_FILE -subj "/CN=localhost"

echo "Self-signed certificate generated: $CERT_FILE"
echo "Private key generated: $KEY_FILE"