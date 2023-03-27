
use super::*;

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
struct Gen<T>(T);

type TestFunc<T,K> = fn(Gen<T>) -> K;

enum Exp<T,K> {
    A(TestFunc<T,K>)
}

#[test]
fn testEnum() {
    let a = TestFunc::from(|resp:String| {resp});
}