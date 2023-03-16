var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
import jwt_decode from "jwt-decode";
// Todo verify
export function hello(r) {
    var _a, _b;
    return __awaiter(this, void 0, void 0, function () {
        var jwt, token, decoded, function_data, functionName, serviceName, namespace, serviceUri, res, body;
        return __generator(this, function (_c) {
            switch (_c.label) {
                case 0:
                    if (!r.headersIn.Authorization) {
                        r.log("No token provided");
                        r.return(403, "No token provided");
                        return [2 /*return*/];
                    }
                    jwt = (_a = r.headersIn.Authorization) === null || _a === void 0 ? void 0 : _a.replace("Bearer ", "");
                    token = jwt_decode(jwt);
                    r.log(r.variables.GROUP_ID);
                    if (r.variables.group_id !== token.aud) {
                        r.return(403, "You are not allowed to use this group.");
                    }
                    decoded = atob(r.variables.function_data);
                    function_data = JSON.parse(decoded);
                    functionName = r.uri.split("/").reverse()[0];
                    serviceName = (_b = function_data.find(function (f) { return f.name === functionName; })) === null || _b === void 0 ? void 0 : _b.service;
                    if (!serviceName) {
                        r.return(404, "Not found");
                        r.error(functionName + " does not exist");
                        return [2 /*return*/];
                    }
                    namespace = r.variables.namespace;
                    serviceUri = "http://" + serviceName + "." + namespace + ".svc.cluster.local:8080";
                    r.log(r.requestText);
                    r.log(r.requestBuffer);
                    r.log(r.requestBody);
                    r.log("mapping from " + r.uri + " to " + serviceUri);
                    return [4 /*yield*/, ngx.fetch(serviceUri, { method: "POST", body: r.requestText })];
                case 1:
                    res = _c.sent();
                    return [4 /*yield*/, res.text()];
                case 2:
                    body = _c.sent();
                    r.log("Function " + functionName + " returned status " + res.status + " with body " + body);
                    // const res = await r.subrequest(`http://localhost:${port}`);
                    r.return(r.status, body);
                    return [2 /*return*/];
            }
        });
    });
}
