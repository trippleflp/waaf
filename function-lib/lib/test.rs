use super::*;
use httpcodec::{HttpVersion, ReasonPhrase, Request, RequestDecoder, Response, StatusCode};

// #[test]
// fn test_waaf_wrapper() {
//     let waaf = WaafFunction{handle_func: |request| {
//         Ok(Response::new(
//             HttpVersion::V1_0,
//             StatusCode::new(200)?,
//             ReasonPhrase::new("")?,
//             format!("echo: {}", req.body()),
//         ))
//     },..Default::default()};
// }

type TestFunc<T, K> = dyn Fn(T) -> K;

enum Exp {
    A(TestFunc<String, String>),
    B(TestFunc<i8, i8>),
}

#[test]
fn testEnum() {
    let a: TestFunc<String, String> = Exp::A(|string| string);

    assert!(a("lol".parse().unwrap()) == "lol".parse().unwrap())
}
