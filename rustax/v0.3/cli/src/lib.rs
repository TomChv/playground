use clap::Parser;

/// Simple program to greet a person
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
pub struct Args {
    /// Set the amount of tax to calculate
    #[clap(short, long)]
    pub amount: f32,
}