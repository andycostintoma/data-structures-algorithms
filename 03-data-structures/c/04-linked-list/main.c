#include <stdio.h>
#include <stdlib.h>
#include "linked-list.h"

int main()
{
    // Initialize the linked list
    LinkedList list;
    linkedListInit(&list);

    // Create some nodes
    Node *node1 = createNode(10);
    Node *node2 = createNode(20);
    Node *node3 = createNode(30);
    Node *node4 = createNode(40);

    // Prepend nodes
    linkedListPrepend(&list, node1); // List: [10]
    linkedListPrepend(&list, node2); // List: [20, 10]

    // Append nodes
    linkedListAppend(&list, node3); // List: [20, 10, 30]
    linkedListAppend(&list, node4); // List: [20, 10, 30, 40]

    // Print the current list
    printf("Current list: ");
    linkedListPrint(&list); // Output: [20, 10, 30, 40]

    // Delete node at index 1 (10 should be deleted)
    printf("Deleting node at index 1...\n");
    linkedListDeleteAtIndex(&list, 1); // List: [20, 30, 40]

    // Print the list after deletion
    printf("List after deletion: ");
    linkedListPrint(&list); // Output: [20, 30, 40]

    // Delete the head node (20 should be deleted)
    printf("Deleting node at index 0...\n");
    linkedListDeleteAtIndex(&list, 0); // List: [30, 40]

    // Print the list after deletion
    printf("List after deletion: ");
    linkedListPrint(&list); // Output: [30, 40]

    // Free the list
    linkedListFree(&list);

    return 0;
}
