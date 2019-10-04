#!/bin/bash
cd /webapp/gotodo/
go build
systemctl start gotodo.service