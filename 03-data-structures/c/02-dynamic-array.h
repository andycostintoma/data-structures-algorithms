#ifndef DYNAMIC_ARRAY_H
#define DYNAMIC_ARRAY_H

typedef struct {
    int *data;
    int size;
    int capacity;
} DynamicArray;

// Public functions
void initArray(DynamicArray *array, int initialCapacity);
void initArrayFromStatic(DynamicArray *array, const int *values, int count);
int get(DynamicArray *array, int index);
int search(DynamicArray *array, int value);
void insert(DynamicArray *array, int index, int value);
void delete(DynamicArray *array, int index);
void append(DynamicArray *array, int value);
int pop(DynamicArray *array);
void freeArray(DynamicArray *array);
void printArray(const DynamicArray *array);

#endif
