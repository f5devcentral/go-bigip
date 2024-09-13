var f5 = require("f5-nodejs");

var ilx = new f5.ILXServer();

ilx.addMethod("getCredentials", function (req, res) {
  const [user] = req.params();
  const arc = new CyberArc();
  const credentials = arc.getCredentials(user);
  res.reply(credentials);
});

ilx.listen();

/**
 * @class CyberArc
 */
class CyberArc {
  /**
   * @method getCredentials
   * @param {string} user
   * @returns {Array} [username, password]
   */
  getCredentials(username) {
    return ["admin", "password"];
  }
}
