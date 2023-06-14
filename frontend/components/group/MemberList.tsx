import Member from "./Member";
import useSWR from "swr";
import { authedFetch } from "../../util/reauth";

interface MemberListProps {
	/**
	 * groupID
	 */
	groupID: string;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a group of users
 */
export default function MemberList({ groupID, className }: MemberListProps) {
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: user } = useSWR(`/api/groups/${groupID}/members`, fetcher);
	const { data: requestedUser } = useSWR(
		`/api/groups/${groupID}/members/requests`,
		fetcher
	); //TODO: admissionID löschen
	const { data: invitedUser } = useSWR(
		`/api/groups/${groupID}/members/invitations`,
		fetcher
	);

	console.log(requestedUser);

	return (
		<div id={groupID} className={`flex flex-col ${className ?? ""}`}>
			<div className="snap-y overflow-y-auto">
				{user?.users &&
					user.users.map((user: User) => (
						<Member
							className="snap-center"
							key={user.userID}
							user={{ ...user, groupID: groupID }}
						/>
					))}
				{requestedUser?.memberRequests &&
					requestedUser.memberRequests.map(
						(requestedUser: {
							userID: string;
							userName: string;
						}) => (
							<Member
								className="snap-center"
								key={requestedUser.userID}
								user={{
									...requestedUser,
									groupRole: "REQUESTED",
									groupID: groupID,
								}}
							/>
						)
					)}
				{invitedUser?.groupInvitations &&
					invitedUser.groupInvitations.map(
						(invitedUser: {
							admissionID: string;
							userID: string;
							userName: string;
						}) => (
							<Member
								className="snap-center"
								key={invitedUser.userID}
								user={{
									...invitedUser,
									groupRole: "INVITED",
									groupID: groupID,
								}}
							/>
						)
					)}
				<Member
					user={{ userID: "", userName: "", groupID: groupID }}
				></Member>
			</div>
		</div>
	);
}
