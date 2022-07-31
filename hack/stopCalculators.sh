#!/usr/bin/env bash

kill $(ps aux | grep 'consistent-demo simple' | awk '{print $2}')