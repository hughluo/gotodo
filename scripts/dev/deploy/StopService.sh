#!/bin/bash
systemctl stop gotodo.service
rm -rf /webapp/gotodo/gotodo
go build