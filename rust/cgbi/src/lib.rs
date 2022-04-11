#![warn(clippy::pedantic, clippy::cargo)]

#[cfg_attr(not(feature = "unstable"), path = "stable.rs")]
#[cfg_attr(feature = "unstable", path = "unstable.rs")]
mod _impl;

pub use _impl::*;
