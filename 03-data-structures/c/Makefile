CC = gcc
CFLAGS = -Wall -Wextra -g

# Automatically find all subdirectories for includes
INCLUDES = $(patsubst %,-I%,$(shell find . -type d))

# Targets for each individual example
run-01: 01-static-array/main.c
	$(CC) $(CFLAGS) $(INCLUDES) 01-static-array/main.c -o run-01
	./run-01
	rm -f run-01

run-02: 02-dynamic-array/main.c 02-dynamic-array/dynamic-array.c 02-dynamic-array/dynamic-array.h
	$(CC) $(CFLAGS) $(INCLUDES) 02-dynamic-array/main.c 02-dynamic-array/dynamic-array.c -o run-02
	./run-02
	rm -f run-02

run-03: 03-stack/main.c 03-stack/stack.c 03-stack/stack.h 02-dynamic-array/dynamic-array.c 02-dynamic-array/dynamic-array.h
	$(CC) $(CFLAGS) $(INCLUDES) 03-stack/main.c 03-stack/stack.c 02-dynamic-array/dynamic-array.c -o run-03
	./run-03
	rm -f run-03

run-04: 04-linked-list/main.c 04-linked-list/linked-list.c 04-linked-list/linked-list.h
	$(CC) $(CFLAGS) $(INCLUDES) 04-linked-list/main.c 04-linked-list/linked-list.c -o run-04
	./run-04
	rm -f run-04

run-05: 05-queue/main.c 05-queue/queue.c 05-queue/queue.h 04-linked-list/linked-list.c 04-linked-list/linked-list.h
	$(CC) $(CFLAGS) $(INCLUDES) 05-queue/main.c 05-queue/queue.c 04-linked-list/linked-list.c -o run-05
	./run-05
	rm -f run-05

run-06: 06-bst/main.c 06-bst/bst.c 06-bst/bst.h 02-dynamic-array/dynamic-array.c 02-dynamic-array/dynamic-array.h
	$(CC) $(CFLAGS) $(INCLUDES) 06-bst/main.c 06-bst/bst.c 02-dynamic-array/dynamic-array.c -o run-06
	./run-06
	rm -f run-06
