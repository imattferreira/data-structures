use std::ptr;

struct Node {
    value: i64,
    next: *mut Node,
    previous: *mut Node,
}

impl Node {
    pub fn new(value: i64) -> Self {
        Self {
            value,
            next: ptr::null_mut(),
            previous: ptr::null_mut(),
        }
    }
}

struct DoublyLinkedList {
    size: u64,
    head: *mut Node,
}

impl DoublyLinkedList {
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
                    (*node).previous = tmp;
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

        if self.head.is_null() {
            self.head = node;
            return;
        }

        unsafe {
            (*node).next = self.head;
            (*self.head).previous = node;
            self.head = node;
        }
    }

    pub fn pop(&mut self) {
        if self.head.is_null() {
            return;
        }

        self.size -= 1;
        let mut tmp = self.head;

        unsafe {
            while !tmp.is_null() {
                if (*tmp).next.is_null() {
                    (*(*tmp).previous).next = ptr::null_mut();
                    break;
                }

                tmp = (*tmp).next;
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
