mod target;
pub use target::Target;

use crate::Element;

pub struct A {
    href: String,
    target: Option<Target>,
}

impl A {
    pub fn new(href: impl Into<String>) -> Self {
        Self {
            href: href.into(),
            target: None,
        }
    }

    pub fn target(mut self, target: Target) -> Self {
        self.target = Some(target);
        self
    }
}

impl Element for A {
    fn render() {}
}
