import jwt_decode, { JwtPayload } from "jwt-decode";
let cr = require("crypto");

interface WaafJwtPayload extends JwtPayload {
  tempTokens?: string[];
}

const decodeBase64 = (str: String) =>
  atob(str.replace(/-/g, "+").replace(/_/g, "/"));

function base64urlEscape(str: String) {
  return str.replace(/\+/g, "-").replace(/\//g, "_").replace(/=/g, "");
}

function sign(headAndPayload: String, key: String) {
  const base64str = cr
    .createHmac("sha256", key)
    .update(headAndPayload)
    .digest("base64");

  return base64urlEscape(base64str);
}

function verify(headAndPayload: String, key: String, signature: String) {
  return signature === sign(headAndPayload, key);
}

export async function hello(r: NginxHTTPRequest) {
  if (!r.headersIn.Authorization) {
    r.log("No token provided");
    r.return(403, "No token provided");
    return;
  }

  const jwt = r.headersIn.Authorization?.replace("Bearer ", "");
  const [header, payload, signature] = jwt.split(".");

  const valid = verify(
    [header, payload].join("."),
    "signing_secret",
    signature
  );

  if (!valid) {
    r.return(503, "Token not valid");
    r.error(`Token not valid`);
    return;
  }

  const token = jwt_decode<WaafJwtPayload>(jwt);
  r.log("Token: " + JSON.stringify(token));
  r.log(("Temp token: " + r.variables.temp_token) as NjsStringOrBuffer);
  if (
    !token.tempTokens ||
    !token.tempTokens.find((h) => h === r.variables.temp_token)
  ) {
    r.return(403, "You are not allowed to use this group.");
  }

  const decoded = decodeBase64(r.variables.function_data as string);
  const function_data: [{ name: string; service: string }] =
    JSON.parse(decoded);
  const functionName = r.uri.split("/").reverse()[0];

  const serviceName = function_data.find(
    (f) => f.name === functionName
  )?.service;
  if (!serviceName) {
    r.return(404, "Not found");
    r.error(`${functionName} does not exist`);
    return;
  }
  const namespace = r.variables.namespace as string;
  const serviceUri = `http://${serviceName}.${namespace}.svc.cluster.local:8080`;

  r.log(`mapping from ${r.uri} to ${serviceUri}`);
  const res = await ngx.fetch(serviceUri, {
    method: "POST",
    body: r.requestText,
  });

  const body = await res.text();
  r.log(
    `Function ${functionName} returned status ${res.status} with body ${body}`
  );

  r.return(r.status, body);
}
