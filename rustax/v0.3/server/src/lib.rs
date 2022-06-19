#[macro_use]
extern crate rocket;

use rocket::serde::{json::Json, Serialize};
use tax::Calculator;

#[get("/world")]
fn world() -> &'static str {
    "Hello, world!"
}

#[derive(Serialize)]
#[serde(crate = "rocket::serde")]
struct Amount {
    amount: f32,
}

#[get("/?<amount>")]
fn calc(amount: f32) -> Json<Amount> {
    let calc = Calculator::new(amount);
    let result = calc.taxe();

    return Json(Amount { amount: result });
}

#[rocket::main]
pub async fn server() -> Result<(), rocket::Error> {
    let _rocket = rocket::build()
        .mount("/hello", routes![world])
        .mount("/tax", routes![calc])
        .launch()
        .await?;

    Ok(())
}
