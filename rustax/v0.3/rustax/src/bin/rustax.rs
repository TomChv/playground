use clap::Parser;

use cli::{Cli, Commands};

fn main() {
    let cli = Cli::parse();

    match &cli.command {
        Commands::Calc { amount } => {
            println!("Calc {}", amount)
        }

        Commands::Serv { port, host } => {
            println!("Serv {}:{}", port, host)
        }
    }
}
