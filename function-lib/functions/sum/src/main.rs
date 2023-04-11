use httpcodec::{HttpVersion, ReasonPhrase, Request, Response, StatusCode};
use waaf::handle_func::HandleFunc;
use waaf::WaafFunction;

fn main() {
    let waaf = WaafFunction {
        handle_func: HandleFunc::StringBody(|req| {
            Ok(Response::new(
                HttpVersion::V1_0,
                StatusCode::new(200)?,
                ReasonPhrase::new("")?,
                format!("echo: {}", req.body()),
            ))
        }),
    };
    waaf.start();
}
