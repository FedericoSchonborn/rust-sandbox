use measure::{Kilometers, Measure};

fn main() {
    let m = 5.m();
    let km: Kilometers = m.into();
    println!("{} meters are {} kilometers", m.get(), km.get())
}
