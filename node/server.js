var os = require("os");
var http = require('http');
var hostname = os.hostname();

var handleRequest = function(request, response) {
    console.log('Received request for URL: ' + request.url);
    response.writeHead(200);
    response.end('Hello World from ' + hostname + "!");
};
var www = http.createServer(handleRequest);
www.listen(8080);
