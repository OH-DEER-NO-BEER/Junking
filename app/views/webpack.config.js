const VueLoaderPlugin = require("vue-loader/lib/plugin");

module.exports = {
	mode: "development",
	module: {
		rules: [
			{
				test: /\.vue$/,
				loader: "vue-loader",
			},
			{
				test: /\.js$/,
				loader: "babel-loader",
			},
			{
				test: /\.css$/,
				use: ["vue-style-loader", "css-loader"],
			},
			{
				test: /\.s(c|a)ss$/,
				use: [
					"vue-style-loader",
					"css-loader",
					{
						loader: "sass-loader",
						// Requires sass-loader@^7.0.0
						options: {
							implementation: require("sass"),
							indentedSyntax: true, // optional
						},
						// Requires >= sass-loader@^8.0.0
						options: {
							implementation: require("sass"),
							sassOptions: {
								indentedSyntax: true, // optional
							},
						},
					},
				],
			},
		],
	},
	plugins: [new VueLoaderPlugin()],
	resolve: {
		extensions: [".vue", ".js"],
		alias: {
			vue$: "vue/dist/vue.esm.js",
		},
	},
};
