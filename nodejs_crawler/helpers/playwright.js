const { chromium } = require("playwright");
const config = require("../config/config.js");

/** @return {import('playwright').Browser} */
const initBrowser = async () => {
	try {
		let option = { headless: true, devtools: true };
		if (config.proxyBrowser) option.proxy = {
			server: config.proxyBrowser.server,
			username: config.proxyBrowser.username,
			password: config.proxyBrowser.password
		}
		if (config.headlessBrowser === false) option.headless = false;
		return await chromium.launch(option);
	} catch (error) {
		console.error(error);
		process.exit(1);
	}
};

module.exports = {
	initBrowser,
};
