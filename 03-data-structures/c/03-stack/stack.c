#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

/**
 * Time complexity: O(1)
 * Returns the element at the top of the stack.
 * Returns an error if the stack is empty.
 */
int stackPeek(Stack *stack)
{
    if (stack->array.size == 0)
    {
        fprintf(stderr, "Stack is empty!\n");
        exit(EXIT_FAILURE);
    }
    return stack->array.data[stack->array.size - 1];
}

/**
 * Time complexity: O(n)
 * Searches for an item in the stack using the arraySearch function from DynamicArray.
 */
int stackSearch(Stack *stack, int value)
{
    return arraySearch(&stack->array, value);
}

/**
 * Time complexity: O(1)
 * Pushes an item onto the stack.
 */
void stackPush(Stack *stack, int value)
{
    arrayAppend(&stack->array, value);
}

/**
 * Time complexity: O(1)
 * Pops an item from the stack and returns its value.
 * Returns an error if the stack is empty.
 */
int stackPop(Stack *stack)
{
    if (stack->array.size == 0)
    {
        fprintf(stderr, "Stack underflow!\n");
        exit(EXIT_FAILURE);
    }
    return arrayPop(&stack->array);
}

/**
 * Time complexity: O(1)
 * Returns the current size of the stack.
 */
int stackSize(Stack *stack)
{
    return stack->array.size;
}

/**
 * Time complexity: O(1)
 * Initializes the stack with a given initial capacity.
 */
void stackInit(Stack *stack, int capacity)
{
    arrayInit(&stack->array, capacity);
}

/**
 * Time complexity: O(1)
 * Frees the memory allocated for the stack and resets its properties.
 */
void stackFree(Stack *stack)
{
    arrayFree(&stack->array);
}

/**
 * Time complexity: O(n)
 * Prints the elements of the stack using the arrayPrint function from DynamicArray.
 */
void stackPrint(const Stack *stack)
{
    arrayPrint(&stack->array);
}
