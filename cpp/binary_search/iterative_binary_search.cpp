#include <iostream>
#include <vector>

using namespace std;

int iterativeBinarySearch(vector<int> vec, int v)
{
  int left = 0;
  int right = vec.size() - 1;

  while (left <= right)
  {
    int mid = (left + right) / 2;

    if (vec[mid] > v)
    {
      right = mid - 1;
      continue;
    }

    if (vec[mid] < v)
    {
      left = mid + 1;
      continue;
    }

    if (vec[mid] == v)
    {
      return mid;
    }
  }

  return -1;
}

int main()
{
  vector<int> vec{10, 20, 30, 40, 50, 60};

  cout << iterativeBinarySearch(vec, 40) << endl;

  return 0;
}
