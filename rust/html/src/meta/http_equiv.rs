#[derive(Debug)]
pub enum HttpEquiv {
    ContentSecurityPolicy(String),
    ContentType(String),
    DefaultStyle(String),
    XUaCompatible(String),
    Refresh(String),
}
