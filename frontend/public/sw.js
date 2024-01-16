if (!self.define) {
	let e,
		s = {};
	const i = (i, a) => (
		(i = new URL(i + ".js", a).href),
		s[i] ||
			new Promise((s) => {
				if ("document" in self) {
					const e = document.createElement("script");
					(e.src = i), (e.onload = s), document.head.appendChild(e);
				} else (e = i), importScripts(i), s();
			}).then(() => {
				let e = s[i];
				if (!e)
					throw new Error(`Module ${i} didn’t register its module`);
				return e;
			})
	);
	self.define = (a, c) => {
		const n =
			e ||
			("document" in self ? document.currentScript.src : "") ||
			location.href;
		if (s[n]) return;
		let t = {};
		const r = (e) => i(e, n),
			o = { module: { uri: n }, exports: t, require: r };
		s[n] = Promise.all(a.map((e) => o[e] || r(e))).then(
			(e) => (c(...e), t)
		);
	};
}
define(["./workbox-7c2a5a06"], function (e) {
	"use strict";
	importScripts(),
		self.skipWaiting(),
		e.clientsClaim(),
		e.precacheAndRoute(
			[
				{
					url: "/_next/static/0TrjcPL_syZUg_Qd74P44/_buildManifest.js",
					revision: "ffc674da68150e4cc26933cfd68c9a6f",
				},
				{
					url: "/_next/static/0TrjcPL_syZUg_Qd74P44/_ssgManifest.js",
					revision: "b6652df95db52feb4daf4eca35380933",
				},
				{
					url: "/_next/static/chunks/386.ff2c0dce6ef7f62a.js",
					revision: "ff2c0dce6ef7f62a",
				},
				{
					url: "/_next/static/chunks/401-42b839cb9f0bc285.js",
					revision: "42b839cb9f0bc285",
				},
				{
					url: "/_next/static/chunks/578-2e8cc1449169ec98.js",
					revision: "2e8cc1449169ec98",
				},
				{
					url: "/_next/static/chunks/683.3c8f97b1c5d2bbca.js",
					revision: "3c8f97b1c5d2bbca",
				},
				{
					url: "/_next/static/chunks/734-cbfa3b35476b4d1a.js",
					revision: "cbfa3b35476b4d1a",
				},
				{
					url: "/_next/static/chunks/762-bb7a7b8ebb21a774.js",
					revision: "bb7a7b8ebb21a774",
				},
				{
					url: "/_next/static/chunks/framework-0c7baedefba6b077.js",
					revision: "0c7baedefba6b077",
				},
				{
					url: "/_next/static/chunks/main-309167e615f41835.js",
					revision: "309167e615f41835",
				},
				{
					url: "/_next/static/chunks/pages/_app-75269e7aaac127af.js",
					revision: "75269e7aaac127af",
				},
				{
					url: "/_next/static/chunks/pages/_error-ee5b5fb91d29d86f.js",
					revision: "ee5b5fb91d29d86f",
				},
				{
					url: "/_next/static/chunks/pages/deck/%5Bid%5D-5356bb2184b938e2.js",
					revision: "5356bb2184b938e2",
				},
				{
					url: "/_next/static/chunks/pages/deck/%5Bid%5D/learn-7c79b398344d9508.js",
					revision: "7c79b398344d9508",
				},
				{
					url: "/_next/static/chunks/pages/features-54d238a86ad9c41c.js",
					revision: "54d238a86ad9c41c",
				},
				{
					url: "/_next/static/chunks/pages/group/%5Bid%5D-3a8b07e7f0138df7.js",
					revision: "3a8b07e7f0138df7",
				},
				{
					url: "/_next/static/chunks/pages/home-46864f456089d7c1.js",
					revision: "46864f456089d7c1",
				},
				{
					url: "/_next/static/chunks/pages/index-a34d057aff37e4eb.js",
					revision: "a34d057aff37e4eb",
				},
				{
					url: "/_next/static/chunks/pages/login-b1d83236fbb06377.js",
					revision: "b1d83236fbb06377",
				},
				{
					url: "/_next/static/chunks/polyfills-c67a75d1b6f99dc8.js",
					revision: "837c0df77fd5009c9e46d446188ecfd0",
				},
				{
					url: "/_next/static/chunks/webpack-5fbbe0967e3f7a7d.js",
					revision: "5fbbe0967e3f7a7d",
				},
				{
					url: "/_next/static/css/431944509084d071.css",
					revision: "431944509084d071",
				},
				{
					url: "/_next/static/css/58dabbfb732e126e.css",
					revision: "58dabbfb732e126e",
				},
				{
					url: "/_next/static/media/05a31a2ca4975f99-s.woff2",
					revision: "f1b44860c66554b91f3b1c81556f73ca",
				},
				{
					url: "/_next/static/media/513657b02c5c193f-s.woff2",
					revision: "c4eb7f37bc4206c901ab08601f21f0f2",
				},
				{
					url: "/_next/static/media/51ed15f9841b9f9d-s.woff2",
					revision: "bb9d99fb9bbc695be80777ca2c1c2bee",
				},
				{
					url: "/_next/static/media/c9a5bc6a7c948fb0-s.p.woff2",
					revision: "74c3556b9dad12fb76f84af53ba69410",
				},
				{
					url: "/_next/static/media/d6b16ce4a6175f26-s.woff2",
					revision: "dd930bafc6297347be3213f22cc53d3e",
				},
				{
					url: "/_next/static/media/ec159349637c90ad-s.woff2",
					revision: "0e89df9522084290e01e4127495fae99",
				},
				{
					url: "/_next/static/media/fd4db3eb5472fc27-s.woff2",
					revision: "71f3fcaf22131c3368d9ec28ef839831",
				},
				{
					url: "/_next/static/media/kioku-logo.f641a1d4.svg",
					revision: "352ee972a954584ebb50fb799922547b",
				},
				{
					url: "/favicon.ico",
					revision: "188a2323861b6d0105fd39d5e021d23a",
				},
				{
					url: "/github.png",
					revision: "eb94bb97c3410733ce017b184d314723",
				},
				{
					url: "/kioku-logo-horizontal.svg",
					revision: "5f87f378458661eeee96a54d98b71230",
				},
				{
					url: "/kioku-logo-title.svg",
					revision: "03e503cf05f44992295b7bfc7d95f354",
				},
				{
					url: "/kioku-logo.png",
					revision: "f2c70172c93869027b0420a9b0e7270e",
				},
				{
					url: "/kioku-logo.svg",
					revision: "352ee972a954584ebb50fb799922547b",
				},
				{
					url: "/loading_spinner.svg",
					revision: "da67c2b397c01205d708ef5f43039e2d",
				},
				{
					url: "/manifest.json",
					revision: "db1d1de3e04ec22d9315644742cf97a3",
				},
				{
					url: "/mockServiceWorker.js",
					revision: "c2a56920b0fe589c64486fa544d0d5c5",
				},
			],
			{ ignoreURLParametersMatching: [] }
		),
		e.cleanupOutdatedCaches(),
		e.registerRoute(
			"/",
			new e.NetworkFirst({
				cacheName: "start-url",
				plugins: [
					{
						cacheWillUpdate: async ({
							request: e,
							response: s,
							event: i,
							state: a,
						}) =>
							s && "opaqueredirect" === s.type
								? new Response(s.body, {
										status: 200,
										statusText: "OK",
										headers: s.headers,
								  })
								: s,
					},
				],
			}),
			"GET"
		),
		e.registerRoute(
			/^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,
			new e.CacheFirst({
				cacheName: "google-fonts-webfonts",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 4,
						maxAgeSeconds: 31536e3,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,
			new e.StaleWhileRevalidate({
				cacheName: "google-fonts-stylesheets",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 4,
						maxAgeSeconds: 604800,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,
			new e.StaleWhileRevalidate({
				cacheName: "static-font-assets",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 4,
						maxAgeSeconds: 604800,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,
			new e.StaleWhileRevalidate({
				cacheName: "static-image-assets",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 64,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\/_next\/image\?url=.+$/i,
			new e.StaleWhileRevalidate({
				cacheName: "next-image",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 64,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\.(?:mp3|wav|ogg)$/i,
			new e.CacheFirst({
				cacheName: "static-audio-assets",
				plugins: [
					new e.RangeRequestsPlugin(),
					new e.ExpirationPlugin({
						maxEntries: 32,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\.(?:mp4)$/i,
			new e.CacheFirst({
				cacheName: "static-video-assets",
				plugins: [
					new e.RangeRequestsPlugin(),
					new e.ExpirationPlugin({
						maxEntries: 32,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\.(?:js)$/i,
			new e.StaleWhileRevalidate({
				cacheName: "static-js-assets",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 32,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\.(?:css|less)$/i,
			new e.StaleWhileRevalidate({
				cacheName: "static-style-assets",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 32,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\/_next\/data\/.+\/.+\.json$/i,
			new e.StaleWhileRevalidate({
				cacheName: "next-data",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 32,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			/\.(?:json|xml|csv)$/i,
			new e.NetworkFirst({
				cacheName: "static-data-assets",
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 32,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			({ url: e }) => {
				if (!(self.origin === e.origin)) return !1;
				const s = e.pathname;
				return !s.startsWith("/api/auth/") && !!s.startsWith("/api/");
			},
			new e.NetworkFirst({
				cacheName: "apis",
				networkTimeoutSeconds: 10,
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 16,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			({ url: e }) => {
				if (!(self.origin === e.origin)) return !1;
				return !e.pathname.startsWith("/api/");
			},
			new e.NetworkFirst({
				cacheName: "others",
				networkTimeoutSeconds: 10,
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 32,
						maxAgeSeconds: 86400,
					}),
				],
			}),
			"GET"
		),
		e.registerRoute(
			({ url: e }) => !(self.origin === e.origin),
			new e.NetworkFirst({
				cacheName: "cross-origin",
				networkTimeoutSeconds: 10,
				plugins: [
					new e.ExpirationPlugin({
						maxEntries: 32,
						maxAgeSeconds: 3600,
					}),
				],
			}),
			"GET"
		);
});
