use crate::Element;

#[derive(Debug, Default)]
pub struct P {}

impl P {
    pub fn new() -> Self {
        Self::default()
    }
}

impl Element for P {
    fn render(&self) {}
}
