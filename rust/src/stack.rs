use std::io::{Error, ErrorKind};

struct Stack {
    capacity: u64,
    items: Vec<Option<i64>>,
    last_index: i64,
}

impl Stack {
    pub fn new(capacity: u64) -> Self {
        Self {
            capacity,
            items: Vec::with_capacity(capacity as usize),
            last_index: -1,
        }
    }

    pub fn clear(&mut self) {
        while self.last_index > -1 {
            self.items[self.last_index as usize] = None;
            self.last_index -= 1;
        }
    }

    pub fn empty(&self) -> bool {
        self.last_index == -1
    }

    pub fn full(&self) -> bool {
        self.size() == self.capacity
    }

    pub fn insert(&mut self, item: i64) -> Result<(), Error> {
        if self.full() {
            return Err(Error::new(ErrorKind::Other, "[stack] is full!"));
        }

        self.last_index += 1;
        self.items[self.last_index as usize] = Some(item);

        Ok(())
    }

    pub fn pop(&mut self) {
        if self.empty() {
            return;
        }

        self.items[self.last_index as usize] = None;
        self.last_index -= 1
    }

    pub fn size(&self) -> u64 {
        (self.last_index + 1) as u64
    }

    pub fn show(&self) {
        let mut i = self.last_index as usize;

        while i != 0 {
            println!("{}", self.items[i].unwrap());
            i -= 1
        }
    }

    pub fn top(&self) -> Option<i64> {
        self.items[self.last_index as usize]
    }
}
