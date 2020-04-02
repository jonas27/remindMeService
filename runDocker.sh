#!/bin/bash
docker build -t remind-me-service .
# docker build -t min -f Dockerfile.local .
docker run -p 2001:2001 -v /volumes/remindMeService:/volumes/remindMeService -d remind-me-service