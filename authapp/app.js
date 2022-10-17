const dotenv = require('dotenv').config();
const app = require('express')();
const http = require('http');
const swaggerUi = require('swagger-ui-express')
const swaggerFile = require('./docs/swagger_output.json')
const sqlite3 = require('sqlite3').verbose();

class App {
    constructor(db) {
        console.log(db);
        // database preparation
        this.db = new sqlite3.Database(db);

        // model preparation
        const User = require('./business/model/user');
        const UserModel = new User(this.db);

        // service preparation
        const Auth = require('./business/service/auth');
        const AuthService = new Auth(UserModel);

        app.use('/swagger', swaggerUi.serve, swaggerUi.setup(swaggerFile))
        require('./handler/rest')(app, AuthService)
    }
    getApp = () => {
        return app;
    }
}

module.exports = App;