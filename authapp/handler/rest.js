const express = require('express');
const bodyParser = require('body-parser');
const x = require("../business/model/error");
const v = require('../docs/schema/validator');

module.exports = function (app, authService) {

    app.use(bodyParser.urlencoded({ extended: false }));
    app.use(bodyParser.json());

    app.post('/register', (req, res) => {
        new Promise((accepted, rejected) => {
            console.log(req.body);
            const val = new v.Validator(v.UserCreateRequestSchema).validate(req.body, v.UserCreateRequestSchema);
            if (val) return rejected(x.newWithCode(x.CodeHTTPBadRequest, val));

            const phone = req.body.phone;
            const role = req.body.role;
            authService.registerUser(phone, role, (err, user) => {
                if (err) return rejected(err);

                return accepted(user);
            })
        }).then(
            (data) => {
                res.status(201).send({
                    "user": { "id": data.id, "phone": data.phone, "role": data.role, "password": data.password }, "access_token": data.accessToken, "expires_in": data.expiresIn
                });
            },
            (error) => {
                res.status(error.code).send({ "code": error.code, "description": error.description, "message": error.message, "err": error.err });
            }
        );
    });

    app.post('/login', (req, res) => {
        new Promise((accepted, rejected) => {
            const val = new v.Validator(v.UserLoginRequestSchema).validate(req.body, v.UserLoginRequestSchema);
            if (val) return rejected(x.newWithCode(x.CodeHTTPBadRequest, val));

            const phone = req.body.phone;
            const password = req.body.password;
            authService.login(phone, password, (err, user) => {
                if (err) return rejected(err);
                if (!user) return rejected(x.newWithCode(x.CodeHTTPUnauthorized, "Invalid phone/password"));
                return accepted(user);
            })
        }).then(
            (data) => {
                res.status(201).send({
                    "user": { "id": data.id, "phone": data.phone, "role": data.role, "password": data.password }, "access_token": data.accessToken, "expires_in": data.expiresIn
                });
            },
            (error) => {
                res.status(error.code).send({ "code": error.code, "description": error.description, "message": error.message, "err": error.err });
            }
        );
    });

    app.delete('/:phone', (req, res) => {
        new Promise((accepted, rejected) => {
            const phone = req.params.phone;
            authService.deleteUser(phone, (err) => {
                if (err) return rejected(err);
                return accepted(phone);
            })
        }).then(
            (data) => {
                res.status(200).send({
                    "message": "delete success"
                });
            },
            (error) => {
                res.status(error.code).send({ "code": error.code, "description": error.description, "message": error.message, "err": error.err });
            }
        );
    });
}