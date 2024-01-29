import { Trans, msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import Link from "next/link";
import { ReactElement } from "react";
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
import {
	Browser,
	OperatingSystem,
	getBrowser,
	getOperatingSystem,
} from "@/util/client";

export const InstallPWAModal = ({ setVisible, ...props }: ModalProps) => {
	const { _ } = useLingui();
	const os = getOperatingSystem(navigator.userAgent);
	const browser = getBrowser(navigator.userAgent);

	return (
		<Modal
			header={_(msg`How to install Kioku`)}
			setVisible={setVisible}
			{...props}
		>
			<div className="space-y-5">
				<Text textSize="xs">
					{os === OperatingSystem.UNKNOWN ||
					browser === Browser.UNKNOWN ? (
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
								<Trans>here</Trans>
							</Link>{" "}
							<Trans>
								so that we can support even more devices in the
								future. We recommend using Chrome on Android and
								Safari on iOS for the best experience.
							</Trans>
						</>
					) : (
						<Trans>
							You first need to install Kioku before you can
							subscribe and receive notifications. Follow the
							instructions below to install the app and enjoy a
							native experience.
						</Trans>
					)}
				</Text>
				{os === OperatingSystem.IOS && browser === Browser.SAFARI && (
					<IosSafariInstructions />
				)}
				{os === OperatingSystem.ANDROID &&
					browser === Browser.CHROME && <AndroidChromeInstructions />}
				{os === OperatingSystem.ANDROID &&
					browser === Browser.SAMSUNG && (
						<AndroidSamsungInstructions />
					)}
				{os === OperatingSystem.ANDROID &&
					browser === Browser.FIREFOX && (
						<AndroidFirefoxInstructions />
					)}
				<Button
					buttonStyle="secondary"
					onClick={() => setVisible(false)}
					className="ml-auto"
				>
					<Trans>Done</Trans>
				</Button>
			</div>
		</Modal>
	);
};

const IosSafariInstructions = () => {
	const { _ } = useLingui();
	return (
		<>
			<Instruction
				icon={<Share size={20} />}
				text={_(msg`Tap on share`)}
			/>
			<Instruction
				icon={<PlusSquare size={20} />}
				text={_(msg`Select "Add to Home Screen"`)}
			/>
		</>
	);
};

const AndroidChromeInstructions = () => {
	const { _ } = useLingui();
	return (
		<>
			<Instruction
				icon={<MoreVertical size={20} />}
				text={_(msg`Open menu`)}
			/>
			<Instruction
				icon={<Download size={20} />}
				text={_(msg`Select "Install app"`)}
			/>
		</>
	);
};

const AndroidSamsungInstructions = () => {
	const { _ } = useLingui();
	return (
		<>
			<Text textSize="5xs" className="text-kiokuRed">
				<Trans>
					Since the Samsung Internet Browser is not fully supporting
					PWAs, we recommend using Chrome for the best experience.
				</Trans>
			</Text>
			<Instruction icon={<Menu size={20} />} text={_(msg`Open menu`)} />
			<Instruction
				icon={<Plus size={20} />}
				text={_(msg`Click on "Add page to"`)}
			/>
			<Instruction
				icon={<Home size={20} />}
				text={_(msg`Select "Home screen"`)}
			/>
		</>
	);
};

const AndroidFirefoxInstructions = () => {
	const { _ } = useLingui();
	return (
		<>
			<Text textSize="5xs" className="text-kiokuRed">
				<Trans>
					Since Firefox is not fully supporting PWAs, we recommend
					using Chrome for the best experience.
				</Trans>
			</Text>
			<Instruction
				icon={<MoreVertical size={20} />}
				text={_(msg`Open menu`)}
			/>
			<Instruction
				icon={<Download size={20} />}
				text={_(msg`Select "Install"`)}
			/>
		</>
	);
};

interface InstructionProps {
	icon: ReactElement;
	text: string;
}

const Instruction = ({ icon, text }: InstructionProps) => {
	return (
		<div className="flex flex-row items-center space-x-3">
			{icon}
			<Text textSize="5xs">{text}</Text>
		</div>
	);
};
