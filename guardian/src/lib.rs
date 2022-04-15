#![warn(clippy::pedantic, clippy::cargo)]

mod private {
    pub trait Sealed {}

    impl<T> Sealed for T {}
}

pub trait True: private::Sealed {}

pub trait False: private::Sealed {}

#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct If<const C: bool>;

impl False for If<false> {}
impl True for If<true> {}

#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct Not<const C: bool>;

impl True for Not<false> {}
impl False for Not<true> {}

#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct And<const A: bool, const B: bool>;

impl False for And<false, false> {}
impl False for And<true, false> {}
impl False for And<false, true> {}
impl True for And<true, true> {}

#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct Or<const A: bool, const B: bool>;

impl False for Or<false, false> {}
impl True for Or<true, false> {}
impl True for Or<false, true> {}
impl True for Or<true, true> {}

#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct Nand<const A: bool, const B: bool>;

impl True for Nand<false, false> {}
impl True for Nand<true, false> {}
impl True for Nand<false, true> {}
impl False for Nand<true, true> {}

#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct Nor<const A: bool, const B: bool>;

impl True for Nor<false, false> {}
impl False for Nor<true, false> {}
impl False for Nor<false, true> {}
impl False for Nor<true, true> {}

#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct Xor<const A: bool, const B: bool>;

impl False for Xor<false, false> {}
impl True for Xor<true, false> {}
impl True for Xor<false, true> {}
impl False for Xor<true, true> {}

#[derive(Debug, Clone, Copy, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct Xnor<const A: bool, const B: bool>;

impl True for Xnor<false, false> {}
impl False for Xnor<true, false> {}
impl False for Xnor<false, true> {}
impl True for Xnor<true, true> {}
