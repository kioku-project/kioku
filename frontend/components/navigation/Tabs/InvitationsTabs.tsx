import { Invitation } from "../../../types/Invitation";
import DeckOverview from "../../deck/DeckList";

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
				<DeckOverview
					key={invitation.groupID}
					group={{
						...invitation,
						isDefault: false,
						groupRole: "INVITED",
					}}
				/>
			))}
		</div>
	);
};
