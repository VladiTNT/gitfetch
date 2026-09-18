# Variables
CC = g++
FLAGS = -O3

# PHONY
.PHONY: bin/main

# Build project
bin/main:
	${CC} ${FLAGS} src/main.cpp -o $@

# Clean
clean:
	rm -rf bin/*
