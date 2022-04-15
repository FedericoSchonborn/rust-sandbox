use cui::ConstrainedU8;

// This type only accepts ASCII characters ranging from 'a' to 'z'.
type AsciiLowerAlpha = ConstrainedU8<0x61, 0x7A>;

fn main() {
    assert!(AsciiLowerAlpha::new(b'a').is_some());
    assert!(AsciiLowerAlpha::new(b'z').is_some());
    assert!(AsciiLowerAlpha::new(b'`').is_none());
    assert!(AsciiLowerAlpha::new(b'{').is_none());

    // This should build.
    let _ = AsciiLowerAlpha::from_const::<b'f'>();
    // This should not build.
    // let _ = AsciiLowerAlpha::from_const::<b'?'>();
}
