#include "dynamic-array.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

/**
 * Time complexity: O(1)
 * Initializes the dynamic array with a given initial capacity.
 */
void arrayInit(DynamicArray *array, int initialCapacity)
{
    if (initialCapacity <= 0)
    {
        fprintf(stderr, "Error: Initial capacity must be greater than 0!\n");
        exit(EXIT_FAILURE);
    }
    array->data = malloc(initialCapacity * sizeof(int));
    if (!array->data)
    {
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
void arrayInitFromStatic(DynamicArray *array, const int *values, int count)
{
    if (count < 0)
    {
        fprintf(stderr, "Error: Count must be non-negative!\n");
        exit(EXIT_FAILURE);
    }
    arrayInit(array, count);
    for (int i = 0; i < count; i++)
    {
        array->data[i] = values[i];
    }
    array->size = count;
}

/**
 * Time complexity: O(n)
 * Resizes the array to the new capacity.
 */
static void arrayResize(DynamicArray *array, int newCapacity)
{
    int *newData = realloc(array->data, newCapacity * sizeof(int));
    if (!newData)
    {
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
static void arrayGrowIfNeeded(DynamicArray *array)
{
    if (array->size >= array->capacity)
    {
        arrayResize(array, array->capacity * 2); // O(n) time complexity
    }
}

/**
 * Time complexity: O(n)
 * Shrinks the array's capacity when it is less than one-quarter full.
 */
static void arrayShrinkIfNeeded(DynamicArray *array)
{
    if (array->capacity > 1 && array->size <= array->capacity / 4)
    {
        int newCapacity = array->capacity / 2;
        if (newCapacity < 1)
            newCapacity = 1;
        arrayResize(array, newCapacity);
    }
}

/**
 * Time complexity: O(1)
 * Returns the element at the specified index in the array.
 */
int arrayGet(DynamicArray *array, int index)
{
    if (index < 0 || index >= array->size)
    {
        fprintf(stderr, "Index %d out of bounds!\n", index);
        exit(EXIT_FAILURE);
    }
    return array->data[index];
}

/**
 * Time complexity: O(n)
 * Searches for the value in the array and returns its index, or -1 if not found.
 */
int arraySearch(DynamicArray *array, int value)
{
    for (int i = 0; i < array->size; i++)
    {
        if (array->data[i] == value)
        {
            return i; // Return index of found element
        }
    }
    return -1; // Return -1 if not found
}

/**
 * Time complexity: O(n)
 * Inserts a value at the specified index, shifting elements if needed.
 */
void arrayInsert(DynamicArray *array, int index, int value)
{
    if (index < 0 || index > array->size)
    {
        fprintf(stderr, "Insert index %d out of bounds!\n", index);
        return;
    }
    arrayGrowIfNeeded(array); // O(n) in the worst case
    for (int i = array->size; i > index; i--)
    {
        array->data[i] = array->data[i - 1];
    }
    array->data[index] = value;
    array->size++;
}

/**
 * Time complexity: O(n)
 * Deletes the element at the specified index and shifts elements if needed.
 */
void arrayDelete(DynamicArray *array, int index)
{
    if (index < 0 || index >= array->size)
    {
        fprintf(stderr, "Delete index %d out of bounds!\n", index);
        return;
    }
    for (int i = index; i < array->size - 1; i++)
    {
        array->data[i] = array->data[i + 1];
    }
    array->size--;
    arrayShrinkIfNeeded(array); // O(n) in the worst case
}

/**
 * Time complexity: O(1)
 * Appends a value to the end of the array.
 */
void arrayAppend(DynamicArray *array, int value)
{
    arrayInsert(array, array->size, value);
}

/**
 * Time complexity: O(1)
 * Removes the last element from the array and returns its value.
 */
int arrayPop(DynamicArray *array)
{
    if (array->size == 0)
    {
        fprintf(stderr, "Cannot pop from an empty array!\n");
        exit(EXIT_FAILURE);
    }
    int value = arrayGet(array, array->size - 1);
    arrayDelete(array, array->size - 1);
    return value;
}

/**
 * Time complexity: O(1)
 * Frees the memory allocated for the array and resets its properties.
 */
void arrayFree(DynamicArray *array)
{
    free(array->data);
    array->data = NULL;
    array->size = 0;
    array->capacity = 0;
}

/**
 * Time complexity: O(n)
 * Prints the elements of the array.
 */
void arrayPrint(const DynamicArray *array)
{
    printf("[");
    for (int i = 0; i < array->size; i++)
    {
        printf("%d", array->data[i]);
        if (i < array->size - 1)
            printf(", ");
    }
    printf("]\n");
}
