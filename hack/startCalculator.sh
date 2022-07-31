#!/usr/bin/env bash

for i in {10000..10010}
do
nohup .././consistent-demo simple --port $i > "/tmp/simple.out.$i" &
done