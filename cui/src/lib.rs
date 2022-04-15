#![warn(clippy::pedantic, clippy::cargo)]
#![allow(incomplete_features)]
#![feature(generic_const_exprs)]

macro_rules! cui {
    ($name:ident($ty:ty)) => {
        pub struct $name<const MIN: $ty, const MAX: $ty>($ty);

        impl<const MIN: $ty, const MAX: $ty> $name<MIN, MAX> where ::guardian::Guard::<{ MIN <= MAX }>: ::guardian::True {
            #[must_use]
            pub fn new(value: $ty) -> Option<Self> {
                if value <= MIN || value >= MAX {
                    None
                } else {
                    Some(Self(value))
                }
            }

            #[must_use]
            pub fn from_const<const N: $ty>() -> Self where ::guardian::And<{ N >= MIN }, { N <= MAX }>: ::guardian::True {
                Self(N)
            }

            #[must_use]
            #[allow(clippy::missing_safety_doc)]
            pub unsafe fn new_unchecked(value: $ty) -> Self {
                Self(value)
            }

            #[must_use]
            pub fn get(self) -> $ty {
                self.0
            }
        }
    };
    ($name:ident($ty:ty), $($names:ident($tys:ty)),+) => {
        cui!($name($ty));
        cui!($($names($tys)),+);
    }
}

cui![
    Usize(usize),
    Isize(isize),
    U8(u8),
    I8(i8),
    U16(u16),
    I16(i16),
    U32(u32),
    I32(i32),
    U64(u64),
    I64(i64),
    U128(u128),
    I128(i128)
];
