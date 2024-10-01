#include <iostream>
#include <vector>

using namespace std;

int recursiveBinarySearch(vector<int> vec, int v, int left, int right)
{
  if (left > right)
  {
    return -1;
  }

  int mid = (left + right) / 2;

  if (vec[mid] > v)
  {
    return recursiveBinarySearch(vec, v, left, mid - 1);
  }

  if (vec[mid] < v)
  {
    return recursiveBinarySearch(vec, v, mid + 1, right);
  }

  if (vec[mid] == v)
  {
    return mid;
  }

  return -1;
}

int main()
{
  vector<int> vec{10, 20, 30, 40, 50, 60};

  cout << recursiveBinarySearch(vec, 40, 0, vec.size() - 1) << endl;

  return 0;
}
