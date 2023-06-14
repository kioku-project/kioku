import Member from "./Member";
import useSWR from "swr";
import { authedFetch } from "../../util/reauth";
import { User } from "../../types/User";
import { Group } from "../../types/Group";

interface MemberListProps {
	/**
	 * group entity
	 */
	group: Group;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a group of users
 */
export default function MemberList({ group, className }: MemberListProps) {
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: user } = useSWR(
		`/api/groups/${group.groupID}/members`,
		fetcher
	);
	const { data: requestedUser } = useSWR(
		`/api/groups/${group.groupID}/members/requests`,
		fetcher
	); //TODO: admissionID löschen
	const { data: invitedUser } = useSWR(
		`/api/groups/${group.groupID}/members/invitations`,
		fetcher
	);

	console.log(requestedUser);

	return (
		<div id={group.groupID} className={`flex flex-col ${className ?? ""}`}>
			<div className="snap-y overflow-y-auto">
				{user?.users &&
					user.users.map((user: User) => (
						<Member
							className="snap-center"
							key={user.userID}
							user={{ ...user, groupID: group.groupID }}
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
									groupID: group.groupID,
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
									groupID: group.groupID,
								}}
							/>
						)
					)}
				{group.groupRole == "ADMIN" && (
					<Member
						user={{
							userID: "",
							userName: "",
							groupID: group.groupID,
						}}
					></Member>
				)}
			</div>
		</div>
	);
}
