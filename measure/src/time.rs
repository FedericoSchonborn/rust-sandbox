use std::ops::{Deref, DerefMut};

pub trait Time: Sized {
    fn h(self) -> Hours;
    fn m(self) -> Minutes;
}

impl Time for usize {
    fn h(self) -> Hours {
        Hours(self as f32)
    }

    fn m(self) -> Minutes {
        Minutes(self as f32)
    }
}

impl Time for f32 {
    fn h(self) -> Hours {
        Hours(self)
    }

    fn m(self) -> Minutes {
        Minutes(self)
    }
}

#[derive(Debug, Clone, Copy, PartialEq, PartialOrd)]
pub struct Hours(f32);

impl Hours {
    pub fn m(self) -> Minutes {
        Minutes(self.0 * 60_f32)
    }
}

impl Deref for Hours {
    type Target = f32;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for Hours {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}

#[derive(Debug, Clone, Copy, PartialEq, PartialOrd)]
pub struct Minutes(f32);

impl Minutes {
    pub fn h(self) -> Hours {
        Hours(self.0 / 60_f32)
    }
}

impl Deref for Minutes {
    type Target = f32;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for Minutes {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}
