#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define STATIC_ARRAY_CAPACITY 10  // Maximum number of elements the array can hold

typedef struct {
    int data[STATIC_ARRAY_CAPACITY];
    int size;
} StaticArray;

/*
 * O(1) - Initializing a StaticArray with a given capacity
 */
void initArray(StaticArray *array) {
    array->size = 0;
}

/*
 * O(1) - Direct access to the index in the array.
 */
int get(StaticArray *array, int index) {
    if (index < 0 || index >= array->size) {
        printf("Index %d out of bounds!\n", index);
        exit(EXIT_FAILURE);
    }
    return array->data[index];
}

/*
 * O(n) - Linear search through the array to find the element.
 */
int search(StaticArray *array, int value) {
    for (int i = 0; i < array->size; i++) {
        if (array->data[i] == value) {
            return i;
        }
    }
    return -1;
}

/*
 * O(n) - Linear shift of elements to insert the new value.
 */
void insert(StaticArray *array, int index, int value) {
    if (index < 0 || index > array->size) {
        printf("Insert index %d out of bounds!\n", index);
        return;
    }
    if (array->size >= STATIC_ARRAY_CAPACITY) {
        printf("Array is full!\n");
        return;
    }
    for (int i = array->size; i > index; i--) {
        array->data[i] = array->data[i - 1];
    }
    array->data[index] = value;
    array->size++;
}

/*
 * O(n) - Linear shift of elements to delete the value at the given index.
 */
void delete(StaticArray *array, int index) {
    if (index < 0 || index >= array->size) {
        printf("Delete index %d out of bounds!\n", index);
        return;
    }
    for (int i = index; i < array->size - 1; i++) {
        array->data[i] = array->data[i + 1];
    }
    array->size--;
}

/*
 * O(1) - Appending to the end of the array.
 */
void append(StaticArray *array, int value) {
    if (array->size >= STATIC_ARRAY_CAPACITY) {
        printf("Array is full!\n");
        return;
    }
    array->data[array->size] = value;
    array->size++;
}

/*
 * O(1) - Removing and returning the last element.
 */
int pop(StaticArray *array) {
    if (array->size == 0) {
        printf("Cannot pop from an empty array!\n");
        exit(EXIT_FAILURE);
    }
    int value = array->data[array->size - 1];
    array->size--;
    return value;
}

/*
 * O(1) - Direct printing of the array.
 */
void printArray(const StaticArray *array) {
    printf("[");

    for (int i = 0; i < array->size; i++) {
        printf("%d", array->data[i]);
        if (i < array->size - 1) {
            printf(", ");
        }
    }

    printf("]\n");
}

int main() {
    StaticArray array;
    initArray(&array);  // Initialize array with fixed capacity

    append(&array, 0);
    append(&array, 1);
    append(&array, 2);
    append(&array, 3);
    append(&array, 4);

    printf("After appending: ");
    printArray(&array);

    insert(&array, 2, 99);
    printf("After insert at index 2: ");
    printArray(&array);

    delete(&array, 1);
    printf("After delete at index 1: ");
    printArray(&array);

    int popped = pop(&array);
    printf("Popped value: %d\n", popped);
    printf("After popping: ");
    printArray(&array);

    int foundIndex = search(&array, 3);
    if (foundIndex != -1) {
        printf("Found 3 at index %d\n", foundIndex);
    } else {
        printf("3 not found in the array\n");
    }

    printf("Final array: ");
    printArray(&array);

    return 0;
}
