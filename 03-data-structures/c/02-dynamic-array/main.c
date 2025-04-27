#include "dynamic-array.h"
#include <stdio.h>

int main()
{
    DynamicArray array;
    int staticValues[] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
    arrayInitFromStatic(&array, staticValues, 10);

    printf("After initialization: ");
    arrayPrint(&array);

    arrayInsert(&array, 5, 99);
    printf("After insert at index 5: ");
    arrayPrint(&array);

    arrayDelete(&array, 2);
    printf("After delete at index 2: ");
    arrayPrint(&array);

    arrayAppend(&array, 100);
    printf("After appending 100: ");
    arrayPrint(&array);

    int popped = arrayPop(&array);
    printf("Popped value: %d\n", popped);
    printf("After popping: ");
    arrayPrint(&array);

    int foundIndex = arraySearch(&array, 1);
    printf("Found 1 at index: %d\n", foundIndex);

    printf("Final array: ");
    arrayPrint(&array);

    arrayFree(&array);
    return 0;
}