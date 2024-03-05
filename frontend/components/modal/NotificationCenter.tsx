import { useState } from "react";
import { Bell, Info } from "react-feather";

import { Invitation } from "@/components/group/Invitation";
import { Modal, ModalProps } from "@/components/modal/modal";
import { useInvitations } from "@/util/swr";

import { Text } from "../Text";

/**
 * Component for displaying an icon that opens the notification center
 */
export const NotificationCenter = () => {
	const { invitations } = useInvitations();

	const [showNotificationCenter, setShowNotificationCenter] =
		useState<boolean>(false);

	return (
		<>
			<NotificationCenterModal
				visible={showNotificationCenter}
				setVisible={setShowNotificationCenter}
			/>
			<div className="relative">
				<Bell
					className="cursor-pointer"
					onClick={() => {
						setShowNotificationCenter(true);
					}}
				/>
				{invitations && (
					<div className="absolute right-[-0.1rem] top-[-0.15rem] h-3 w-3 flex-none rounded-full bg-kiokuRed">
						<div className="absolute h-full w-full animate-[ping_0.8s_ease-out_3] rounded-full bg-kiokuRed" />
					</div>
				)}
			</div>
		</>
	);
};

interface NotificationCenterModalProps {
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * Modal for creating decks
 */
export const NotificationCenterModal = ({
	className = "",
	setVisible,
	...props
}: NotificationCenterModalProps & ModalProps) => {
	const { invitations } = useInvitations();

	return (
		<Modal header="Notification Center" setVisible={setVisible} {...props}>
			{invitations ? (
				invitations.map((invitation) => (
					<Invitation invitation={invitation} />
				))
			) : (
				<div className="flex items-center justify-center space-x-3 rounded-md border border-dashed border-gray-500 p-10 text-gray-800">
					<Info />
					<Text>You don't have any pending messages</Text>
				</div>
			)}
		</Modal>
	);
};
