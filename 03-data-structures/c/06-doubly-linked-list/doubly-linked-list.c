#include <stdio.h>
#include <stdlib.h>
#include "doubly-linked-list.h"

// Helper function to handle memory allocation (centralized check)
Node *createNode(int value)
{
    Node *newNode = (Node *)malloc(sizeof(Node));
    if (!newNode)
    {
        fprintf(stderr, "Memory allocation failed\n");
        exit(EXIT_FAILURE); // Ensure that the program halts on allocation failure
    }
    newNode->value = value;
    newNode->next = NULL;
    newNode->prev = NULL;
    return newNode;
}

// Initializes a doubly linked list (O(1))
void doublyLinkedListInit(DoublyLinkedList *list)
{
    list->head = NULL;
    list->tail = NULL;
    list->size = 0;
}

// Frees all nodes in the doubly linked list (O(n))
void doublyLinkedListFree(DoublyLinkedList *list)
{
    if (!list)
        return; // Check if list is NULL
    Node *current = list->head;
    while (current != NULL)
    {
        Node *next = current->next;
        free(current);
        current = next;
    }
    list->head = NULL;
    list->tail = NULL;
    list->size = 0;
}

// Returns the node at the given index (O(n))
Node *doublyLinkedListGetAtIndex(DoublyLinkedList *list, int index)
{
    if (!list || index < 0 || index >= list->size)
    {
        fprintf(stderr, "Index out of bounds\n");
        return NULL; // Return NULL if the index is out of bounds
    }

    Node *current = list->head;
    for (int i = 0; i < index; i++)
    {
        current = current->next;
    }
    return current; // Return the node at the specified index
}

// Searches for a value, returns its index or -1 (O(n))
int doublyLinkedListSearch(DoublyLinkedList *list, int value)
{
    if (!list)
        return -1; // Check if list is NULL

    Node *current = list->head;
    int index = 0;
    while (current != NULL)
    {
        if (current->value == value)
        {
            return index;
        }
        current = current->next;
        index++;
    }
    return -1;
}

// Inserts a node at the head (O(1))
void doublyLinkedListPrepend(DoublyLinkedList *list, Node *node)
{
    if (!list || !node)
        return; // Check if list or node is NULL

    node->next = list->head;
    if (list->head != NULL)
    {
        list->head->prev = node;
    }
    list->head = node;
    if (list->tail == NULL)
    {
        list->tail = node;
    }
    list->size++;
}

// Inserts a node at the tail (O(1))
void doublyLinkedListAppend(DoublyLinkedList *list, Node *node)
{
    if (!list || !node)
        return; // Check if list or node is NULL

    if (list->tail == NULL)
    {
        list->head = node;
        list->tail = node;
    }
    else
    {
        node->prev = list->tail;
        list->tail->next = node;
        list->tail = node;
    }
    list->size++;
}

// Inserts a node at a given index (O(n))
void doublyLinkedListInsertAtIndex(DoublyLinkedList *list, int index, Node *node)
{
    if (!list || !node || index < 0 || index > list->size)
    {
        fprintf(stderr, "Index out of bounds\n");
        return; // Return without modifying the list
    }

    if (index == 0)
    {
        doublyLinkedListPrepend(list, node);
        return;
    }
    else if (index == list->size)
    {
        doublyLinkedListAppend(list, node);
        return;
    }

    Node *current = list->head;
    for (int i = 0; i < index; i++)
    {
        current = current->next;
    }

    node->next = current;
    node->prev = current->prev;
    if (current->prev != NULL)
    {
        current->prev->next = node;
    }
    current->prev = node;

    list->size++;
}

// Deletes a node from the list (O(1))
void doublyLinkedListDeleteNode(DoublyLinkedList *list, Node *node)
{
    if (!list || !node)
    {
        fprintf(stderr, "List or node is NULL\n");
        return;
    }

    if (node == list->head)
    {
        list->head = node->next;
        if (list->head != NULL)
        {
            list->head->prev = NULL;
        }
    }
    else if (node == list->tail)
    {
        list->tail = node->prev;
        if (list->tail != NULL)
        {
            list->tail->next = NULL;
        }
    }
    else
    {
        node->prev->next = node->next;
        node->next->prev = node->prev;
    }

    free(node);
    list->size--;
}

// Prints the doubly linked list (O(n))
void doublyLinkedListPrint(const DoublyLinkedList *list)
{
    if (!list)
        return; // Check if list is NULL

    Node *current = list->head;
    printf("[");
    while (current != NULL)
    {
        printf("%d", current->value);
        if (current->next != NULL)
        {
            printf(", ");
        }
        current = current->next;
    }
    printf("]\n");
}
