use measure::Time;

fn main() {
    let h = 5.h();
    let m = h.m();
    println!("there are {} minutes in {} hours", *m, *h);
}
