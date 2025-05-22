# Makefile
app_name := uala-posts-service
version ?= latest

.PHONY: build

build:
	docker build -t $(app_name):$(version) .