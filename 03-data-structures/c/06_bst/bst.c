#include <stdio.h>
#include <stdlib.h>
#include "bst.h"

// -------------------- Creation & Destruction --------------------

BSTNode *bst_create(int val)
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

void bst_free(BSTNode *root)
{
    if (root == NULL)
        return;
    bst_free(root->left);
    bst_free(root->right);
    free(root);
}

// -------------------- Access --------------------

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 * Returns NULL if the tree is empty.
 */
int *bst_get_min(BSTNode *root)
{
    if (root == NULL)
        return NULL;
    if (root->left == NULL)
        return &root->val;
    return bst_get_min(root->left);
}

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 * Returns NULL if the tree is empty.
 */
int *bst_get_max(BSTNode *root)
{
    if (root == NULL)
        return NULL;
    if (root->right == NULL)
        return &root->val;
    return bst_get_max(root->right);
}

// -------------------- Search --------------------

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 */
bool bst_exists(BSTNode *root, int val)
{
    if (root == NULL)
        return false;
    if (val == root->val)
        return true;
    if (val < root->val)
        return bst_exists(root->left, val);
    return bst_exists(root->right, val);
}

// -------------------- Insertion --------------------

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 */
void bst_insert(BSTNode *root, int val)
{
    if (root == NULL)
        return;

    if (root->val == val)
        return;

    if (val < root->val)
    {
        if (root->left)
            bst_insert(root->left, val);
        else
            root->left = bst_create(val);
    }
    else
    {
        if (root->right)
            bst_insert(root->right, val);
        else
            root->right = bst_create(val);
    }
}

// -------------------- Deletion --------------------

/*
 * Time Complexity: O(log n) average, O(n) worst-case
 */
BSTNode *bst_delete(BSTNode *root, int val)
{
    if (root == NULL)
        return NULL;

    if (val < root->val)
    {
        root->left = bst_delete(root->left, val);
    }
    else if (val > root->val)
    {
        root->right = bst_delete(root->right, val);
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
        root->right = bst_delete(root->right, minLargerNode->val);
    }

    return root;
}

// -------------------- Other operations --------------------

/*
 * Time Complexity: O(n)
 */
int bst_height(BSTNode *root)
{
    if (root == NULL)
        return 0;

    int lh = bst_height(root->left);
    int rh = bst_height(root->right);

    return (lh > rh ? lh : rh) + 1;
}

/*
 * Time Complexity: O(n)
 */
void bst_preorder_traversal(BSTNode *root, int visited[], int *size)
{
    if (root == NULL)
        return;
    visited[(*size)++] = root->val;
    bst_preorder_traversal(root->left, visited, size);
    bst_preorder_traversal(root->right, visited, size);
}

/*
 * Time Complexity: O(n)
 */
void bst_inorder_traversal(BSTNode *root, int visited[], int *size)
{
    if (root == NULL)
        return;
    bst_inorder_traversal(root->left, visited, size);
    visited[(*size)++] = root->val;
    bst_inorder_traversal(root->right, visited, size);
}

/*
 * Time Complexity: O(n)
 */
void bst_postorder_traversal(BSTNode *root, int visited[], int *size)
{
    if (root == NULL)
        return;
    bst_postorder_traversal(root->left, visited, size);
    bst_postorder_traversal(root->right, visited, size);
    visited[(*size)++] = root->val;
}

/*
 * Time Complexity: O(n)
 */
void bst_print(BSTNode *root, int level)
{
    if (root == NULL)
        return;
    bst_print(root->right, level + 1);
    printf("%*s%d\n", level * 4, "", root->val);
    bst_print(root->left, level + 1);
}
