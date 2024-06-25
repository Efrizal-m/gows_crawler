const { crawlDetik } = require("../controllers/crawlDetik.js");

module.exports = function (app) {
  app.get("/crawl/detik", async function (req, res) {
    const result = await crawlDetik();
    res.json({ result });
  })
};
