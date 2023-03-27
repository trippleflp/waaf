use derivative::Derivative;
use http;
use std::env;

pub mod handle_func;
mod wasmedge_connect;

#[derive(Debug, Derivative)]
#[derivative(Default)]
pub struct WaafFunction {
    pub handle_func: handle_func::HandleFunc,
}

impl WaafFunction {
    pub fn start(self) {
        match wasmedge_connect::start_socket(self) {
            Ok(_) => {}
            Err(e) => {
                eprintln!("Error: {}", e)
            }
        };
    }
}

#[cfg(test)]
mod test;
