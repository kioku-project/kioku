import { Trans } from "@lingui/macro";
import clsx from "clsx";
import Link from "next/link";
import { Check, X } from "react-feather";

import { Button } from "@/components/input/Button";
import { Invitation as InvitationType } from "@/types/Invitation";
import { declineGroupInvitation, sendGroupRequest } from "@/util/api";

import { Text } from "../Text";

interface InvitationProps {
	/**
	 * Invitation
	 */
	invitation: InvitationType;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 *  Component for displaying a group invitation in the notification center
 */
export const Invitation = ({ invitation, className }: InvitationProps) => {
	return (
		<div className={clsx("flex flex-row space-x-3", className)}>
			<div className="h-8 w-8 flex-none rounded-full bg-kiokuLightBlue" />
			<div className="w-full">
				<Text textSize="5xs" className="font-semibold text-black">
					<Trans>Group Invitation</Trans>
				</Text>
				<Text textSize="5xs" className="text-gray-500">
					<Trans>
						You have been invited to join{" "}
						<Link
							href={`/group/${invitation.groupID}`}
							className="underline hover:text-black"
						>
							{invitation.groupName}
						</Link>
					</Trans>
				</Text>
			</div>
			<div className="flex flex-row space-x-1">
				<Button
					buttonSize="px-3"
					buttonStyle="tertiary"
					buttonIcon={<X strokeWidth={3}></X>}
					onClick={() => declineGroupInvitation(invitation.groupID)}
				/>
				<Button
					buttonSize="px-3 py-1"
					buttonStyle="secondary"
					buttonIcon={<Check strokeWidth={3}></Check>}
					onClick={() => sendGroupRequest(invitation.groupID)}
				/>
			</div>
		</div>
	);
};
