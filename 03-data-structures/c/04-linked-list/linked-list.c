#include <stdio.h>
#include <stdlib.h>
#include "linked-list.h"

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
    return newNode;
}

// Initializes a linked list (O(1))
void linkedListInit(LinkedList *list)
{
    list->head = NULL;
    list->size = 0;
}

// Frees all nodes in the linked list (O(n))
void linkedListFree(LinkedList *list)
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
    list->size = 0;
}

// Returns the node at the given index (O(n))
Node *linkedListGetAtIndex(LinkedList *list, int index)
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
int linkedListSearch(LinkedList *list, int value)
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
void linkedListPrepend(LinkedList *list, Node *node)
{
    if (!list || !node)
        return; // Check if list or node is NULL

    node->next = list->head;
    list->head = node;
    list->size++;
}

// Inserts a node at the tail (O(n))
void linkedListAppend(LinkedList *list, Node *node)
{
    if (!list || !node)
        return; // Check if list or node is NULL

    linkedListInsertAtIndex(list, list->size, node);
}

// Inserts a node at a given index (O(n))
void linkedListInsertAtIndex(LinkedList *list, int index, Node *node)
{
    if (!list || !node || index < 0 || index > list->size)
    {
        fprintf(stderr, "Index out of bounds\n");
        return; // Return without modifying the list
    }

    if (index == 0)
    {
        linkedListPrepend(list, node);
        return;
    }

    Node *current = list->head;
    for (int i = 0; i < index - 1; i++)
    {
        current = current->next;
    }
    node->next = current->next;
    current->next = node;
    list->size++;
}

// Inserts a node after the given node (O(1))
void linkedListInsertAfterNode(LinkedList *list, Node *prevNode, Node *node)
{
    if (!prevNode || !node)
    {
        fprintf(stderr, "Previous node or node cannot be NULL\n");
        return;
    }

    node->next = prevNode->next;
    prevNode->next = node;

    list->size++;
}

// Deletes a node at a given index (O(n))
void linkedListDeleteAtIndex(LinkedList *list, int index)
{
    if (!list || index < 0 || index >= list->size)
    {
        fprintf(stderr, "Index out of bounds\n");
        return; // Return without modifying the list
    }

    Node *toDelete;

    if (index == 0)
    {
        toDelete = list->head;
        list->head = list->head->next;
    }
    else
    {
        Node *current = list->head;
        for (int i = 0; i < index - 1; i++)
        {
            current = current->next;
        }
        toDelete = current->next;
        current->next = toDelete->next;
    }

    free(toDelete);
    list->size--;
}

// Deletes a specific node (O(1))
void linkedListDeleteNode(LinkedList *list, Node *nodeToDelete)
{
    if (!list || !nodeToDelete)
    {
        fprintf(stderr, "List or node to delete cannot be NULL\n");
        return;
    }

    // Special case: If the node to delete is the head
    if (nodeToDelete == list->head)
    {
        list->head = list->head->next;
    }
    else
    {
        // Find the previous node
        Node *prevNode = list->head;
        while (prevNode && prevNode->next != nodeToDelete)
        {
            prevNode = prevNode->next;
        }

        if (!prevNode)
        {
            fprintf(stderr, "Node not found in the list\n");
            return; // Handle error gracefully
        }

        // Adjust pointers to bypass nodeToDelete
        prevNode->next = nodeToDelete->next;
    }

    free(nodeToDelete); // Free the node
    list->size--;
}

// Prints the linked list (O(n))
void linkedListPrint(const LinkedList *list)
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
