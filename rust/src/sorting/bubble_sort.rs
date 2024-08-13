// Big O = O**2

pub fn bubble_sort(v: &mut Vec<i32>) -> &mut Vec<i32> {
    for i in 0..v.len() {
        // small optimization: iterate until unsorted position
        for j in 0..v.len() - i - 1 {
            if v[j] > v[j + 1] {
                v.swap(j, j + 1);
            }
        }
    }

    v
}
