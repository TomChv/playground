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

pub async fn server(port: &i32, host: &String) -> Result<(), rocket::Error> {
    let figment = rocket::Config::figment()
        .merge(("port", port))
        .merge(("address", host));

    let _rocket = rocket::custom(figment)
        .mount("/hello", routes![world])
        .mount("/tax", routes![calc])
        .launch()
        .await?;

    Ok(())
}
