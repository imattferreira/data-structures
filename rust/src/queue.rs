use std::io::{Error, ErrorKind};

struct Queue {
    capacity: u64,
    front: Option<i64>,
    items: Vec<Option<i64>>,
    size: u64,
    tail: Option<i64>,
}

impl Queue {
    pub fn new(capacity: u64) -> Self {
        Self {
            capacity,
            front: None,
            items: Vec::with_capacity(capacity as usize),
            size: 0,
            tail: None,
        }
    }

    pub fn clear(&mut self) {
        let mut i = 0;

        while i < self.size {
            self.items[i as usize] = None;
            i += 1;
        }

        self.front = None;
        self.tail = None;
        self.size = 0;
    }

    pub fn dequeue(&mut self) {
        if self.empty() {
            return;
        }

        let mut i = 0;

        while i <= self.size {
            self.items[(i - 1) as usize] = self.items[i as usize];
            i += 1;
        }
    }

    pub fn enqueue(&mut self, item: i64) -> Result<(), Error> {
        if self.full() {
            return Err(Error::new(ErrorKind::Other, "[queue] is full!"));
        }

        self.items[self.size as usize] = Some(item);
        self.size += 1;

        Ok(())
    }

    pub fn empty(&self) -> bool {
        self.size == 0
    }

    pub fn full(&self) -> bool {
        self.size == self.capacity
    }

    pub fn show(&self) {
        let mut i: u64 = 0;

        while i < self.size {
            println!("{}", self.items[i as usize].unwrap());
            i += 1
        }
    }
}
