extern crate rocket;

use clap::Parser;

use server::server;
use tax::Calculator;
use cli::{Cli, Commands};

#[rocket::main]
async fn main() {
    let cli = Cli::parse();

    match &cli.command {
        Commands::Calc { amount } => {
            let calc = Calculator::new(*amount);
            let result = calc.taxe();
        
            println!("For an income of {}, you'll pay {}€ of tax", amount, result);
        }

        Commands::Serv { port, host } => {
            match server(port, host).await {
                Ok(_) => {},
                Err(err) => panic!("error {}", err),
            };
        }
    }
}
