#!/bin/bash

DIR=`dirname $0`
ABS_PATH=`cd $DIR; pwd`

docker run --rm --network=host -v $ABS_PATH/etc/properties/gateway-client/log:/etc/hercules vstk/hercules-gateway-client:0.44.1-SNAPSHOT
docker run --rm --network=host -v $ABS_PATH/etc/properties/gateway-client/metric:/etc/hercules vstk/hercules-gateway-client:0.44.1-SNAPSHOT
docker run --rm --network=host -v $ABS_PATH/etc/properties/gateway-client/trace:/etc/hercules vstk/hercules-gateway-client:0.44.1-SNAPSHOT
