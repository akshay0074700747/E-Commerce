build_dir = bin/ecommerce-executable
wire_dir = ../wireeeeeee
version ?= latest

build:
	go build -o $(build_dir)

run: build
	$(build_dir)

wire-gen :
	cd $(wire_dir) && wire

docker-run :
	cd .. && docker build -t ecommerce:$(version) . --no-cache

	# use this --no-cache  if you updated the code

docker-compose :
	cd .. && docker-compose up