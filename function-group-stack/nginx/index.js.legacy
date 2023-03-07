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
// Todo verify
function hello(r) {
  var _a;
  return __awaiter(this, void 0, void 0, function () {
    var jwt, token;
    return __generator(this, function (_b) {
      if (!r.headersIn.Authorization) {
        r.log("No token provided");
        r["return"](403, "No token provided");
        return [2 /*return*/];
      }

      jwt = (_a = r.headersIn.Authorization) === null || _a === void 0 ? void 0 : _a.replace("Bearer ", "");
      token = o(jwt);
      r.log(r.variables.group_id);
      r.log(r.variables.GROUP_ID);
      if (r.variables.group_id !== token.aud) {
        r["return"](403, "You are not allowed to use this group.");
      }
      r["return"](200, JSON.stringify(r.variables.group_id));
      return [2 /*return*/];
    });
  });
}

export default { hello };
