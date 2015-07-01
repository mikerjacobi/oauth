#!/usr/bin/python
import os

cmd = "docker ps -a | grep oauth | awk '{print $1}' | xargs docker rm -f"
os.system(cmd)
cmd = "go build && docker-compose up -d"
os.system(cmd)
