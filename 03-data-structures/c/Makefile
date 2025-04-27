CC = gcc
CFLAGS = -Wall -Wextra -g

# Targets for each individual example
run-01: 01-static-array.c
	$(CC) $(CFLAGS) 01-static-array.c -o run-01
	./run-01
	rm -f run-01

run-02: 02-dynamic-array.c 02-dynamic-array.h
	$(CC) $(CFLAGS) 02-dynamic-array.c -o run-02
	./run-02
	rm -f run-02

run-03: 03-stack.c 03-stack.h 02-dynamic-array.c 02-dynamic-array.h
	$(CC) $(CFLAGS) 03-stack.c 02-dynamic-array.c -o run-03
	./run-03
	rm -f run-03

