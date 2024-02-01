import DeckList from "@/components/deck/DeckList";
import { Invitation } from "@/types/Invitation";

interface InvitationsTabProps {
	/**
	 * List of all invitations
	 */
	invitations: Invitation[];
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the InvitationsTab
 */
export const InvitationsTab = ({
	invitations,
	className = "",
}: InvitationsTabProps) => {
	return (
		<div className={`${className}`}>
			{invitations?.map((invitation) => (
				<DeckList
					key={invitation.groupID}
					header={invitation.groupName}
				/>
			))}
		</div>
	);
};
