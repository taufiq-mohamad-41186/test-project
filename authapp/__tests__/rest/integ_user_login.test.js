const tcase = require('../../docs/test_case/test_case');
const request = require('supertest');
const App = require("../../app");
const { hasUncaughtExceptionCaptureCallback } = require('process');

const app = new App('./db/auth_test.db').getApp();

const testCase = new tcase.TestCase(
    "./docs/test_case/integ_post_user_login_meta.csv",
    "./docs/test_case/integ_post_user_login_req.csv",
    "./docs/test_case/integ_post_user_login_resp.csv"
).generate();

beforeEach(() => {
    // register user
});

afterEach(() => {
    // delete user
});

describe('Integ Test User Login', () => {
    testCase.meta.forEach((tc, i) => {
        it(tc.test_number + ": " + tc.test_desc, async () => {
            const respp = await request(app).post('/register').send({ "phone": testCase.req[i].phone, "role": "admin" });
            const payload = {
                "phone": testCase.req[i].phone,
                "password": respp.body.user.password,
            }
            const response = await request(app).post('/login').send(payload);
            request(app).delete('/' + testCase.req[i].phone);
            expect(response.statusCode).toEqual(parseInt(testCase.resp[i].response_code, 10));
        });
    });
});