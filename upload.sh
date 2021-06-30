#!/bin/sh

curl -vvv --form file='@test.txt' localhost:8888/upload
