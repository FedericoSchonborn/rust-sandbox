mod private {
    pub trait Sealed {}
}

pub trait True: private::Sealed {}

pub trait False: private::Sealed {}

pub struct Guard<const __: bool>;

impl<const __: bool> private::Sealed for Guard<__> {}

impl True for Guard<true> {}

impl False for Guard<false> {}
