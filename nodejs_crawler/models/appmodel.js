var mongoose = require('mongoose');
var config = require('../config/config.js');

mongoose.connect(config.database.crawlerUrl, { 
    promiseLibrary: global.Promise,
    useNewUrlParser: true,
    useUnifiedTopology: true
});
mongoose.set('useCreateIndex', true);

module.exports = mongoose;