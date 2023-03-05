import jwt_decode, { JwtPayload } from "jwt-decode";
// Todo verify
export async function hello(r: NginxHTTPRequest) {
  if (!r.headersIn.Authorization) {
    r.log("No token provided");
    r.return(403, "No token provided");
    return;
  }

  const jwt = r.headersIn.Authorization?.replace("Bearer ", "");

  const token = jwt_decode<JwtPayload>(jwt);

  r.log(r.variables.group_id as NjsStringOrBuffer);
  r.log(r.variables.GROUP_ID as NjsStringOrBuffer);
  if (r.variables.group_id !== token.aud) {
    r.return(403, "You are not allowed to use this group.");
  }

  r.return(200, JSON.stringify(r.variables.group_id));
}
