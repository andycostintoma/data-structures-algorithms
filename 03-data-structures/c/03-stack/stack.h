#ifndef STACK_H
#define STACK_H

#include "dynamic-array.h"

// Stack structure
typedef struct
{
    DynamicArray array;
} Stack;

// Public functions for Stack
void stackInit(Stack *stack, int capacity);
void stackFree(Stack *stack);
int stackPeek(Stack *stack);
int stackSearch(Stack *stack, int value);
void stackPush(Stack *stack, int value);
int stackPop(Stack *stack);
int stackSize(Stack *stack);
void stackPrint(const Stack *stack);

#endif
