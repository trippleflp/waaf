use httpcodec::{HttpVersion, ReasonPhrase, Request, Response, StatusCode};
use waaf::WaafFunction;
use waaf::handle_func::HandleFunc;

fn main() {
    let waaf = WaafFunction{
        handle_func: HandleFunc::StringBody(|req|     Ok(Response::new(
            HttpVersion::V1_0,
            StatusCode::new(200)?,
            ReasonPhrase::new("")?,
            format!("echo: {}", req.body()),
        ))),
        // ..Default::default()
    };
    waaf.start();
}
