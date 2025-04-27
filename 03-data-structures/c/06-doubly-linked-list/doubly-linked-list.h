#ifndef DOUBLY_LINKED_LIST_H
#define DOUBLY_LINKED_LIST_H

typedef struct Node
{
    int value;
    struct Node *next;
    struct Node *prev;
} Node;

typedef struct DoublyLinkedList
{
    Node *head;
    Node *tail;
    int size;
} DoublyLinkedList;

Node *createNode(int value);
void doublyLinkedListInit(DoublyLinkedList *list);
void doublyLinkedListFree(DoublyLinkedList *list);
Node *doublyLinkedListGetAtIndex(DoublyLinkedList *list, int index);
int doublyLinkedListSearch(DoublyLinkedList *list, int value);
void doublyLinkedListPrepend(DoublyLinkedList *list, Node *node);
void doublyLinkedListAppend(DoublyLinkedList *list, Node *node);
void doublyLinkedListInsertAtIndex(DoublyLinkedList *list, int index, Node *node);
void doublyLinkedListDeleteNode(DoublyLinkedList *list, Node *node); // Updated function
void doublyLinkedListPrint(const DoublyLinkedList *list);

#endif // DOUBLY_LINKED_LIST_H
