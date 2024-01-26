export function getOS() {
	const userAgent = navigator.userAgent;
	if (/android/i.test(userAgent)) {
		return "android";
	}
	if (/iPad|iPhone|iPod/.test(userAgent)) {
		return "ios";
	}
	return "unknown";
}

export function getBrowser() {
	const userAgent = navigator.userAgent;
	if (/samsungbrowser/i.test(userAgent)) {
		return "samsung";
	}
	if (/chrome/i.test(userAgent)) {
		return "chrome";
	}
	if (/safari/i.test(userAgent)) {
		return "safari";
	}
	if (/firefox/i.test(userAgent)) {
		return "firefox";
	}
	return "unknown";
}
