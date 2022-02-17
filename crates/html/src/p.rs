use crate::Element;

pub struct P {}

impl P {
    pub fn new() -> Self {
        Self {}
    }
}

impl Element for P {
    fn render() {}
}
