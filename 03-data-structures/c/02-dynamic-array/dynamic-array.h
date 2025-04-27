#ifndef DYNAMIC_ARRAY_H
#define DYNAMIC_ARRAY_H

typedef struct
{
    int *data;
    int size;
    int capacity;
} DynamicArray;

// Public functions
void arrayInit(DynamicArray *array, int initialCapacity);
void arrayInitFromStatic(DynamicArray *array, const int *values, int count);
int arrayGet(DynamicArray *array, int index);
int arraySearch(DynamicArray *array, int value);
void arrayInsert(DynamicArray *array, int index, int value);
void arrayDelete(DynamicArray *array, int index);
void arrayAppend(DynamicArray *array, int value);
int arrayPop(DynamicArray *array);
void arrayFree(DynamicArray *array);
void arrayPrint(const DynamicArray *array);

#endif
