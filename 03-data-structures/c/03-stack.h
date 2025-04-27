#ifndef STACK_H
#define STACK_H

#include "02-dynamic-array.h"

// Stack structure
typedef struct {
    DynamicArray array;
} Stack;

// Public functions for Stack
void initStack(Stack *stack, int capacity);
void freeStack(Stack *stack);
void push(Stack *stack, int value);
int pop(Stack *stack);
int peek(Stack *stack);
int search(Stack *stack, int value);
int size(Stack *stack);
void printStack(const Stack *stack);

#endif
