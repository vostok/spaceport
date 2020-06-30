#!/bin/bash
sleep 20
curl -v -X POST 'management-api:6307/streams/create' --header 'Content-Type: text/plain' -H 'masterApiKey: 123' --data '{
    "type": "base",
    "name": "logs",
    "partitions": 1,
    "shardingKey": [],
    "ttl": 86400000
}'
curl -v -X POST 'management-api:6307/streams/create' --header 'Content-Type: text/plain' -H 'masterApiKey: 123' --data '{
    "type": "base",
    "name": "metrics",
    "partitions": 1,
    "shardingKey": [],
    "ttl": 86400000
}'
curl -v -X POST 'management-api:6307/streams/create' --header 'Content-Type: text/plain' -H 'masterApiKey: 123' --data '{
    "type": "base",
    "name": "traces",
    "partitions": 1,
    "shardingKey": [],
    "ttl": 86400000
}'
curl -v -X POST 'management-api:6307/rules/set?key=key_11111111111111111111111111111111&pattern=metrics*&rights=rwm' -H 'masterApiKey: 123'
curl -v -X POST 'management-api:6307/rules/set?key=key_11111111111111111111111111111111&pattern=logs*&rights=rwm' -H 'masterApiKey: 123'
curl -v -X POST 'management-api:6307/rules/set?key=key_11111111111111111111111111111111&pattern=traces*&rights=rwm' -H 'masterApiKey: 123'
