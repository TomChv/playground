use clap::Parser;

use cli::Args;
use tax::Calculator;

fn main() {
    let args = Args::parse();

    let amount = args.amount;
    let calc = Calculator::new(amount);
    let result = calc.taxe();

    println!("For an income of {}, you'll pay {}€ of tax", amount, result);
}