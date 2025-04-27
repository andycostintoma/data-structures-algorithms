#include "02-dynamic-array.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

/**
 * Time complexity: O(1)
 * Initializes the dynamic array with a given initial capacity.
 */
void initArray(DynamicArray *array, int initialCapacity) {
    if (initialCapacity <= 0) {
        fprintf(stderr, "Error: Initial capacity must be greater than 0!\n");
        exit(EXIT_FAILURE);
    }
    array->data = malloc(initialCapacity * sizeof(int));
    if (!array->data) {
        fprintf(stderr, "Memory allocation failed!\n");
        exit(EXIT_FAILURE);
    }
    array->size = 0;
    array->capacity = initialCapacity;
}

/**
 * Time complexity: O(n)
 * Initializes the dynamic array from a static array of values.
 */
void initArrayFromStatic(DynamicArray *array, const int *values, int count) {
    if (count < 0) {
        fprintf(stderr, "Error: Count must be non-negative!\n");
        exit(EXIT_FAILURE);
    }
    initArray(array, count);
    for (int i = 0; i < count; i++) {
        array->data[i] = values[i];
    }
    array->size = count;
}

/**
 * Time complexity: O(n)
 * Resizes the array to the new capacity.
 */
static void resizeArray(DynamicArray *array, int newCapacity) {
    int *newData = realloc(array->data, newCapacity * sizeof(int));
    if (!newData) {
        fprintf(stderr, "Memory reallocation failed!\n");
        free(array->data);
        exit(EXIT_FAILURE);
    }
    array->data = newData;
    array->capacity = newCapacity;
}

/**
 * Time complexity: O(n)
 * Doubles the array's capacity if the size exceeds the current capacity.
 */
static void growIfNeeded(DynamicArray *array) {
    if (array->size >= array->capacity) {
        resizeArray(array, array->capacity * 2);  // O(n) time complexity
    }
}

/**
 * Time complexity: O(n)
 * Shrinks the array's capacity when it is less than one-quarter full.
 */
static void shrinkIfNeeded(DynamicArray *array) {
    if (array->capacity > 1 && array->size <= array->capacity / 4) {
        int newCapacity = array->capacity / 2;
        if (newCapacity < 1) newCapacity = 1;
        resizeArray(array, newCapacity);
    }
}

/**
 * Time complexity: O(1)
 * Returns the element at the specified index in the array.
 */
int get(DynamicArray *array, int index) {
    if (index < 0 || index >= array->size) {
        fprintf(stderr, "Index %d out of bounds!\n", index);
        exit(EXIT_FAILURE);
    }
    return array->data[index];
}

/**
 * Time complexity: O(n)
 * Searches for the value in the array and returns its index, or -1 if not found.
 */
int search(DynamicArray *array, int value) {
    for (int i = 0; i < array->size; i++) {
        if (array->data[i] == value) {
            return i;  // Return index of found element
        }
    }
    return -1;  // Return -1 if not found
}

/**
 * Time complexity: O(n)
 * Inserts a value at the specified index, shifting elements if needed.
 */
void insert(DynamicArray *array, int index, int value) {
    if (index < 0 || index > array->size) {
        fprintf(stderr, "Insert index %d out of bounds!\n", index);
        return;
    }
    growIfNeeded(array);  // O(n) in the worst case
    for (int i = array->size; i > index; i--) {
        array->data[i] = array->data[i - 1];
    }
    array->data[index] = value;
    array->size++;
}

/**
 * Time complexity: O(n)
 * Deletes the element at the specified index and shifts elements if needed.
 */
void delete(DynamicArray *array, int index) {
    if (index < 0 || index >= array->size) {
        fprintf(stderr, "Delete index %d out of bounds!\n", index);
        return;
    }
    for (int i = index; i < array->size - 1; i++) {
        array->data[i] = array->data[i + 1];
    }
    array->size--;
    shrinkIfNeeded(array);  // O(n) in the worst case
}

/**
 * Time complexity: O(1)
 * Appends a value to the end of the array.
 */
void append(DynamicArray *array, int value) {
    insert(array, array->size, value);
}

/**
 * Time complexity: O(1)
 * Removes the last element from the array and returns its value.
 */
int pop(DynamicArray *array) {
    if (array->size == 0) {
        fprintf(stderr, "Cannot pop from an empty array!\n");
        exit(EXIT_FAILURE);
    }
    int value = get(array, array->size - 1);
    delete(array, array->size - 1);
    return value;
}

/**
 * Time complexity: O(1)
 * Frees the memory allocated for the array and resets its properties.
 */
void freeArray(DynamicArray *array) {
    free(array->data);
    array->data = NULL;
    array->size = 0;
    array->capacity = 0;
}

/**
 * Time complexity: O(n)
 * Prints the elements of the array.
 */
void printArray(const DynamicArray *array) {
    printf("[");
    for (int i = 0; i < array->size; i++) {
        printf("%d", array->data[i]);
        if (i < array->size - 1) printf(", ");
    }
    printf("]\n");
}

int main() {
    DynamicArray array;
    int staticValues[] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
    initArrayFromStatic(&array, staticValues, 10);

    printf("After initialization: ");
    printArray(&array);

    insert(&array, 5, 99);
    printf("After insert at index 5: ");
    printArray(&array);

    delete(&array, 2);
    printf("After delete at index 2: ");
    printArray(&array);

    append(&array, 100);
    printf("After appending 100: ");
    printArray(&array);

    int popped = pop(&array);
    printf("Popped value: %d\n", popped);
    printf("After popping: ");
    printArray(&array);

    int foundIndex = search(&array, 1);
    printf("Found 1 at index: %d\n", foundIndex);

    printf("Final array: ");
    printArray(&array);

    freeArray(&array);
    return 0;
}
