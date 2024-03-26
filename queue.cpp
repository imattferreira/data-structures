#include <stdio.h>
#include <iostream>

using namespace std;

struct Queue {
  private:
    unsigned capacity;
    int front;
    int* items;
    unsigned int size = 0;
    int tail;

  public:
    Queue(unsigned int c): capacity(c) {
      this->items = new int[c];
    }

    void clear() {
      for (int i = 0; i < this->size; i++) {
        this->items[i] = NULL;

        for (int j = i; j < this->size; j++) {
          this->items[j] = this->items[j + 1];
        }
      }
    }

    int dequeue() {
      if (this->empty()) {
        return NULL;
      }

      for (int i = 0; i < this->size; i++) {
        this->items[i] = this->items[i + 1];
      }

      this->front = this->items[0];

      return this->front;
    }

    bool empty() {
      return this->size == 0;
    }

    void enqueue(int item) {
      if (this->full()) {
        throw runtime_error("[Queue]: is full!");
      }

      this->tail = item;
      this->items[this->size] = item;
      this->size++;
    }

    int front() {
      return this->front;
    }

    bool full() {
      return this->size == this->capacity;
    }

    void show() {
      for (int i = 0; i < this->size; i++) {
        cout << "Item: " << this->items[i] << " | Position: " << i << endl;
      }
    }

    int size() {
      return this->size;
    }

    int tail() {
      return this->tail;
    }
};

int main() {
  return 0;
}
