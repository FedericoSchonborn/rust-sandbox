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
    ConstrainedUsize(usize),
    ConstrainedIsize(isize),
    ConstrainedU8(u8),
    ConstrainedI8(i8),
    ConstrainedU16(u16),
    ConstrainedI16(i16),
    ConstrainedU32(u32),
    ConstrainedI32(i32),
    ConstrainedU64(u64),
    ConstrainedI64(i64),
    ConstrainedU128(u128),
    ConstrainedI128(i128)
];
