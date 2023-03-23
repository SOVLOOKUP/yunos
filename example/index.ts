import WebSocket from 'ws';
import { JsonRpcEngine, JsonRpcRequest } from 'json-rpc-engine';
import destr from 'destr'
const engine = new JsonRpcEngine();
const ws = new WebSocket('ws://127.0.0.1:3000/ws');
import { nanoid } from 'nanoid';

ws.on('error', console.error);

ws.on('open', function open() {
    event({ "hello": "world" })
});

engine.push((req, res, next, end) => {
    switch (req.method) {
        case "meta":
            res.result = {
                user_id: "",
                device_id: "",
                endpoints: []
            }
            end();
            break;

        default:
            return next()
        // break;
    }

})
ws.on('message', async (data) => {
    const req = destr(data.toString())
    console.log(req);
    const r = await engine.handle(req)
    console.log(r);
    const result = JSON.stringify(r)
    ws.send(result)
});

const event = <T>(params?: T) => ws.send(JSON.stringify({
    jsonrpc: "2.0",
    method: "event",
    id: nanoid(),
    params
}))

