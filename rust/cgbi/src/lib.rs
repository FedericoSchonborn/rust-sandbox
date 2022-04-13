#![warn(clippy::pedantic, clippy::cargo)]
#![cfg_attr(feature = "guard", allow(incomplete_features))]
#![cfg_attr(feature = "guard", feature(generic_const_exprs))]

#[cfg_attr(feature = "guard", path = "guard.rs")]
#[cfg_attr(not(feature = "guard"), path = "no-guard.rs")]
mod _impl;

pub use _impl::*;
