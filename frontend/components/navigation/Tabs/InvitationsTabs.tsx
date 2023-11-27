import { Invitation } from "../../../types/Invitation";
import DeckList from "../../deck/DeckList";

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
