use measure::{Measure, Meters};

fn main() {
    let km = 5.km();
    let m: Meters = km.into();
    println!("there are {} meters in {} kilometers", *m, *km);
}
