#!/bin/bash
kill `pidof server`
./bin/server -conn_port 55555 -bufsize 16384 >> out.txt &
#1.为啥client自然关闭之后,强制关闭server会出现timeout?
#2.golang默认不reseue address吗
timeout=10
for bufsize in 16384 32768 65536
do
  for nothreads in 1 2 4
  do
    for nosessions in 1 10 100
    do
    echo "Bufsize: $bufsize Threads: $nothreads Sessions: $nosessions"
    #./bin/server -conn_port 55555 -bufsize $bufsize >> out.txt & srvpid=$!
    ../asio/build/bin/client localhost 55555 $nothreads $bufsize $nosessions $timeout
    #kill -9 $srvpid
    done
  done
done

kill `pidof server`
