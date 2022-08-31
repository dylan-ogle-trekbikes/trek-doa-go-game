var express = require("express");
var app = express();

app.use(express.static("./dist"));

app.get("/", function (req, res) {
  res.sendFile("test.html", { root: __dirname + "/dist" });
});

app.listen(4000, () => {
  console.log(`listening on port 4000`);
});
