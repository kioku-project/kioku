import useSWR from "swr";
import { authedFetch } from "../../../util/reauth";
import DeckOverview from "../../deck/DeckOverview";

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
	const { data: invitations } = useSWR(`/api/user/invitations`, fetcher);

	return (
		<div className={`${className}`}>
			{invitations?.groupInvitation && (
				<DeckOverview
					group={{
						groupID: invitations.groupInvitation[0].groupID,
						groupName: invitations.groupInvitation[0].groupName,
						groupRole: "INVITED",
					}}
				></DeckOverview>
			)}
		</div>
	);
};
