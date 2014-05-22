
  var io = require('socket.io-client')

  socket = io.connect('http://localhost:3000');
  
  //console.log( socket )
  socket.on("connect", function(){
    socket.emit("news", "this is title", "this is body", 1)
  })
  socket.on("error", function (error) {
    console.log("error", error);
  })

  socket.on("connect_timeout",function  (argument) {
    console.log( argument );
  })


  socket.on("news", function(message, urgency){
    console.log(message + urgency);
    socket.emit("ping")
  })
  socket.on("pong", function() {
    console.log("got pong")
  })
  socket.of("/pol").on("news", function(message, urgency){
    console.log(message + urgency);
    socket.emit("ping")
  })
  socket.of("/pol").on("pong", function() {
    console.log("got pong")
  })
  socket.on("disconnect", function() {
    alert("You have disconnected from the server")
  })
  var pol = io.connect("http://localhost:3000/pol");
  pol.on("pong", function() {
    console.log("got pong from pol")
  })
  pol.on("news", function(message, urgency){
    console.log(message + urgency);
    socket.emit("ping")
  })