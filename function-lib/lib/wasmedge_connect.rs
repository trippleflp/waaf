use crate::handle_func::HandleFunc;
use crate::WaafFunction;
use httpcodec::{HttpVersion, ReasonPhrase, Request, Response, StatusCode};
use std::io::{Read, Write};
#[cfg(not(feature = "std"))]
use wasmedge_wasi_socket::{Shutdown, TcpListener, TcpStream};
use std::env;

// used as example
fn handle_http(req: Request<String>) -> bytecodec::Result<Response<String>> {
    Ok(Response::new(
        HttpVersion::V1_0,
        StatusCode::new(200)?,
        ReasonPhrase::new("")?,
        format!("echo: {}", req.body()),
    ))
}

fn handle_client(mut stream: TcpStream, handle_func: &HandleFunc) -> std::io::Result<()> {
    let mut buff = [0u8; 1024];
    let mut data = Vec::new();

    loop {
        let n = stream.read(&mut buff)?;
        data.extend_from_slice(&buff[0..n]);
        if n < 1024 {
            break;
        }
    }

    let response = handle_func.call(data);
    stream.write(response.as_bytes())?;
    stream.shutdown(Shutdown::Both)?;
    Ok(())
}

pub fn start_socket(waaf: WaafFunction) -> std::io::Result<()> {
    let port = std::env::var("PORT").unwrap_or("1234".to_string());
    println!("new connection at 0.0.0.0:{}", port);
    let listener = TcpListener::bind(format!("0.0.0.0:{}", port), false)?;
    loop {
        let res = handle_client(listener.accept(false)?.0, &waaf.handle_func);
    }

    println!("exiting");

    Ok(())
}
