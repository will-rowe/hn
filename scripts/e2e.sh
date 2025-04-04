#!/usr/bin/env sh
set -euo pipefail

echo "Sending test payload to report-service..."

RESPONSE=$(curl -s -w "%{http_code}" -o /tmp/e2e_response.json -X POST http://report-service:8080/v1/reports \
  -H "Authorization: Bearer testtoken" \
  -H "Content-Type: application/json" \
  -d '{
        "dataset_id": "compose-dataset",
        "data_id": "compose-data",
        "media_type": "MEDIA_TYPE_TEXT",
        "violation_type": "VIOLATION_TYPE_PII",
        "description": "test from compose"
      }')

HTTP_CODE="${RESPONSE:(-3)}"
BODY=$(cat /tmp/e2e_response.json)

echo "Response code: $HTTP_CODE"
echo "Response body: $BODY"

if [[ "$HTTP_CODE" != "200" ]]; then
  echo "Expected 200 OK, got $HTTP_CODE"
  exit 1
fi

if ! echo "$BODY" | grep -q '"status":"RECEIVED"'; then
  echo "Response missing expected status field"
  exit 1
fi

if ! echo "$BODY" | grep -q '"reportId":"'; then
  echo "Response missing reportId"
  exit 1
fi

echo "End-to-end test passed!"
