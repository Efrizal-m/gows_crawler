var mongoose = require("./appmodel.js");

var postSchema = new mongoose.Schema(
  {
    id: { type: String, required: true },
    articleurl: String,
    categoryauto: String,
    desktopurl: String,
    type: String,
    idparentkanal: Number,
    imageurl: String,
    mobileurl: String,
    parent: Number,
    penulis: String,
    publishdate: String,
    resume: String,
    tag: String,
    title: String,
    idarticletype: String,
    unixtime: Number
  },
  { timestamps: true }
);
var Post = mongoose.model("NewsDetik", postSchema, "detik");

module.exports = Post;
