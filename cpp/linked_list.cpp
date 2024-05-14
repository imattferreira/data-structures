#include <stdio.h>
#include <iostream>

using namespace std;

struct Node
{
public:
  int value;
  Node *next = nullptr;

  Node(int v) : value(v) {}
};

struct LinkedList
{
private:
  int size = 0;
  Node *head = nullptr;

public:
  void insert(int value)
  {
    Node *tmp = this->head;
    Node *node = new Node(value);

    do
    {
      if (this->head == nullptr)
      {
        this->head = node;
        break;
      }

      if (tmp->next == nullptr)
      {
        tmp->next = node;
        break;
      }

      tmp = tmp->next;
    } while (tmp != nullptr);
  }

  void insertAtBeginning(int value)
  {
    Node *node = new Node(value);

    node->next = this->head;
    this->head = node;
  }

  void pop()
  {
    Node *prev = nullptr;
    Node *curr = this->head;

    while (curr != nullptr)
    {
      if (curr->next == nullptr)
      {
        prev->next = nullptr;
        delete curr;
        break;
      }

      prev = curr;
      curr = curr->next;
    }
  }

  void show()
  {
    Node *tmp = this->head;

    while (tmp != nullptr)
    {
      cout << tmp->value << endl;
      tmp = tmp->next;
    }
  }
};

int main()
{
  LinkedList *list = new LinkedList();

  list->insert(10);
  list->insertAtBeginning(5);
  list->insert(20);
  list->insert(30);
  list->insertAtBeginning(2);
  list->insert(40);
  list->pop();
  list->show();

  return 0;
}
