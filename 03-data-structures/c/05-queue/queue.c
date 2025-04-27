#include <stdio.h>
#include <stdlib.h>
#include "queue.h"

// Initializes the queue (O(1))
void queueInit(Queue *queue)
{
    linkedListInit(&queue->list); // Initialize the internal linked list
    queue->tail = NULL;           // Set the tail to NULL initially
}

// Frees the queue (O(n))
void queueFree(Queue *queue)
{
    linkedListFree(&queue->list); // Free the internal linked list
    queue->tail = NULL;           // Reset the tail pointer to NULL
}

// Enqueues an element to the queue (add to tail) (O(1))
void enqueue(Queue *queue, int value)
{
    // Create a new node for the value
    Node *newNode = createNode(value);
    if (!newNode)
    {
        return; // Allocation failed
    }

    // If the queue is empty, both head and tail should point to the new node
    if (queueIsEmpty(queue))
    {
        queue->list.head = newNode; // Set head to the new node
        queue->tail = newNode;      // Set tail to the new node
    }
    else
    {
        // If the queue is not empty, link the new node after the current tail
        queue->tail->next = newNode; // Link the new node after the current tail
        queue->tail = newNode;       // Update the tail to the new node
    }

    // Increment the size of the queue
    queue->list.size++;
}

// Dequeues an element from the queue (remove from head) (O(1))
Node *dequeue(Queue *queue)
{
    if (queueIsEmpty(queue))
    {
        fprintf(stderr, "Queue is empty, cannot dequeue\n");
        return NULL; // Error value (NULL for Node)
    }

    // Get the front element (head) of the queue
    Node *nodeToReturn = queue->list.head;

    // Remove the element from the linked list (O(1))
    queue->list.head = queue->list.head->next;

    // Decrease size and reset tail if necessary
    queue->list.size--;
    if (queue->list.size == 0)
    {
        queue->tail = NULL; // Reset tail to NULL if the queue is empty
    }

    // Return the dequeued node
    nodeToReturn->next = NULL; // Ensure the dequeued node doesn't link to anything else
    return nodeToReturn;       // Return the node that was dequeued
}

// Peek the front element without removing it (O(1))
Node *queuePeek(Queue *queue)
{
    if (queueIsEmpty(queue))
    {
        fprintf(stderr, "Queue is empty, cannot peek\n");
        return NULL; // Error value (NULL for Node)
    }

    return queue->list.head; // Return the front node of the queue
}

// Checks if the queue is empty (O(1))
bool queueIsEmpty(Queue *queue)
{
    return queue->list.size == 0; // If size is 0, the queue is empty
}

// Get the size of the queue (O(1))
int queueSize(Queue *queue)
{
    return queue->list.size; // Return the size of the queue
}
