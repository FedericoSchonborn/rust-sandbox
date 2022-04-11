#![warn(clippy::pedantic, clippy::cargo)]

macro_rules! bounded_impl {
    ($ty:ty => $name:ident) => {

        pub struct $name<const MAX: $ty, const MIN: $ty>($ty);

        impl<const MAX: $ty, const MIN: $ty> $name<MAX, MIN> {
            #[must_use]
            pub fn new(value: $ty) -> Option<Self> {
                if value > MAX && value < MIN {
                    Some(Self(value))
                } else {
                    None
                }
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
    ($ty:ty => $name:ident, $($tys:ty => $names:ident),+) => {
        bounded_impl!($ty => $name);
        bounded_impl!($($tys => $names),+);
    }
}

bounded_impl! {
    usize => BoundedUsize,
    isize => BoundedIsize,
    u8 => BoundedU8,
    i8 => BoundedI8,
    u16 => BoundedU16,
    i16 => BoundedI16,
    u32 => BoundedU32,
    i32 => BoundedI32,
    u64 => BoundedU64,
    i64 => BoundedI64,
    u128 => BoundedU128,
    i128 => BoundedI128
}
