#include <stdio.h>
#include <stdlib.h>
#include "bst.h"

// -------------------- Creation & Destruction --------------------

BSTNode *bstCreate(int val)
{
    BSTNode *newNode = (BSTNode *)malloc(sizeof(BSTNode));
    if (newNode == NULL)
    {
        fprintf(stderr, "Memory allocation failed\n");
        exit(EXIT_FAILURE);
    }
    newNode->val = val;
    newNode->left = NULL;
    newNode->right = NULL;
    return newNode;
}

void bstFree(BSTNode *root)
{
    if (root == NULL)
        return;
    bstFree(root->left);
    bstFree(root->right);
    free(root);
}

// -------------------- Access --------------------

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 * Returns NULL if the tree is empty.
 */
int *bstGetMin(BSTNode *root)
{
    if (root == NULL)
        return NULL;
    if (root->left == NULL)
        return &root->val;
    return bstGetMin(root->left);
}

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 * Returns NULL if the tree is empty.
 */
int *bstGetMax(BSTNode *root)
{
    if (root == NULL)
        return NULL;
    if (root->right == NULL)
        return &root->val;
    return bstGetMax(root->right);
}

// -------------------- Search --------------------

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 */
bool bstExists(BSTNode *root, int val)
{
    if (root == NULL)
        return false;
    if (val == root->val)
        return true;
    if (val < root->val)
        return bstExists(root->left, val);
    return bstExists(root->right, val);
}

// -------------------- Insertion --------------------

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 */
void bstInsert(BSTNode *root, int val)
{
    if (root == NULL)
        return;

    if (root->val == val)
        return;

    if (val < root->val)
    {
        if (root->left)
            bstInsert(root->left, val);
        else
            root->left = bstCreate(val);
    }
    else
    {
        if (root->right)
            bstInsert(root->right, val);
        else
            root->right = bstCreate(val);
    }
}

// -------------------- Deletion --------------------

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 */
BSTNode *bstDelete(BSTNode *root, int val)
{
    if (root == NULL)
        return NULL;

    if (val < root->val)
    {
        root->left = bstDelete(root->left, val);
    }
    else if (val > root->val)
    {
        root->right = bstDelete(root->right, val);
    }
    else
    {
        if (root->right == NULL)
        {
            BSTNode *temp = root->left;
            free(root);
            return temp;
        }
        if (root->left == NULL)
        {
            BSTNode *temp = root->right;
            free(root);
            return temp;
        }

        BSTNode *minLargerNode = root->right;
        while (minLargerNode->left)
        {
            minLargerNode = minLargerNode->left;
        }

        root->val = minLargerNode->val;
        root->right = bstDelete(root->right, minLargerNode->val);
    }

    return root;
}

// -------------------- Other operations --------------------

/*
 * Time Complexity: O(n)
 */
int bstHeight(BSTNode *root)
{
    if (root == NULL)
        return 0;

    int lh = bstHeight(root->left);
    int rh = bstHeight(root->right);

    return (lh > rh ? lh : rh) + 1;
}

/**
 * Time Complexity: O(n)
 * Performs an inorder traversal and returns the result in a dynamic array.
 */
DynamicArray bstInorderTraversal(BSTNode *root)
{
    DynamicArray visited;
    arrayInit(&visited, 10);

    if (root == NULL)
        return visited;

    DynamicArray left = bstInorderTraversal(root->left);
    for (int i = 0; i < left.size; i++)
        arrayAppend(&visited, left.data[i]);

    arrayAppend(&visited, root->val);

    DynamicArray right = bstInorderTraversal(root->right);
    for (int i = 0; i < right.size; i++)
        arrayAppend(&visited, right.data[i]);

    return visited;
}

/**
 * Time Complexity: O(n)
 * Performs a preorder traversal and returns the result in a dynamic array.
 */
DynamicArray bstPreorderTraversal(BSTNode *root)
{
    DynamicArray visited;
    arrayInit(&visited, 10);

    if (root == NULL)
        return visited;

    arrayAppend(&visited, root->val);

    DynamicArray left = bstPreorderTraversal(root->left);
    for (int i = 0; i < left.size; i++)
        arrayAppend(&visited, left.data[i]);

    DynamicArray right = bstPreorderTraversal(root->right);
    for (int i = 0; i < right.size; i++)
        arrayAppend(&visited, right.data[i]);

    return visited;
}

/**
 * Time Complexity: O(n)
 * Performs a postorder traversal and returns the result in a dynamic array.
 */
DynamicArray bstPostorderTraversal(BSTNode *root)
{
    DynamicArray visited;
    arrayInit(&visited, 10);

    if (root == NULL)
        return visited;

    DynamicArray left = bstPostorderTraversal(root->left);
    for (int i = 0; i < left.size; i++)
        arrayAppend(&visited, left.data[i]);

    DynamicArray right = bstPostorderTraversal(root->right);
    for (int i = 0; i < right.size; i++)
        arrayAppend(&visited, right.data[i]);

    arrayAppend(&visited, root->val);

    return visited;
}

/*
 * Time Complexity: O(n)
 */
void bstPrint(BSTNode *root, int level)
{
    if (root == NULL)
        return;
    bstPrint(root->right, level + 1);
    printf("%*s%d\n", level * 4, "", root->val);
    bstPrint(root->left, level + 1);
}
