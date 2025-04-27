#include <stdio.h>
#include <stdlib.h>
#include "queue.h"

int main()
{
    Queue queue;
    queueInit(&queue); // Initialize the queue

    // Enqueue some elements
    printf("Enqueuing 10, 20, and 30...\n");
    enqueue(&queue, 10);
    enqueue(&queue, 20);
    enqueue(&queue, 30);

    // Peek the front element
    Node *front = queuePeek(&queue);
    if (front != NULL)
    {
        printf("Front element (peek): %d\n", front->value);
    }

    // Get the size of the queue
    printf("Queue size: %d\n", queueSize(&queue));

    // Dequeue elements
    printf("\nDequeuing elements...\n");
    Node *dequeued = dequeue(&queue);
    if (dequeued != NULL)
    {
        printf("Dequeued: %d\n", dequeued->value);
        free(dequeued); // Don't forget to free the dequeued node
    }

    dequeued = dequeue(&queue);
    if (dequeued != NULL)
    {
        printf("Dequeued: %d\n", dequeued->value);
        free(dequeued);
    }

    // Check queue size again
    printf("Queue size after dequeueing: %d\n", queueSize(&queue));

    // Peek again
    front = queuePeek(&queue);
    if (front != NULL)
    {
        printf("Front element (peek) after dequeue: %d\n", front->value);
    }

    // Free the queue
    queueFree(&queue);

    return 0;
}
