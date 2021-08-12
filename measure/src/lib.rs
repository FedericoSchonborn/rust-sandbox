use std::ops::Div;

#[derive(Debug, Clone, Copy, Default, PartialEq, PartialOrd)]
pub struct Meters(f64);

impl Meters {
    pub fn get(self) -> f64 {
        self.0
    }
}

#[derive(Debug, Clone, Copy, Default, PartialEq, PartialOrd)]
pub struct Kilometers(f64);

impl Kilometers {
    pub fn get(self) -> f64 {
        self.0
    }
}

impl Into<Kilometers> for Meters
where
    f64: Div,
{
    fn into(self) -> Kilometers {
        Kilometers(self.0 / 100_f64)
    }
}

pub trait Measure: Sized {
    fn m(self) -> Meters;
    fn km(self) -> Kilometers;
}

impl Measure for usize {
    fn m(self) -> Meters {
        Meters(self as f64)
    }

    fn km(self) -> Kilometers {
        Kilometers(self as f64)
    }
}

impl Measure for f64 {
    fn m(self) -> Meters {
        Meters(self)
    }

    fn km(self) -> Kilometers {
        Kilometers(self)
    }
}
