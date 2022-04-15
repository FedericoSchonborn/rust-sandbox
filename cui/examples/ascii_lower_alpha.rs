use cui::U8;

fn main() {
    // This type only accepts ASCII characters ranging from 'a' to 'z'.
    type AsciiLowerAlpha = U8<0x61, 0x7A>;
    assert!(AsciiLowerAlpha::new(b'a').is_some());
    assert!(AsciiLowerAlpha::new(b'z').is_some());
    assert!(AsciiLowerAlpha::new(b'`').is_none());
    assert!(AsciiLowerAlpha::new(b'{').is_none());

    #[cfg(feature = "guard")]
    AsciiLowerAlpha::from_const::<b'f'>();
}
