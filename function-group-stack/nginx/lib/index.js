function e(e) {
  this.message = e;
}
e.prototype = new Error(), e.prototype.name = "InvalidCharacterError";
var r = "undefined" != typeof window && window.atob && window.atob.bind(window) || function (r) {
  var t = String(r).replace(/=+$/, "");
  if (t.length % 4 == 1) throw new e("'atob' failed: The string to be decoded is not correctly encoded.");
  for (var n, o, a = 0, i = 0, c = ""; o = t.charAt(i++); ~o && (n = a % 4 ? 64 * n + o : o, a++ % 4) ? c += String.fromCharCode(255 & n >> (-2 * a & 6)) : 0) o = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=".indexOf(o);
  return c;
};
function t(e) {
  var t = e.replace(/-/g, "+").replace(/_/g, "/");
  switch (t.length % 4) {
    case 0:
      break;
    case 2:
      t += "==";
      break;
    case 3:
      t += "=";
      break;
    default:
      throw "Illegal base64url string!";
  }
  try {
    return function (e) {
      return decodeURIComponent(r(e).replace(/(.)/g, function (e, r) {
        var t = r.charCodeAt(0).toString(16).toUpperCase();
        return t.length < 2 && (t = "0" + t), "%" + t;
      }));
    }(t);
  } catch (e) {
    return r(t);
  }
}
function n(e) {
  this.message = e;
}
function o(e, r) {
  if ("string" != typeof e) throw new n("Invalid token specified");
  var o = !0 === (r = r || {}).header ? 0 : 1;
  try {
    return JSON.parse(t(e.split(".")[o]));
  } catch (e) {
    throw new n("Invalid token specified: " + e.message);
  }
}
n.prototype = new Error(), n.prototype.name = "InvalidTokenError";
var __awaiter = undefined && undefined.__awaiter || function (thisArg, _arguments, P, generator) {
  function adopt(value) {
    return value instanceof P ? value : new P(function (resolve) {
      resolve(value);
    });
  }
  return new (P || (P = Promise))(function (resolve, reject) {
    function fulfilled(value) {
      try {
        step(generator.next(value));
      } catch (e) {
        reject(e);
      }
    }
    function rejected(value) {
      try {
        step(generator["throw"](value));
      } catch (e) {
        reject(e);
      }
    }
    function step(result) {
      result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected);
    }
    step((generator = generator.apply(thisArg, _arguments || [])).next());
  });
};
var __generator = undefined && undefined.__generator || function (thisArg, body) {
  var _ = {
      label: 0,
      sent: function sent() {
        if (t[0] & 1) throw t[1];
        return t[1];
      },
      trys: [],
      ops: []
    },
    f,
    y,
    t,
    g;
  return g = {
    next: verb(0),
    "throw": verb(1),
    "return": verb(2)
  }, typeof Symbol === "function" && (g[Symbol.iterator] = function () {
    return this;
  }), g;
  function verb(n) {
    return function (v) {
      return step([n, v]);
    };
  }
  function step(op) {
    if (f) throw new TypeError("Generator is already executing.");
    while (_) try {
      if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
      if (y = 0, t) op = [op[0] & 2, t.value];
      switch (op[0]) {
        case 0:
        case 1:
          t = op;
          break;
        case 4:
          _.label++;
          return {
            value: op[1],
            done: false
          };
        case 5:
          _.label++;
          y = op[1];
          op = [0];
          continue;
        case 7:
          op = _.ops.pop();
          _.trys.pop();
          continue;
        default:
          if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) {
            _ = 0;
            continue;
          }
          if (op[0] === 3 && (!t || op[1] > t[0] && op[1] < t[3])) {
            _.label = op[1];
            break;
          }
          if (op[0] === 6 && _.label < t[1]) {
            _.label = t[1];
            t = op;
            break;
          }
          if (t && _.label < t[2]) {
            _.label = t[2];
            _.ops.push(op);
            break;
          }
          if (t[2]) _.ops.pop();
          _.trys.pop();
          continue;
      }
      op = body.call(thisArg, _);
    } catch (e) {
      op = [6, e];
      y = 0;
    } finally {
      f = t = 0;
    }
    if (op[0] & 5) throw op[1];
    return {
      value: op[0] ? op[1] : void 0,
      done: true
    };
  }
};
var __read = undefined && undefined.__read || function (o, n) {
  var m = typeof Symbol === "function" && o[Symbol.iterator];
  if (!m) return o;
  var i = m.call(o),
    r,
    ar = [],
    e;
  try {
    while ((n === void 0 || n-- > 0) && !(r = i.next()).done) ar.push(r.value);
  } catch (error) {
    e = {
      error: error
    };
  } finally {
    try {
      if (r && !r.done && (m = i["return"])) m.call(i);
    } finally {
      if (e) throw e.error;
    }
  }
  return ar;
};
var cr = require("crypto");
var decodeBase64 = function decodeBase64(str) {
  return atob(str.replace(/-/g, "+").replace(/_/g, "/"));
};
function base64urlEscape(str) {
  return str.replace(/\+/g, "-").replace(/\//g, "_").replace(/=/g, "");
}
function sign(headAndPayload, key) {
  var base64str = cr.createHmac("sha256", key).update(headAndPayload).digest("base64");
  return base64urlEscape(base64str);
}
function verify(headAndPayload, key, signature) {
  return signature === sign(headAndPayload, key);
}
function hello(r) {
  var _a, _b;
  return __awaiter(this, void 0, void 0, function () {
    var jwt, _c, header, payload, signature, valid, token, decoded, function_data, functionName, serviceName, namespace, serviceUri, res, body;
    return __generator(this, function (_d) {
      switch (_d.label) {
        case 0:
          if (!r.headersIn.Authorization) {
            r.log("No token provided");
            r["return"](403, "No token provided");
            return [2 /*return*/];
          }

          jwt = (_a = r.headersIn.Authorization) === null || _a === void 0 ? void 0 : _a.replace("Bearer ", "");
          _c = __read(jwt.split("."), 3), header = _c[0], payload = _c[1], signature = _c[2];
          valid = verify([header, payload].join("."), "signing_secret", signature);
          if (!valid) {
            r["return"](503, "Token not valid");
            r.error("Token not valid");
            return [2 /*return*/];
          }

          token = o(jwt);
          r.log("Token: " + JSON.stringify(token));
          r.log("Temp token: " + r.variables.temp_token);
          if (!token.tempTokens || !token.tempTokens.find(function (h) {
            return h === r.variables.temp_token;
          })) {
            r["return"](403, "You are not allowed to use this group.");
          }
          decoded = decodeBase64(r.variables.function_data);
          function_data = JSON.parse(decoded);
          functionName = r.uri.split("/").reverse()[0];
          serviceName = (_b = function_data.find(function (f) {
            return f.name === functionName;
          })) === null || _b === void 0 ? void 0 : _b.service;
          if (!serviceName) {
            r["return"](404, "Not found");
            r.error(functionName + " does not exist");
            return [2 /*return*/];
          }

          namespace = r.variables.namespace;
          serviceUri = "http://" + serviceName + "." + namespace + ".svc.cluster.local:8080";
          r.log("mapping from " + r.uri + " to " + serviceUri);
          return [4 /*yield*/, ngx.fetch(serviceUri, {
            method: "POST",
            body: r.requestText
          })];
        case 1:
          res = _d.sent();
          return [4 /*yield*/, res.text()];
        case 2:
          body = _d.sent();
          r.log("Function " + functionName + " returned status " + res.status + " with body " + body);
          r["return"](r.status, body);
          return [2 /*return*/];
      }
    });
  });
}

export default { hello };
