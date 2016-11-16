#!/usr/bin/env bash
bash scripts/wait-for-it.sh go-app-db:3306
go-crud-template