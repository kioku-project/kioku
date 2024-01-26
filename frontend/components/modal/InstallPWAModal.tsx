import { Trans, msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import Link from "next/link";
import {
	Download,
	Home,
	Menu,
	MoreVertical,
	Plus,
	PlusSquare,
	Share,
} from "react-feather";

import { Text } from "@/components/Text";
import { Button } from "@/components/input/Button";
import { Modal, ModalProps } from "@/components/modal/modal";
import { getBrowser, getOS } from "@/util/utils";

export const InstallPWAModal = ({ setVisible, ...props }: ModalProps) => {
	const { _ } = useLingui();
	const os = getOS();
	const browser = getBrowser();

	return (
		<Modal
			header={_(msg`How to install Kioku`)}
			setVisible={setVisible}
			{...props}
		>
			<div className="space-y-5">
				<Text textSize="xs">
					{os === "unknown" || browser === "unknown" ? (
						<>
							<Trans>
								You are using an unknown browser or operating
								system, so we cannot provide you with
								instructions on how to install Kioku on your
								device. We would be happy if you report this
								issue
							</Trans>{" "}
							<Link
								href="https://github.com/kioku-project/kioku/issues/new/choose"
								className="underline"
							>
								<Trans> here </Trans>
							</Link>{" "}
							<Trans>
								so that we can support even more devices in the
								future. We recommend to use Chrome on Android
								and Safari on iOS for the best experience.
							</Trans>
						</>
					) : (
						<Trans>
							Install Kioku for a native experience and to receive
							notifications.
						</Trans>
					)}
				</Text>
				{os === "ios" && browser === "safari" && (
					<IosSafariInstructions />
				)}
				{os === "android" && browser === "chrome" && (
					<AndroidChromeInstructions />
				)}
				{os === "android" && browser === "samsung" && (
					<AndroidSamsungInstructions />
				)}
				{os === "android" && browser === "firefox" && (
					<AndroidFirefoxInstructions />
				)}
				<div className="flex flex-row justify-end space-x-1">
					<Button
						buttonStyle="secondary"
						onClick={() => setVisible(false)}
					>
						<Trans>Done</Trans>
					</Button>
				</div>
			</div>
		</Modal>
	);
};

const IosSafariInstructions = () => {
	return (
		<>
			<div className="flex flex-row items-center space-x-3">
				<Share size={20} />
				<Text textSize="5xs">
					<Trans>Tap on share</Trans>
				</Text>
			</div>
			<div className="flex flex-row items-center space-x-3">
				<PlusSquare size={20} />
				<Text textSize="5xs">
					<Trans>Select &quot;Add to Home Screen&quot;</Trans>
				</Text>
			</div>
		</>
	);
};

const AndroidChromeInstructions = () => {
	return (
		<>
			<div className="flex flex-row items-center space-x-3">
				<MoreVertical size={20} />
				<Text textSize="5xs">
					<Trans>Open menu</Trans>
				</Text>
			</div>
			<div className="flex flex-row items-center space-x-3">
				<Download size={20} />
				<Text textSize="5xs">
					<Trans>Select &quot;Install app&quot;</Trans>
				</Text>
			</div>
		</>
	);
};

const AndroidSamsungInstructions = () => {
	return (
		<>
			<Text textSize="5xs" className="text-kiokuRed">
				<Trans>
					Since the Samsung Internet Browser is not fully supporting
					PWAs, we recommend to use Chrome for the best experience.
				</Trans>
			</Text>
			<div className="flex flex-row items-center space-x-3">
				<Menu size={20} />
				<Text textSize="5xs">
					<Trans>Open menu</Trans>
				</Text>
			</div>
			<div className="flex flex-row items-center space-x-3">
				<Plus size={20} />
				<Text textSize="5xs">
					<Trans>Click on &quot;Add page to&quot;</Trans>
				</Text>
			</div>
			<div className="flex flex-row items-center space-x-3">
				<Home size={20} />
				<Text textSize="5xs">
					<Trans>Select &quot;Home screen&quot;</Trans>
				</Text>
			</div>
		</>
	);
};

const AndroidFirefoxInstructions = () => {
	return (
		<>
			<Text textSize="5xs" className="text-kiokuRed">
				<Trans>
					Since Firefox is not fully supporting PWAs, we recommend to
					use Chrome for the best experience.
				</Trans>
			</Text>
			<div className="flex flex-row items-center space-x-3">
				<MoreVertical size={20} />
				<Text textSize="5xs">
					<Trans>Open menu</Trans>
				</Text>
			</div>
			<div className="flex flex-row items-center space-x-3">
				<Download size={20} />
				<Text textSize="5xs">
					<Trans>Select &quot;Install&quot;</Trans>
				</Text>
			</div>
		</>
	);
};
