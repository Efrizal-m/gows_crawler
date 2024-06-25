const config = require("../config/config.js");
const Logger = require("../helpers/logger.js");
const NewsDetik = require("../models/newsDetik.js");

module.exports = {
	crawlDetik: async () => {
		let context = undefined;
		let page = undefined;
		let webData;
		try {
			context = await global.browser.newContext();
			page = await context.newPage();

			await Promise.all([
				page.goto(`https://www.detik.com`, {
					timeout: 600000,
					waitUntil: 'domcontentloaded',
				}),
				page
					.waitForResponse(
						(response) => response.url().includes(`/rech.detik.com/article-recommendation/wp/`) && response.status() === 200,
						{
							timeout: 240000,
						}
					)
					.then(async (webProfileInfoResponse) => {
						const webProifleInfojson = await webProfileInfoResponse.json();
						console.log("web data",webProifleInfojson);
						webData = webProifleInfojson.body;
					}),
			]);

			try {
				for await (const data of webData) {
					let checkNews = await NewsDetik.findOne({ id: data.id });
					if (!checkNews) checkNews = new NewsDetik(data);
					const saveNews = await checkNews.save();
					Logger.info(`Save data news ${data.id}`);		
				}
			} catch (e) {
				console.error(e)
				return null;
			}
			return webData				
			await browser.close()
		} catch (error) {
			Logger.err(error);
			return { meta: "bad-gateway" };
		}
	}
}
