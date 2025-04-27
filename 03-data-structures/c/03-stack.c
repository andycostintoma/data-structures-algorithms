#include <stdio.h>
#include <stdlib.h>
#include "03-stack.h"

/**
 * Time complexity: O(1)
 * Returns the element at the top of the stack.
 * Returns an error if the stack is empty.
 */
int peek(Stack *stack) {
    if (stack->array.size == 0) {
        fprintf(stderr, "Stack is empty!\n");
        exit(EXIT_FAILURE);
    }
    return stack->array.data[stack->array.size - 1];
}

/**
 * Time complexity: O(n)
 * Searches for an item in the stack using the search function from DynamicArray.
 */
int search(Stack *stack, int value) {
    return search(&stack->array, value);
}

/**
 * Time complexity: O(1)
 * Pushes an item onto the stack.
 */
void push(Stack *stack, int value) {
    append(&stack->array, value);
}

/**
 * Time complexity: O(1)
 * Pops an item from the stack and returns its value.
 * Returns an error if the stack is empty.
 */
int pop(Stack *stack) {
    if (stack->array.size == 0) {
        fprintf(stderr, "Stack underflow!\n");
        exit(EXIT_FAILURE);
    }
    return pop(&stack->array);
}

/**
 * Time complexity: O(1)
 * Returns the current size of the stack.
 */
int size(Stack *stack) {
    return stack->array.size;
}

/**
 * Time complexity: O(1)
 * Initializes the stack with a given initial capacity.
 */
void initStack(Stack *stack, int capacity) {
    initArray(&stack->array, capacity);
}

/**
 * Time complexity: O(1)
 * Frees the memory allocated for the stack and resets its properties.
 */
void freeStack(Stack *stack) {
    freeArray(&stack->array);
}

/**
 * Time complexity: O(n)
 * Prints the elements of the stack using the printArray function from DynamicArray.
 */
void printStack(const Stack *stack) {
    printArray(&stack->array);
}

int main() {
    Stack stack;
    initStack(&stack, 10);

    printf("After initialization: ");
    printStack(&stack);

    push(&stack, 1);
    push(&stack, 2);
    push(&stack, 3);
    printf("After pushing 1, 2, 3: ");
    printStack(&stack);

    printf("Peek top element: %d\n", peek(&stack));

    int poppedValue = pop(&stack);
    printf("Popped value: %d\n", poppedValue);
    printf("After popping: ");
    printStack(&stack);

    int foundIndex = search(&stack, 2);
    printf("Found 2 at index: %d\n", foundIndex);

    printf("Final stack: ");
    printStack(&stack);

    freeStack(&stack);
    return 0;
}
