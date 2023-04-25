/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  output: 'standalone',
  i18n: {
    locales: ["en"],
    defaultLocale: "en",
  },
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://envoy:8081/api/:path*',
      },
    ];
  },
}

module.exports = nextConfig
