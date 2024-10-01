pub fn recursive_binary_search(vec: Vec<i32>, v: i32, left: i32, right: i32) -> i32 {
    if left > right {
        return -1;
    }

    let mid = (left + right) / 2;

    if vec[mid as usize] > v {
        return recursive_binary_search(vec, v, left, mid - 1);
    }

    if vec[mid as usize] < v {
        return recursive_binary_search(vec, v, mid + 1, right);
    }

    if vec[mid as usize] == v {
        return mid;
    }

    -1
}
