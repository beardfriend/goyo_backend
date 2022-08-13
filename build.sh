#!/bin/bash

scp 
GO_ENV=production
go mod download
go build -o main .