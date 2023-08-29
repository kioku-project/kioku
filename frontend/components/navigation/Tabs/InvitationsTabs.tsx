import useSWR from "swr";
import { authedFetch } from "../../../util/reauth";
import DeckOverview from "../../deck/DeckOverview";
import { Group } from "../../../types/Group";

interface InvitationsTabProps {
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the InvitationsTab
 */
export const InvitationsTab = ({ className = "" }: InvitationsTabProps) => {
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: invitations } = useSWR<{
		groupInvitation: Pick<Group, "groupID" | "groupName">[];
	}>(`/api/user/invitations`, fetcher);

	return (
		<div className={`${className}`}>
			{invitations?.groupInvitation?.map((invitation) => (
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
