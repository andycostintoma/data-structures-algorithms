#include "stack.h"
#include <stdio.h>

int main()
{
    Stack stack;
    stackInit(&stack, 10);

    printf("After initialization: ");
    stackPrint(&stack);

    stackPush(&stack, 1);
    stackPush(&stack, 2);
    stackPush(&stack, 3);
    printf("After pushing 1, 2, 3: ");
    stackPrint(&stack);

    printf("Peek top element: %d\n", stackPeek(&stack));

    int poppedValue = stackPop(&stack);
    printf("Popped value: %d\n", poppedValue);
    printf("After popping: ");
    stackPrint(&stack);

    int foundIndex = stackSearch(&stack, 2);
    printf("Found 2 at index: %d\n", foundIndex);

    printf("Final stack: ");
    stackPrint(&stack);

    stackFree(&stack);
    return 0;
}