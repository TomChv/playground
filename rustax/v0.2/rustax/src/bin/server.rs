#[macro_use] extern crate rocket;

use rocket::serde::{Serialize, json::Json};

#[get("/world")]
fn world() -> &'static str {
    "Hello, world!"
}

#[derive(Serialize)]
#[serde(crate = "rocket::serde")]
struct Amount {
    amount: f32,
}

#[get("/urssaf?<amount>")]
fn calc(amount: f32) -> Json<Amount> {
    return Json(Amount { amount })
}

#[rocket::main]
async fn main() -> Result<(), rocket::Error> {
    let _rocket = rocket::build()
        .mount("/hello", routes![world])
        .mount("/tax", routes![calc])
        .launch()
        .await?;

    Ok(())
}