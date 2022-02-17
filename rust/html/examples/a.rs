use html::{a::Target::Blank, A};

fn main() {
    A::new("https://rust-lang.org").target(Blank);
}
