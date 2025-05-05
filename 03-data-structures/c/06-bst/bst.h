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
BSTNode *bstCreate(int val);
void bstFree(BSTNode *root);

// -------------------- Access --------------------
int *bstGetMin(BSTNode *root);
int *bstGetMax(BSTNode *root);

// -------------------- Search --------------------
bool bstExists(BSTNode *root, int val);

// -------------------- Insertion --------------------
void bstInsert(BSTNode *root, int val);

// -------------------- Deletion --------------------
BSTNode *bstDelete(BSTNode *root, int val);

// -------------------- Other operations --------------------
int bstHeight(BSTNode *root);
DynamicArray bstPreorderTraversal(BSTNode *root);
DynamicArray bstInorderTraversal(BSTNode *root);
DynamicArray bstPostorderTraversal(BSTNode *root);
void bstPrint(BSTNode *root, int level);

#endif // BST_H
