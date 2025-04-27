#ifndef QUEUE_H
#define QUEUE_H

#include "linked-list.h"
#include <stdbool.h>

typedef struct Queue
{
    LinkedList list;
    Node *tail;
} Queue;

// Queue operations
void queueInit(Queue *queue);
void queueFree(Queue *queue);
void enqueue(Queue *queue, int value);
Node *dequeue(Queue *queue);
Node *queuePeek(Queue *queue);
bool queueIsEmpty(Queue *queue);
int queueSize(Queue *queue);

#endif
