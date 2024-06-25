//Dependencies
require('dotenv').config();
const express = require("express");
const app = require("express")();
const http = require("http").Server(app);
const config = require("./config/config.js");
const bodyParser = require("body-parser");
const mongoose = require("mongoose");
const cors = require("cors");
let { initBrowser } = require("./helpers/playwright.js");

app.use(express.static("public"));
app.use(bodyParser.urlencoded({ extended: true }));

process.on("SIGINT", function () {
  mongoose.disconnect(function () {
    console.log("Mongoose disconnected on app termination");
    process.exit(0);
  });
});

(async () => {
  global.browser = await initBrowser();

  require("./routes/route.js")(app);

  app.use(function (req, res, next) {
    res.sendStatus(404);
  });

  http.listen(config.port, function () {
    console.log("listening on *:", config.port);
  });
})();
