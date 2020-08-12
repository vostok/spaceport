#!/usr/bin/env sh
ATTEMPTS_MAX_COUNT=20
TIMEOUT=10

ATTEMPT=1
while [ "$ATTEMPT" -le "$ATTEMPTS_MAX_COUNT" ]; do
  if test -f "/init/results/init_success"; then
      echo "Init container - success!";
      break;
  fi;
  echo "$ATTEMPT attempt failed."

  if [ "$ATTEMPT" -eq "$ATTEMPTS_MAX_COUNT" ]; then
      echo "Init container failed!"
      exit 1;
  fi;

  ATTEMPT=$(($ATTEMPT+1))
  sleep $TIMEOUT;
done

java -jar app.jar application.properties=file:///etc/hercules/application.properties
