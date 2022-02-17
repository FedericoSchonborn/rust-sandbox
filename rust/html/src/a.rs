mod target;
pub use target::Target;

use crate::Element;

#[derive(Debug, Default)]
pub struct A {
    _href: String,
    target: Option<Target>,
    children: Option<Vec<Box<dyn Element>>>,
}

impl A {
    pub fn new(href: impl Into<String>) -> Self {
        Self {
            _href: href.into(),
            target: None,
            children: None,
        }
    }

    pub fn target(mut self, target: Target) -> Self {
        self.target = Some(target);
        self
    }

    pub fn children(mut self, children: Vec<Box<dyn Element>>) -> Self {
        self.children = Some(children);
        self
    }
}

impl Element for A {
    fn render(&self) {}
}
