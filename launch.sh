#!/bin/bash 

env -S $(grep -v '^#' .env) go run cmd/provico.go
