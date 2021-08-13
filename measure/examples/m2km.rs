use measure::Length;

fn main() {
    let km = 5.km();
    let m = km.m();
    println!("there are {} meters in {} kilometers", *m, *km);
}
