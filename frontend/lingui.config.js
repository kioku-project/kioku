/** @type {import('@lingui/conf').LinguiConfig} */
module.exports = {
	locales: ["en", "de"],
	catalogs: [
		{
			path: "<rootDir>/locales/{locale}/messages",
			include: ["<rootDir>"],
			exclude: ["**/node_modules/**"],
		},
	],
	format: "po",
	sourceLocale: "en",
	compileNamespace: "ts",
};
