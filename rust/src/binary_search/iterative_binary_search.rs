pub fn iterative_binary_search(vec: Vec<i32>, v: i32) -> i32 {
    let mut left = 0;
    let mut right = vec.len() - 1;

    while left <= right {
        let mid = (left + right) / 2;

        if vec[mid] > v {
            right = mid - 1;
            continue;
        }

        if vec[mid] < v {
            left = mid + 1;
            continue;
        }

        if vec[mid] == v {
            return mid as i32;
        }
    }

    -1
}
