const CodeHTTPInternalServerError = Object.fromEntries(new Map([
    ["Code", 500],
    ["Description", "Internal Server Error"]
]));
const CodeHTTPBadRequest = Object.fromEntries(new Map([
    ["Code", 400],
    ["Description", "Bad Request"]
]));
const CodeHTTPSQLDuplicateEntity = Object.fromEntries(new Map([
    ["Code", 409],
    ["Description", "Duplicate Entity"]
]));
const CodeHTTPUnauthorized = Object.fromEntries(new Map([
    ["Code", 401],
    ["Description", "Unauthorized"]
]));

exports.CodeHTTPInternalServerError = CodeHTTPInternalServerError;
exports.CodeHTTPBadRequest = CodeHTTPBadRequest;
exports.CodeHTTPSQLDuplicateEntity = CodeHTTPSQLDuplicateEntity;
exports.CodeHTTPUnauthorized = CodeHTTPUnauthorized;

exports.wrapWithCode = (code, err, message = null) => {
    var stack = new Error(err).stack;
    return { "code": code.Code, "description": code.Description, "message": message, "err": stack }
}

exports.wrap = (err, message = null) => {
    var stack = new Error(err).stack;
    return { "code": 500, "description": "Internal Server Error", "message": message, "err": stack }
}

exports.newWithCode = (code, message = null) => {
    var stack = new Error().stack;
    return { "code": code.Code, "description": code.Description, "message": message, "err": stack }
}

exports.new = (message = null) => {
    var stack = new Error().stack;
    return { "code": 500, "description": "Internal Server Error", "message": message, "err": stack }
}

