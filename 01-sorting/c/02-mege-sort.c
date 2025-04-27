#include <stdio.h>
#include <stdlib.h>

void merge(int* arr, int left, int mid, int right) {
    int size1 = mid - left + 1;
    int size2 = right - mid;

    int* L = malloc(size1 * sizeof(int));
    int* R = malloc(size2 * sizeof(int));

    for (int i = 0; i < size1; i++) L[i] = arr[left + i];
    for (int j = 0; j < size2; j++) R[j] = arr[mid + 1 + j];

    int i = 0, j = 0, k = left;
    while (i < size1 && j < size2) {
        if (L[i] < R[j]) {
            arr[k++] = L[i++];
        } else {
            arr[k++] = R[j++];
        }
    }

    while (i < size1) arr[k++] = L[i++];
    while (j < size2) arr[k++] = R[j++];

    free(L);
    free(R);
}

void mergeSort(int* arr, int left, int right) {
    if (left < right) {
        int mid = (left + right) / 2;
        mergeSort(arr, left, mid);
        mergeSort(arr, mid + 1, right);
        merge(arr, left, mid, right);
    }
}

void printArray(int* arr, int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main() {
    int arr[] = {3, 7, 4, 12, 6, 5};
    int size = sizeof(arr) / sizeof(arr[0]);

    printf("Before sorting: ");
    printArray(arr, size);

    mergeSort(arr, 0, size - 1);

    printf("After sorting: ");
    printArray(arr, size);

    return 0;
}
