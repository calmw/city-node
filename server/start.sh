#!/bin/bash

# city_node
/wkspace/city_node  2>&1 &

 # task
/wkspace/task  2>&1 &

# just keep this script running
while [[ true ]]; do
    sleep 1
done