const fs = require("fs");

class TestCase {
    constructor(rawMeta, rawReq, rawResp) {
        this.rawMeta = rawMeta;
        this.rawReq = rawReq;
        this.rawResp = rawResp;
        this.meta = [];
        this.metaHeader = [];
        this.metaData = [];
        this.req = [];
        this.reqHeader = [];
        this.reqData = [];
        this.resp = [];
        this.respHeader = [];
        this.respData = [];
        this.meta = (() => {
            fs.readFileSync(rawMeta, "utf8").split("\n").forEach((row, i) => {
                if (i == 0) {
                    this.metaHeader = row.split("|");
                } else {
                    this.metaData.push(row.split("|"));
                }
            });


            this.metaData.forEach((data, i) => {
                let tmpMeta = []
                this.metaHeader.forEach((header, j) => {
                    tmpMeta.push([header, data[j]]);
                })
                const map = new Map(tmpMeta);
                const obj = Object.fromEntries(map);
                this.meta.push(obj);
            });

            return this.meta;
        })();

        this.req = (() => {
            fs.readFileSync(rawReq, "utf8").split("\n").forEach((row, i) => {
                if (i == 0) {
                    this.reqHeader = row.split("|");
                } else {
                    this.reqData.push(row.split("|"));
                }
            });


            this.reqData.forEach((data, i) => {
                let tmpReq = []
                this.reqHeader.forEach((header, j) => {
                    tmpReq.push([header, data[j]]);
                })
                const map = new Map(tmpReq);
                const obj = Object.fromEntries(map);
                this.req.push(obj);
            });

            return this.req;
        })();

        this.resp = (() => {
            fs.readFileSync(rawResp, "utf8").split("\n").forEach((row, i) => {
                if (i == 0) {
                    this.respHeader = row.split("|");
                } else {
                    this.respData.push(row.split("|"));
                }
            });


            this.respData.forEach((data, i) => {
                let tmpResp = []
                this.respHeader.forEach((header, j) => {
                    tmpResp.push([header, data[j]]);
                })
                const map = new Map(tmpResp);
                const obj = Object.fromEntries(map);
                this.resp.push(obj);
            });

            return this.resp;
        })();
    }

    generate = () => {
        return {
            meta: this.meta,
            req: this.req,
            resp: this.resp
        }
    }
}

exports.TestCase = TestCase;
