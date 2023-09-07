import useSWR from "swr";
import { authedFetch } from "../../../util/reauth";
import DeckOverview from "../../deck/DeckOverview";
import { Group } from "../../../types/Group";

interface InvitationsTabProps {
	/**
	 * List of all invitations
	 */
	invitations: Pick<Group, "groupID" | "groupName">[];
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
						groupRole: "INVITED",
					}}
				/>
			))}
		</div>
	);
};
