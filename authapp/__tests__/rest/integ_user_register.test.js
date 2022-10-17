const tcase = require('../../docs/test_case/test_case');
const request = require('supertest');
const App = require("../../app");
const { hasUncaughtExceptionCaptureCallback } = require('process');

const app = new App('./db/auth_test.db').getApp();

const testCase = new tcase.TestCase(
    "./docs/test_case/integ_post_user_register_meta.csv",
    "./docs/test_case/integ_post_user_register_req.csv",
    "./docs/test_case/integ_post_user_register_resp.csv"
).generate();

describe('Integ Test User Register', () => {
    testCase.meta.forEach((tc, i) => {
        it(tc.test_number + ": " + tc.test_desc, async () => {
            const payload = {
                "phone": testCase.req[i].phone,
                "role": testCase.req[i].role
            }
            // console.log(payload);
            const response = await request(app).post('/register').send(payload);
            // console.log(response.body);
            expect(response.statusCode).toEqual(parseInt(testCase.resp[i].response_code, 10));
        });
    });
});