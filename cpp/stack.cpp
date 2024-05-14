#include <stdio.h>
#include <iostream>

using namespace std;

struct Stack
{
private:
  unsigned capacity;
  int *items;
  unsigned last_index = 0;
  unsigned size = 0;

public:
  Stack(unsigned int c = 4) : capacity(c)
  {
    this->items = new int[c];
  }

  void clear()
  {
    while (this->size != 0)
    {
      this->items[this->size - 1] = NULL;
    }
  }

  bool empty()
  {
    return this->size == 0;
  }

  bool full()
  {
    return this->size == this->capacity;
  }

  void insert(int item)
  {
    if (this->full())
    {
      throw runtime_error("[Stack]: is full!");
    }

    this->last_index++;
    this->items[this->last_index] = item;
  }

  int pop()
  {
    int last_item = this->items[this->last_index];

    this->items[this->last_index] = NULL;
    this->last_index--;

    return last_index;
  }

  void show()
  {
    int i = this->size - 1;

    while (i != 0)
    {
      cout << this->items[i] << endl;
      i--;
    }
  }

  int size()
  {
    return this->size;
  }

  int top()
  {
    return this->items[this->last_index];
  }
};
