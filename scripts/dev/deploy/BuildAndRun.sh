#!/bin/bash
cd /webapp/gotodo/
docker build -t gotodo_dev
docker run -d -p 80:8080 --name gotodo_dev_running gotodo_dev