trait IntExt {
    fn meters(self) -> Self;
}

impl IntExt for usize {
    fn meters(self) -> Self {
        self * 1 // ???
    }
}
