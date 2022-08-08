use clap::{Parser, Subcommand};

/// A binary to calculate or run a tax calculator server
#[derive(Parser)]
#[clap(author, name = "rustax", version = "0.3", about)]
pub struct Cli {
    #[clap(subcommand)]
    pub command: Commands,
}

#[derive(Subcommand)]
pub enum Commands {
    /// Directly calculate tax
    Calc {
        /// Set the amount of tax to calculate
        #[clap(long)]
        amount: f32,
    },

    /// Start rustax server
    Serv {
        /// Set server port
        #[clap(long, default_value_t = 8888)]
        port: i32,

        /// Set server host
        #[clap(long, default_value = "0.0.0.0")]
        host: String
    },
}
