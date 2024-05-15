use std::io::{Error, ErrorKind};

struct CircularQueue {
    capacity: u64,
    front_index: u64,
    tail_index: u64,
    size: u64,
    items: Vec<Option<i64>>,
}

impl CircularQueue {
    pub fn new(capacity: u64) -> Self {
        Self {
            capacity,
            front_index: 0,
            tail_index: 0,
            size: 0,
            items: Vec::with_capacity(capacity as usize),
        }
    }

    pub fn clear(&mut self) {
        let mut i = self.front_index;

        while i != self.tail_index {
            self.items[i as usize] = None;
            i = self.nextIndex(i);
        }
    }

    pub fn dequeue(&mut self) {
        if self.empty() {
            return;
        }

        self.front_index = self.nextIndex(self.front_index);
        self.size -= 1;
    }

    pub fn enqueue(&mut self, item: i64) -> Result<(), Error> {
        if self.full() {
            return Err(Error::new(ErrorKind::Other, "[queue] is full!"));
        }

        self.tail_index = self.nextIndex(self.tail_index);
        self.items[self.tail_index as usize] = Some(item);
        self.size += 1;

        Ok(())
    }

    pub fn front(&self) -> Option<i64> {
        self.items[self.front_index as usize]
    }

    pub fn tail(&self) -> Option<i64> {
        self.items[self.tail_index as usize]
    }

    pub fn show(&self) {
        let mut i = self.front_index;

        while i != self.tail_index {
            println!("{}", self.items[i as usize].unwrap());
            i = self.nextIndex(i);
        }
    }

    pub fn nextIndex(&self, index: u64) -> u64 {
        if index == self.capacity - 1 {
            return 0;
        }

        index + 1
    }

    pub fn full(&self) -> bool {
        self.size == self.capacity
    }

    pub fn empty(&self) -> bool {
        self.size == 0
    }
}
