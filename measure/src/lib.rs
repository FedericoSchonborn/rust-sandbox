#[derive(Debug, Clone, Copy, Default, PartialEq, PartialOrd)]
pub struct Meters(f32);

impl Meters {
    pub fn get(self) -> f32 {
        self.0
    }
}

#[derive(Debug, Clone, Copy, Default, PartialEq, PartialOrd)]
pub struct Kilometers(f32);

impl Kilometers {
    pub fn get(self) -> f32 {
        self.0
    }
}

impl Into<Meters> for Kilometers {
    fn into(self) -> Meters {
        Meters(self.0 * 1000_f32)
    }
}

impl Into<Kilometers> for Meters {
    fn into(self) -> Kilometers {
        Kilometers(self.0 / 1000_f32)
    }
}

pub trait Measure: Sized {
    fn m(self) -> Meters;
    fn km(self) -> Kilometers;
}

impl Measure for usize {
    fn m(self) -> Meters {
        Meters(self as f32)
    }

    fn km(self) -> Kilometers {
        Kilometers(self as f32)
    }
}

impl Measure for f32 {
    fn m(self) -> Meters {
        Meters(self)
    }

    fn km(self) -> Kilometers {
        Kilometers(self)
    }
}
