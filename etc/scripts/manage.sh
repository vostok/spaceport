#!/usr/bin/env sh
ATTEMPTS_MAX_COUNT=40
TIMEOUT=10

ATTEMPT=1
while [ "$ATTEMPT" -le "$ATTEMPTS_MAX_COUNT" ]; do
  if nc -z management-api 6307; then
      break;
  fi;
  echo "$ATTEMPT attempt failed."

  if [ "$ATTEMPT" -eq "$ATTEMPTS_MAX_COUNT" ]; then
      echo "Can't ping management-api!"
      exit 1;
  fi;

  ATTEMPT=$(($ATTEMPT+1))
  sleep $TIMEOUT;
done

curl -v -X POST 'management-api:6307/streams/create' --header 'Content-Type: text/plain' -H 'Authorization: Hercules masterApiKey 123' --data '{
    "type": "base",
    "name": "logs",
    "partitions": 1,
    "shardingKey": [],
    "ttl": 86400000
}'
curl -v -X POST 'management-api:6307/streams/create' --header 'Content-Type: text/plain' -H 'Authorization: Hercules masterApiKey 123' --data '{
    "type": "base",
    "name": "metrics",
    "partitions": 1,
    "shardingKey": [],
    "ttl": 86400000
}'
curl -v -X POST 'management-api:6307/streams/create' --header 'Content-Type: text/plain' -H 'Authorization: Hercules masterApiKey 123' --data '{
    "type": "base",
    "name": "traces",
    "partitions": 1,
    "shardingKey": [],
    "ttl": 86400000
}'
curl -v -X POST 'management-api:6307/rules/set?key=key_11111111111111111111111111111111&pattern=metrics*&rights=rwm' -H 'Authorization: Hercules masterApiKey 123'
curl -v -X POST 'management-api:6307/rules/set?key=key_11111111111111111111111111111111&pattern=logs*&rights=rwm' -H 'Authorization: Hercules masterApiKey 123'
curl -v -X POST 'management-api:6307/rules/set?key=key_11111111111111111111111111111111&pattern=traces*&rights=rwm' -H 'Authorization: Hercules masterApiKey 123'
