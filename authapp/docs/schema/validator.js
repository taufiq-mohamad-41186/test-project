const JsonSchema = require('jsonschema').Validator;
const v = new JsonSchema();

const userCreateSchema = require('./user-create-request.schema.json');
const userLoginSchema = require('./user-login-request.schema.json');

class Validator {
    constructor(schema) {
        this.validator = v;
        this.validator.addSchema(schema);
    }

    validate = (data, schema) => {
        const val = this.validator.validate(data, schema)
        if (val.errors.length != 0) {
            let message = "";
            val.errors.forEach((error) => {
                message += error.property + " " + error.message + " \n ";
            })
            return message;
        }
    }
}

exports.Validator = Validator;
exports.UserCreateRequestSchema = userCreateSchema;
exports.UserLoginRequestSchema = userLoginSchema;