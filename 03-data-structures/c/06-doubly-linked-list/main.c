#include <stdio.h>
#include <stdlib.h>
#include "doubly-linked-list.h"

int main()
{
    DoublyLinkedList list;
    doublyLinkedListInit(&list);

    // Create nodes
    Node *node1 = createNode(10);
    Node *node2 = createNode(20);
    Node *node3 = createNode(30);

    // Prepend nodes
    doublyLinkedListPrepend(&list, node1);
    doublyLinkedListPrepend(&list, node2);

    // Append a node
    doublyLinkedListAppend(&list, node3);

    // Print the list
    printf("Doubly Linked List after operations:\n");
    doublyLinkedListPrint(&list);

    // Search for a node
    int index = doublyLinkedListSearch(&list, 20);
    if (index != -1)
        printf("Node with value 20 found at index: %d\n", index);
    else
        printf("Node with value 20 not found.\n");

    // Delete a node
    doublyLinkedListDeleteNode(&list, node2);
    printf("Doubly Linked List after deleting node 20:\n");
    doublyLinkedListPrint(&list);

    // Clean up
    doublyLinkedListFree(&list);

    return 0;
}
