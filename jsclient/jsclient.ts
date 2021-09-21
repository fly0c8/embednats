import { connect,StringCodec } from "https://raw.githubusercontent.com/nats-io/nats.ws/master/src/mod.ts";

const sc = StringCodec();
const nc = await connect({servers: ["ws://localhost:8081"]});
nc.subscribe("hello", {callback: (err, msg) => {
    console.log(`${sc.decode(msg.data)}`);
}});
nc.publish("hello", sc.encode("Nats from Websocket"));