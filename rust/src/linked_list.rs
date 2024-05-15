use std::ptr;

struct Node {
    value: i64,
    next: *mut Node,
}

impl Node {
    pub fn new(value: i64) -> Self {
        Self {
            value,
            next: ptr::null_mut(),
        }
    }
}

struct LinkedList {
    size: u64,
    head: *mut Node,
}

impl LinkedList {
    pub fn new() -> Self {
        Self {
            size: 0,
            head: ptr::null_mut(),
        }
    }

    pub fn insert(&mut self, value: i64) {
        let node = Box::into_raw(Box::new(Node::new(value)));
        self.size += 1;

        if self.head.is_null() {
            self.head = node;
            return;
        }

        let mut tmp = self.head;

        unsafe {
            while !tmp.is_null() {
                if (*tmp).next.is_null() {
                    (*tmp).next = node;
                    break;
                }

                tmp = (*tmp).next;
            }
        }
    }

    pub fn insertAtBeginning(&mut self, value: i64) {
        let node = Box::into_raw(Box::new(Node::new(value)));
        self.size += 1;

        let mut tmp = self.head;

        unsafe {
            while !tmp.is_null() {
                if (*tmp).next.is_null() {
                    (*tmp).next = node;
                    break;
                }

                tmp = (*tmp).next;
            }
        }
    }

    pub fn pop(&mut self) {
        if self.head.is_null() {
            return;
        }

        let mut previous = self.head;
        let mut current = self.head;

        unsafe {
            while !current.is_null() {
                if (*current).next.is_null() {
                    (*previous).next = ptr::null_mut();
                    break;
                }

                previous = current;
                current = (*current).next;
            }
        }
    }

    pub fn show(&self) {
        let mut tmp = self.head;

        unsafe {
            while !tmp.is_null() {
                println!("{}", (*tmp).value);
                tmp = (*tmp).next;
            }
        }
    }
}
