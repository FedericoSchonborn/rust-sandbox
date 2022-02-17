use html::{meta::HttpEquiv::XUaCompatible, Meta};

fn main() {
    Meta::charset("utf-8");
    Meta::http_equiv(XUaCompatible(String::from("IE=edge")));
    Meta::other("generator", "rust-html");
}
