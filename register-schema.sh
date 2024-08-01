#!/bin/bash

# Register the user schema
curl -X POST "$SCHEMA_REGISTRY_URL/subjects/user-events-value/versions" \
     -H "Content-Type: application/vnd.schemaregistry.v1+json" \
     --data @user-schema.json \
     --silent \
     --output /dev/null

echo "User schema registered successfully."

# Register the order schema
curl -X POST "$SCHEMA_REGISTRY_URL/subjects/order-events-value/versions" \
     -H "Content-Type: application/vnd.schemaregistry.v1+json" \
     --data @order-schema.json \
     --silent \
     --output /dev/null

echo "Order schema registered successfully."
