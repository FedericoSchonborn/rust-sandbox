#![allow(non_camel_case_types)]

use std::{
    fmt::{self, Display, Formatter},
    ops::Shl,
};

pub struct cout;

impl<T> Shl<T> for cout
where
    T: Display,
{
    type Output = Self;

    fn shl(self, rhs: T) -> Self::Output {
        print!("{}", rhs);
        self
    }
}

pub struct endl;

impl Display for endl {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        f.write_str("\n")
    }
}
