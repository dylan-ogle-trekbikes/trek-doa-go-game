const e = require("express");
var express = require("express");
const { resolve } = require("path");
var app = express();

app.use(express.static("./dist"));

app.get("/", function (req, res) {
  res.sendFile("test.html", { root: __dirname + "/dist" });
});

const port = process.env.PORT || 4000;

app.listen(port, () => {
  console.log(`listening on port ${port}`);
});
