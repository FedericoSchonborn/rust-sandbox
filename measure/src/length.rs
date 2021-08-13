use std::ops::{Deref, DerefMut};

pub trait Length: Sized {
    fn m(self) -> Meters;
    fn km(self) -> Kilometers;
}

impl Length for usize {
    fn m(self) -> Meters {
        Meters(self as f32)
    }

    fn km(self) -> Kilometers {
        Kilometers(self as f32)
    }
}

impl Length for f32 {
    fn m(self) -> Meters {
        Meters(self)
    }

    fn km(self) -> Kilometers {
        Kilometers(self)
    }
}

#[derive(Debug, Clone, Copy, Default, PartialEq, PartialOrd)]
pub struct Meters(f32);

impl Meters {
    pub fn km(self) -> Kilometers {
        Kilometers(self.0 / 1000_f32)
    }
}

impl Deref for Meters {
    type Target = f32;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for Meters {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}

#[derive(Debug, Clone, Copy, Default, PartialEq, PartialOrd)]
pub struct Kilometers(f32);

impl Kilometers {
    pub fn m(self) -> Meters {
        Meters(self.0 * 1000_f32)
    }
}

impl Deref for Kilometers {
    type Target = f32;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for Kilometers {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}
