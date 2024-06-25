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

const savePageContentToFile = (pageContent, query) => {
	const fs = require("fs");
	fs.writeFileSync(
		`tmp/profile-${query || new Date().getTime()}.html`,
		pageContent,
		(err) =>
			err ? console.error(err) : `new page content schema saved! @${query}`
	);
};

const saveJsonToFile = (json, query) => {
	const fs = require("fs");
	fs.writeFileSync(
		`tmp/${query || new Date().getTime()}.json`,
		JSON.stringify(json),
		(err) =>
			err ? console.error(err) : `new page content schema saved! @${query}`
	);
};

module.exports = {
	initBrowser,
	savePageContentToFile,
	saveJsonToFile,
};
