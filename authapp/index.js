const dotenv = require('dotenv').config();
const App = require('./app');
const http = require('http');

// running server
const app = new App('./db/auth.db');
const port = process.env.PORT || 3000;
http.createServer(app.getApp()).listen(port)
console.log('Server listening at http://localhost:' + port)
