#!/bin/bash
kill `pidof server`

timeout=10
for bufsize in 16384 32768 65536
do
  for nothreads in 1 2 4
  do
    for nosessions in 1 10 100
    do
    echo "Bufsize: $bufsize Threads: $nothreads Sessions: $nosessions"
    ./build/bin/server 0.0.0.0 55555 $nothreads $bufsize & srvpid=$!
    ./build/bin/client localhost 55555 $nothreads $bufsize $nosessions $timeout
    kill -9 $srvpid
    done
  done
done
