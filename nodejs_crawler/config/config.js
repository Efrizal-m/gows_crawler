
module.exports = {
	port: Number(process.env.PORT || 0) || 3006,
	database: {
		crawlerUrl: process.env.MONGODB || "mongodb://localhost:27017/ivosight",
	},
	headlessBrowser: process.env.HEADLESS_BROWSER === 'true',
} 