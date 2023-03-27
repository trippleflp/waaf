use bytecodec::{DecodeExt, Error};
use httpcodec::{HttpVersion, ReasonPhrase, Request, RequestDecoder, Response, StatusCode};
use serde_json::Value;

pub type HandleFuncType<T> = fn(req: T) -> bytecodec::Result<Response<String>>;
pub type HandleFuncTwoParameterType<T, K> =
    fn(req: T, data: K) -> bytecodec::Result<Response<String>>;

#[derive(Debug)]
pub enum HandleFunc {
    StringBody(HandleFuncType<Request<String>>),
    BufferBody(HandleFuncType<Vec<u8>>),
    JsonBuffer(HandleFuncTwoParameterType<Request<String>, Value>),
}
impl Default for HandleFunc {
    fn default() -> Self {
        todo!()
    }
}

impl HandleFunc {
    pub fn call(&self, data: Vec<u8>) -> String {
        let response = self.handle(data);

        let r = match response {
            Ok(r) => r,
            Err(e) => {
                let err = format!("{:?}", e);
                Response::new(
                    HttpVersion::V1_0,
                    StatusCode::new(500).unwrap(),
                    ReasonPhrase::new(err.as_str()).unwrap(),
                    err.clone(),
                )
            }
        };

        r.to_string()
    }

    fn handle(&self, data: Vec<u8>) -> bytecodec::Result<Response<String>> {
        match self {
            HandleFunc::StringBody(handle_func) => {
                let req = HandleFunc::decode_string(data)?;
                handle_func(req)
            }
            HandleFunc::BufferBody(handle_func) => handle_func(data),
            HandleFunc::JsonBuffer(handle_func) => {
                let req = HandleFunc::decode_string(data)?;
                let json = match serde_json::from_str(req.body()) {
                    Ok(json) => Ok(json),
                    Err(e) => Err(std::io::Error::from(e)),
                }?;
                handle_func(req, json)
            }
        }
    }

    fn decode_string(data: Vec<u8>) -> Result<Request<String>, Error> {
        let mut decoder =
            RequestDecoder::<httpcodec::BodyDecoder<bytecodec::bytes::Utf8Decoder>>::default();

        return match decoder.decode_from_bytes(data.as_slice()) {
            Ok(req) => Ok(req),
            Err(e) => Err(e),
        };
    }
}
