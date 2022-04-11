#![allow(incomplete_features)]
#![feature(generic_const_exprs)]

macro_rules! bounded_impl {
    ($name:ident($ty:ty)) => {
        pub struct $name<const MIN: $ty, const MAX: $ty>($ty, ::guardian::Guard::<true>);

        impl<const MIN: $ty, const MAX: $ty> $name<MIN, MAX> where ::guardian::Guard::<{ MIN < MAX }>: ::guardian::True {
            #[must_use]
            pub fn new(value: $ty) -> Option<Self> {
                if value < MIN || value > MAX {
                    None
                } else {
                    Some(Self(value, ::guardian::Guard))
                }
            }

            #[must_use]
            #[allow(clippy::missing_safety_doc)]
            pub unsafe fn new_unchecked(value: $ty) -> Self {
                Self(value, ::guardian::Guard)
            }

            #[must_use]
            pub fn get(self) -> $ty {
                self.0
            }
        }
    };
    ($name:ident($ty:ty), $($names:ident($tys:ty)),+) => {
        bounded_impl!($name($ty));
        bounded_impl!($($names($tys)),+);
    }
}

bounded_impl! {
    BoundedUsize(usize),
    BoundedIsize(isize),
    BoundedU8(u8),
    BoundedI8(i8),
    BoundedU16(u16),
    BoundedI16(i16),
    BoundedU32(u32),
    BoundedI32(i32),
    BoundedU64(u64),
    BoundedI64(i64),
    BoundedU128(u128),
    BoundedI12(i128)
}
