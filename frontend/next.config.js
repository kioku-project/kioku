/** @type {import('next').NextConfig} */
const nextConfig = {
	reactStrictMode: true,
	swcMinify: true,
	experimental: {
		swcPlugins: [["@lingui/swc-plugin", {}]],
	},
	output: "standalone",
	i18n: {
		locales: ["en", "de"],
		defaultLocale: "en",
	},
	async rewrites() {
		return [
			{
				source: "/api/:path*",
				destination: "http://frontend_proxy:80/api/:path*",
			},
		];
	},
};

const withPWA = require("next-pwa")({
	dest: "public",
	register: false,
	skipWaiting: false,
});

module.exports = withPWA(nextConfig);
