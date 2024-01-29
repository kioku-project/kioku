import { mapIgnoreCase } from "@/util/utils";

export enum Platform {
	DESKTOP,
	MOBILE,
	UNKNOWN,
}

export enum OperatingSystem {
	ANDROID,
	IOS,
	UNKNOWN,
}

export enum Browser {
	SAMSUNG,
	CHROME,
	FIREFOX,
	SAFARI,
	UNKNOWN,
}

export function getPlatform(userAgent: string): Platform {
	const os = getOperatingSystem(userAgent);
	if (os === OperatingSystem.ANDROID || os === OperatingSystem.IOS) {
		return Platform.MOBILE;
	}
	return Platform.UNKNOWN;
}

export function getOperatingSystem(userAgent: string): OperatingSystem {
	return (
		mapIgnoreCase<OperatingSystem>(
			{
				android: OperatingSystem.ANDROID,
				iphone: OperatingSystem.IOS,
				ipad: OperatingSystem.IOS,
				ipod: OperatingSystem.IOS,
			},
			userAgent
		) ?? OperatingSystem.UNKNOWN
	);
}

export function getBrowser(userAgent: string): Browser {
	return (
		mapIgnoreCase<Browser>(
			{
				samsungbrowser: Browser.SAMSUNG,
				chrome: Browser.CHROME,
				safari: Browser.SAFARI,
				firefox: Browser.FIREFOX,
			},
			userAgent
		) ?? Browser.UNKNOWN
	);
}
