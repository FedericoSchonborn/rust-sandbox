#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct BoundedUsize<const S: usize, const E: usize>(usize);

impl<const S: usize, const E: usize> BoundedUsize<S, E> {
    pub fn new(value: usize) -> Option<Self> {
        if value > S && value < E {
            Some(Self(value))
        } else {
            None
        }
    }

    #[allow(clippy::missing_safety_doc)]
    pub unsafe fn new_unchecked(value: usize) -> Self {
        Self(value)
    }

    pub fn get(self) -> usize {
        self.0
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct BoundedIsize<const S: isize, const E: isize>(isize);

impl<const S: isize, const E: isize> BoundedIsize<S, E> {
    pub fn new(value: isize) -> Option<Self> {
        if value > S && value < E {
            Some(Self(value))
        } else {
            None
        }
    }

    #[allow(clippy::missing_safety_doc)]
    pub unsafe fn new_unchecked(value: isize) -> Self {
        Self(value)
    }

    pub fn get(self) -> isize {
        self.0
    }
}
