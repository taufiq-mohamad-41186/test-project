class UserModel {
    constructor(database) {
        this.database = database;

        const sqlQuery = 'CREATE TABLE IF NOT EXISTS users(id integer PRIMARY KEY, phone text UNIQUE, role text, encrypted_password text)';
        console.log("init database");

        this.database.run(sqlQuery, (err) => {
            if (err) { console.log(err) }
        });
    }

    findUserByPhone = (phone, cb) => {
        this.database.get(`SELECT * FROM users WHERE phone = ?`, [phone], (err, row) => {
            if (err) return cb(err, undefined);
            return cb(undefined, row);
        });
    }

    createUser = (user, cb) => {
        this.database.run('INSERT INTO users (phone, role, encrypted_password) VALUES (?,?,?)', user, (err) => {
            if (err) return cb(err);
            return cb();
        });
    }

    deleteUser = (phone, cb) => {
        this.database.run('DELETE FROM users WHERE phone = ?', phone, (err) => {
            if (err) return cb(error);
            return cb();
        });
    }
}

module.exports = UserModel;