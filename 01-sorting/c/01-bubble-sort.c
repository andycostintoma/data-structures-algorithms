#include <stdio.h>
#include <stdbool.h>

void bubbleSort(int arr[], int size) {
    bool swapping = true;
    while (swapping) {
        swapping = false;
        for (int i = 1; i < size; i++) {
            if (arr[i - 1] > arr[i]) {
                // Swap
                int temp = arr[i];
                arr[i] = arr[i - 1];
                arr[i - 1] = temp;
                swapping = true;
            }
        }
    }
}

void printArray(int arr[], int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main() {
    int data[] = {5, 3, 8, 4, 2};
    int size = sizeof(data) / sizeof(data[0]);

    printf("Before sorting: ");
    printArray(data, size);

    bubbleSort(data, size);

    printf("After sorting: ");
    printArray(data, size);

    return 0;
}
