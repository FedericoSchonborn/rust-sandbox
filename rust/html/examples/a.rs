use html::{a::Target::Blank, A, P};

fn main() {
    A::new("https://rust-lang.org")
        .target(Blank)
        .children(vec![Box::new(P::new())]);
}