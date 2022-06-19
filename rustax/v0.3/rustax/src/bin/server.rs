use server::server;

fn main() {
    let foo = server();
    match foo {
        Ok(_) => {},
        Err(err) => panic!("error {}", err),
    };
}