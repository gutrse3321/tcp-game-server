let net = require("net");

let options = {
    host: "127.0.0.1",
    port: 9000
};

let tcpClient = net.Socket();

tcpClient.connect(options, () => {
    console.log("connected to go rpcServer");
    let str = JSON.stringify({service: "getMyRealName", model: JSON.stringify({nickName: "yuki"})});
    tcpClient.write(str);
});

tcpClient.on("data", data => {
    console.log(new Date() + ": received data: %s from server", data.toString());
})

tcpClient.on("end", () => {
    console.log("data end!");
})

tcpClient.on("error", () => {
    console.log("tcpClient error!");
})

let str = JSON.stringify({service: "getMyRealName", model: JSON.stringify({nickName: "tomonori"})});
tcpClient.write(str);
