#ifndef BST_H
#define BST_H

#include <stdbool.h>
#include "dynamic-array.h"

typedef struct BSTNode
{
    int val;
    struct BSTNode *left;
    struct BSTNode *right;
} BSTNode;

// -------------------- Creation & Destruction --------------------
BSTNode *bst_create(int val);
void bst_free(BSTNode *root);

// -------------------- Access --------------------
int *bst_get_min(BSTNode *root);
int *bst_get_max(BSTNode *root);

// -------------------- Search --------------------
bool bst_exists(BSTNode *root, int val);

// -------------------- Insertion --------------------
void bst_insert(BSTNode *root, int val);

// -------------------- Deletion --------------------
BSTNode *bst_delete(BSTNode *root, int val);

// -------------------- Other operations --------------------
int bst_height(BSTNode *root);
void bst_preorder_traversal(BSTNode *root, DynamicArray visited);
void bst_inorder_traversal(BSTNode *root, DynamicArray visited);
void bst_postorder_traversal(BSTNode *root, DynamicArray visited);
void bst_print(BSTNode *root, int level);

#endif // BST_H
