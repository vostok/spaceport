#!/usr/bin/env sh
rm -r /init/results/init_success

ATTEMPTS_MAX_COUNT=40
TIMEOUT=10

ATTEMPT=1
while [ "$ATTEMPT" -le "$ATTEMPTS_MAX_COUNT" ]; do
  if nc -z cassandra 9042; then
      break;
  fi;
  echo "$ATTEMPT attempt failed."

  if [ "$ATTEMPT" -eq "$ATTEMPTS_MAX_COUNT" ]; then
      echo "Can't ping cassandra!"
      exit 1;
  fi;

  ATTEMPT=$(($ATTEMPT+1))
  sleep $TIMEOUT;
done

java -jar app.jar application.properties=file:///etc/hercules/application.properties init-zk=true init-kafka=true init-tracing-cassandra=true
echo 1 > /init/results/init_success
exit 0
