const dotenv = require('dotenv').config();
const jwt = require('jsonwebtoken');
const bcrypt = require('bcryptjs');
const x = require('../model/error');
const SECRET_KEY = process.env.JWT_SECRET_KEY;

class AuthService {
    constructor(userModel) {
        this.userModel = userModel;
    }
    registerUser = (phone, role, cb) => {
        this.userModel.findUserByPhone(phone, (err, user) => {
            if (err) return cb(x.wrap(err), user);
            if (user) return cb(x.newWithCode(x.CodeHTTPSQLDuplicateEntity));
            const password = Math.random().toString(36).slice(-4);
            const encrypted_password = bcrypt.hashSync(password);
            this.userModel.createUser([phone, role, encrypted_password], (err) => {
                if (err) return cb(x.wrap(err), user);
                this.userModel.findUserByPhone(phone, (err, user) => {
                    if (err) return cb(x.wrap(err), user);
                    const expiresIn = 24 * 60 * 60;
                    const accessToken = jwt.sign({ id: user.id, phone: user.phone, role: user.role }, SECRET_KEY, {
                        expiresIn: expiresIn
                    });
                    user.password = password;
                    user.accessToken = accessToken;
                    user.expiresIn = expiresIn;

                    return cb(err, user);
                })
            })
        });
    }

    login = (phone, password, cb) => {
        this.userModel.findUserByPhone(phone, (err, user) => {
            if (err) return cb(x.wrap(err), user);
            if (!user) return cb(x.newWithCode(x.CodeHTTPUnauthorized), user);
            const result = bcrypt.compareSync(password, user.encrypted_password);
            if (!result) return cb(x.newWithCode(x.CodeHTTPUnauthorized), user);

            const expiresIn = 24 * 60 * 60;
            const accessToken = jwt.sign({ id: user.id, phone: user.phone, role: user.role }, SECRET_KEY, {
                expiresIn: expiresIn
            });
            user.accessToken = accessToken;
            user.expiresIn = expiresIn;

            return cb(err, user);
        });
    };

    deleteUser = (phone, cb) => {
        this.userModel.deleteUser(phone, (err) => {
            if (err) return cb(x.wrap(err));

            return cb();
        });
    };
}

module.exports = AuthService;