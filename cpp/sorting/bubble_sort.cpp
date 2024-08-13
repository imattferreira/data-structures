#include <stdio.h>
#include <iostream>

using namespace std;

int *bubbleSort(int v[], int size)
{
  bool swapped = false;

  for (int i = 0; i < size - 1; i++)
  {
    /**
     * small optimization: iterate only over i until last index instead
     * of all elems
     */
    for (int j = 0; j < i - 1; j++)
    {
      if (v[j] > v[j + 1])
      {
        int tmp = v[j];
        v[j] = v[j + 1];
        v[j + 1] = tmp;
      }
    }

    // small optimization: consider vector as already sorted
    if (swapped == false)
    {
      break;
    }
  }

  return v;
}

int main()
{
  int arr[] = {10, 2, 5, 15, 18};
  /**
   * divides total array memory size by the first element size to get total
   * of elems in vector
   */
  int size = sizeof(arr) / sizeof(arr[0]);
  int *sorted = bubbleSort(arr, size);

  for (int i = 0; i < size; i++)
  {
    cout << sorted[i] << " ";
  }

  cout << endl;

  return 0;
}
