use crate::handle_func::HandleFunc;
use crate::WaafFunction;
use httpcodec::{HttpVersion, ReasonPhrase, Request, Response, StatusCode};

fn StringBuffer_Test() {
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

fn JsonBuffer_Test() {
    let waaf = WaafFunction {
        handle_func: HandleFunc::BufferBody(|bytes| {
            Ok(Response::new(
                HttpVersion::V1_0,
                StatusCode::new(200)?,
                ReasonPhrase::new("")?,
                format!("bytes: {}", String::from(bytes)),
            ))
        }),
    };
    waaf.start();
}

fn JsonBuffer_Test() {
    let waaf = WaafFunction {
        handle_func: HandleFunc::JsonBuffer(|req, json| {
            Ok(Response::new(
                HttpVersion::V1_0,
                StatusCode::new(200)?,
                ReasonPhrase::new("")?,
                format!("json: {}", json.to_string()),
            ))
        }),
    };
    waaf.start();
}
