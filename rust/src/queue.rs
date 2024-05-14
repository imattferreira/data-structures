use std::io::{Error, ErrorKind};

struct Queue {
    capacity: i64,
    front: Option<i64>,
    items: Vec<Option<i64>>,
    size: i64,
    tail: Option<i64>,
}

impl Queue {
    pub fn new(capacity: i64) -> Self {
        Self {
            capacity,
            front: None,
            items: Vec::with_capacity(capacity as usize),
            size: 0,
            tail: None,
        }
    }

    pub fn clear(&mut self) {
        let mut i = self.last_index();

        while i != 0 {
            self.items[i] = None;

            i -= 1;
        }

        self.size = 0;
        self.front = None;
        self.tail = None;
    }

    pub fn empty(&self) -> bool {
        self.size == 0
    }

    fn last_index(&self) -> usize {
        (self.size - 1) as usize
    }

    pub fn dequeue(&mut self) -> Result<(), Error> {
        if self.empty() {
            return Err(Error::new(ErrorKind::Other, "[queue] is empty!"));
        }

        let last_i = self.last_index();

        self.items[last_i] = None;
        self.size -= 1;
        self.tail = self.items[last_i - 1];

        Ok(())
    }

    pub fn enqueue(&mut self, item: i64) -> Result<(), Error> {
        if self.full() {
            return Err(Error::new(ErrorKind::Other, "[queue] is full!"));
        }

        let index = self.last_index() + 1;

        self.items[index] = Some(item);
        self.tail = Some(item);
        self.size += 1;

        Ok(())
    }

    pub fn front(&self) -> Option<i64> {
        self.front
    }

    pub fn full(&self) -> bool {
        self.size == self.capacity
    }

    pub fn show(&self) {
        let mut i = self.last_index();

        while i != 0 {
            let item = self.items[i].unwrap();

            println!("{}", item);

            i -= 1;
        }
    }

    pub fn size(&self) -> i64 {
        self.size
    }

    pub fn tail(&self) -> Option<i64> {
        self.tail
    }
}
