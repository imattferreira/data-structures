struct CircularQueue {
    capacity: i64,
    front_index: usize,
    items: Vec<Option<i64>>,
    size: i64,
    tail_index: usize,
}

impl CircularQueue {
    pub fn clear(&mut self) {}

    pub fn dequeue(&mut self) {}

    pub fn empty(&self) -> bool {
        self.size == 0
    }

    pub fn full(&self) -> bool {
        self.size == self.capacity
    }

    pub fn front(&self) -> Option<i64> {
        self.items[self.front_index]
    }

    pub fn tail(&self) -> Option<i64> {
        self.items[self.tail_index]
    }

    pub fn size(&self) -> i64 {
        self.size
    }

    pub fn show() {}
}
