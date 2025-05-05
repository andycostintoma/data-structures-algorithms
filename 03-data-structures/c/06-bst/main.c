#include <stdio.h>
#include "bst.h"
#include "dynamic-array.h"

int main()
{
    // Create a BST with root value 10
    BSTNode *root = bstCreate(10);

    // Insert some values into the BST
    bstInsert(root, 5);
    bstInsert(root, 15);
    bstInsert(root, 3);
    bstInsert(root, 7);
    bstInsert(root, 12);
    bstInsert(root, 18);

    // Print the BST structure
    printf("BST structure:\n");
    bstPrint(root, 0);

    // Search for a value in the BST
    int searchVal = 7;
    if (bstExists(root, searchVal))
    {
        printf("\nValue %d exists in the BST.\n", searchVal);
    }
    else
    {
        printf("\nValue %d does not exist in the BST.\n", searchVal);
    }

    // Get the minimum and maximum values
    int *minVal = bstGetMin(root);
    int *maxVal = bstGetMax(root);
    printf("\nMinimum value in the BST: %d\n", *minVal);
    printf("Maximum value in the BST: %d\n", *maxVal);

    // Get the height of the BST
    int height = bstHeight(root);
    printf("\nHeight of the BST: %d\n", height);

    // Perform in-order traversal and print the result
    DynamicArray inorder = bstInorderTraversal(root);
    printf("\nIn-order traversal: ");
    for (int i = 0; i < inorder.size; i++)
    {
        printf("%d ", inorder.data[i]);
    }
    printf("\n");

    // Perform pre-order traversal and print the result
    DynamicArray preorder = bstPreorderTraversal(root);
    printf("\nPre-order traversal: ");
    for (int i = 0; i < preorder.size; i++)
    {
        printf("%d ", preorder.data[i]);
    }
    printf("\n");

    // Perform post-order traversal and print the result
    DynamicArray postorder = bstPostorderTraversal(root);
    printf("\nPost-order traversal: ");
    for (int i = 0; i < postorder.size; i++)
    {
        printf("%d ", postorder.data[i]);
    }
    printf("\n");

    // Delete a value from the BST
    int deleteVal = 7;
    root = bstDelete(root, deleteVal);
    printf("\nBST structure after deleting value %d:\n", deleteVal);
    bstPrint(root, 0);

    // Free the BST
    bstFree(root);

    return 0;
}
