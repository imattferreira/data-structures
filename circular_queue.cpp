#include <stdio.h>
#include <iostream>

using namespace std;

struct CircularQueue {
  private:
    unsigned capacity;
    unsigned int front_index;
    int* items;
    unsigned size = 0;
    unsigned int tail_index;

    bool is_last(int i) {
      return i == this->capacity - 1;
    }

  public:
    CircularQueue(unsigned int c): capacity(c) {
      this->items = new int[c];
    }

    void clear() {
      while (this->size != 0) {
        this->front_index = this->is_last(this->front_index) ? 0 : this->front_index + 1;
        this->items[this->front_index] = NULL;
        this->size--;
      }
    }

    void dequeue() {
      if (this->empty()) {
        throw runtime_error("[CircularQueue]: is empty!");
      }

      this->front_index = this->is_last(this->front_index) ? 0 : this->front_index + 1;
      this->size--;
    }

    bool empty() {
      return this->size == 0;
    }

    void enqueue(int item) {
      if (this->full()) {
        throw runtime_error("[CircularQueue]: is full!");
      }

      this->tail_index = this->is_last(this->tail_index) ? 0 : this->tail_index--;
      this->items[this->tail_index] = item;
      this->size++;
    }

    int front() {
      return this->items[this->front_index];
    }

    bool full() {
      return this->size == this->capacity;
    }

    void show() {
      int i = this->front_index;

      while (i != this->tail_index) {
        cout << "Item: " << this->items[i] << " | Position: " << i << endl;

        i = this->is_last(i) ? 0 : i + 1;
      }
    }

    int size() {
      return this->size;
    }

    int tail() {
      return this->items[this->tail_index];
    }
};

int main() {
  return 0;
}
