#!/bin/bash

SCRIPTS="01 02 03"

docker-compose -f ./test/docker-compose.yml up -d

# Wait for ready all containers
echo "sleep 15s"
sleep 15

for f in $SCRIPTS
do
	./balerter -config test/config.yml -once -script test/scripts/$f.lua > $f.out
	diff $f.out test/out/$f.lua.txt
	status=$?
	rm $f.out

	if [ $status -eq 0 ]
	then
	  echo "test $f ok"
  else
    echo "test $f failed"
    docker-compose -f ./test/docker-compose.yml down -v
    exit 1
  fi

done

docker-compose -f ./test/docker-compose.yml down -v
