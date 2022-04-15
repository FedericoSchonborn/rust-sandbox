pub trait Boolish {
    fn boolish(&self) -> bool;
}

impl Boolish for bool {
    fn boolish(&self) -> bool {
        *self
    }
}

impl Boolish for str {
    fn boolish(&self) -> bool {
        !self.is_empty()
    }
}

impl Boolish for &str {
    fn boolish(&self) -> bool {
        (*self).boolish()
    }
}

impl Boolish for String {
    fn boolish(&self) -> bool {
        (**self).boolish()
    }
}

impl Boolish for &String {
    fn boolish(&self) -> bool {
        (***self).boolish()
    }
}

macro_rules! int_impl {
    ($ty:ty) => {
        impl $crate::Boolish for $ty {
            fn boolish(&self) -> bool {
                *self != 0
            }
        }
    };
    ($ty:ty, $($tys:ty),+) => {
        int_impl!($ty);
        int_impl!($($tys),+);
    }
}

int_impl![isize, usize, i8, u8, i16, u16, i32, u32, i64, u64, i128, u128];

impl<T> Boolish for Vec<T> {
    fn boolish(&self) -> bool {
        !self.is_empty()
    }
}

impl<T> Boolish for [T] {
    fn boolish(&self) -> bool {
        !self.is_empty()
    }
}

impl<T, const N: usize> Boolish for [T; N] {
    fn boolish(&self) -> bool {
        !self.is_empty()
    }
}

impl<T> Boolish for &[T] {
    fn boolish(&self) -> bool {
        (*self).boolish()
    }
}

impl<T, const N: usize> Boolish for &[T; N] {
    fn boolish(&self) -> bool {
        (*self).boolish()
    }
}
