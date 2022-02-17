pub mod a;
#[doc(inline)]
pub use a::A;

pub mod meta;
#[doc(inline)]
pub use meta::Meta;

pub mod p;
#[doc(inline)]
pub use p::P;

pub trait Element {
    fn render();
}
