import jwt_decode, {JwtPayload} from "jwt-decode";

// Todo verify
export async function hello(r: NginxHTTPRequest) {
    if (!r.headersIn.Authorization) {
        r.log("No token provided");
        r.return(403, "No token provided");
        return;
    }

    const jwt = r.headersIn.Authorization?.replace("Bearer ", "");

    const token = jwt_decode<JwtPayload>(jwt);

    r.log(r.variables.GROUP_ID as NjsStringOrBuffer);
    if (r.variables.group_id !== token.aud) {
        r.return(403, "You are not allowed to use this group.");
    }

    const decoded = atob(r.variables.function_data as string);
    const function_data: [{ name: string; service:string }] = JSON.parse(decoded);
    const functionName = r.uri.split("/").reverse()[0];

    const serviceName = function_data.find((f) => f.name === functionName)?.service;
    if (!serviceName) {
        r.return(404, "Not found");
        r.error(`${functionName} does not exist`);
        return;
    }
    const namespace = r.variables.namespace as string
    const serviceUri = `http://${serviceName}.${namespace}.svc.cluster.local:8080`

    r.log(`mapping from ${r.uri} to ${serviceUri}`)
    const res = await ngx.fetch(serviceUri,{method: "POST", body: r.requestText})

    const body = await res.text()
    r.log(`Function ${functionName} returned status ${res.status} with body ${body}`)

    r.return(r.status, body);
}
