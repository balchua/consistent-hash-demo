#!/usr/bin/env bash

kill $(ps aux | grep 'consistent-demo calculate' | awk '{print $2}')