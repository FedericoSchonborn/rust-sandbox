mod http_equiv;
pub use http_equiv::HttpEquiv;

use crate::Element;

#[derive(Debug)]
pub enum Meta {
    Charset(String),
    HttpEquiv(HttpEquiv),
    Other(String, String),
}

impl Meta {
    pub fn charset(charset: impl Into<String>) -> Self {
        Self::Charset(charset.into())
    }

    pub fn http_equiv(http_equiv: HttpEquiv) -> Self {
        Self::HttpEquiv(http_equiv)
    }

    pub fn other(name: impl Into<String>, content: impl Into<String>) -> Self {
        Self::Other(name.into(), content.into())
    }
}

impl Element for Meta {
    fn render(&self) {}
}
