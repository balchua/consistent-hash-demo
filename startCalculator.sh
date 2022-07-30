#!/usr/bin/env bash

for i in {10000..10010}
do
nohup ./consistent-demo calculate --port $i > "/tmp/calculate.out.$i" &
done