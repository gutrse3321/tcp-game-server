let net = require("net");

let options = {
    host: "127.0.0.1",
    port: 9000
};

let tcpClient = net.Socket();

tcpClient.connect(options, () => {
    console.log("connected to go rpcServer");
});

tcpClient.on("data", data => {
    console.log(new Date() + ": received data: %s from server", data.toString());
    let str = JSON.stringify({handler: "lobby", service: "GetLobby", model: JSON.stringify({nickName: "yuki"})});
    tcpClient.write(str);
})

tcpClient.on("end", () => {
    console.log("data end!");
})

tcpClient.on("error", () => {
    console.log("tcpClient error!");
})
