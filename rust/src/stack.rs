use std::io::{Error, ErrorKind};

struct Stack {
    capacity: i64,
    items: Vec<Option<i64>>,
    size: i64,
}

impl Stack {
    pub fn new(capacity: i64) -> Self {
        Self {
            capacity,
            items: Vec::with_capacity(capacity as usize),
            size: 0,
        }
    }

    pub fn clear(&mut self) {
        let mut i = self.last_index();

        while self.size != 0 {
            self.items[i] = None;
            i -= 1;
        }

        self.size = 0;
    }

    pub fn empty(&self) -> bool {
        self.size == 0
    }

    fn last_index(&self) -> usize {
        (self.size - 1) as usize
    }

    pub fn full(&self) -> bool {
        self.size == self.capacity
    }

    pub fn size(&self) -> i64 {
        self.size
    }

    pub fn top(&self) -> Option<i64> {
        self.items[self.last_index()]
    }

    pub fn show(&self) {
        let mut i = self.last_index();

        while i != 0 {
            let item = self.items[i].unwrap();

            println!("{}", item);

            i -= 1;
        }
    }

    pub fn insert(&mut self, item: i64) -> Result<(), Error> {
        if self.full() {
            return Err(Error::new(ErrorKind::Other, "[stack] is full!"));
        }

        let index = self.last_index() + 1;

        self.items[index] = Some(item);
        self.size += 1;

        Ok(())
    }

    pub fn pop(&mut self) -> Result<i64, Error> {
        if self.empty() {
            return Err(Error::new(ErrorKind::Other, "[stack] is empty!"));
        }

        let item = self.items[self.last_index()].unwrap();
        let index = self.last_index();

        self.items[index] = None;
        self.size -= 1;

        Ok(item)
    }
}
