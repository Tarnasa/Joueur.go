var net = require("net")
var clients = [], i;

var server = net.createServer(function (stream) {
  stream.setEncoding("utf8");
  stream.on("connect", function () {
    clients.push(stream);
  });
  stream.on("data", function (data) {
    console.log(`got: '${data}'`);
  	console.log("> echoing to " + clients.length + " clients.");
  	for(i = 0; i < clients.length; i++){
  		if(clients[i] != stream){
  			clients[i].write(data + "\0");
  		}
  	}
  });
  stream.on("end", function () {
    stream.end();
    for(i = 0; i < clients.length; i++){
    	if(clients[i] == stream){
    		clients.splice(i, 1);
    	}
    }
    console.log('removed client');
  });
});
process.on('uncaughtException', function(exception){
	console.log("uncaught, just ignoring");
});
console.log("Server's running on 192.168.1.198:3000");
server.listen(3000);
