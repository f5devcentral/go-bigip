var f5 = require('f5-nodejs');
var ilx = new f5.ILXServer();

ilx.addMethod('greeter', function(req, res) {
  var arg = req.params()[0];

  res.reply('Hello ' + arg);
});

ilx.listen();
