#ifndef LINKED_LIST_H
#define LINKED_LIST_H

typedef struct Node
{
    int value;
    struct Node *next;
} Node;

typedef struct LinkedList
{
    Node *head;
    int size;
} LinkedList;

Node *createNode(int value);
void linkedListInit(LinkedList *list);
void linkedListFree(LinkedList *list);
Node *linkedListGetAtIndex(LinkedList *list, int index);
int linkedListSearch(LinkedList *list, int value);
void linkedListPrepend(LinkedList *list, Node *node);
void linkedListAppend(LinkedList *list, Node *node);
void linkedListInsertAtIndex(LinkedList *list, int index, Node *node);
void linkedListInsertAfterNode(LinkedList *list, Node *prevNode, Node *node);
void linkedListDeleteAtIndex(LinkedList *list, int index);
void linkedListDeleteNodeAfter(LinkedList *list, Node *prevNode);
void linkedListPrint(const LinkedList *list);

#endif // LINKED_LIST_H
