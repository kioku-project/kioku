import useSWR from "swr";

import { Group as GroupType } from "../../types/Group";
import { User } from "../../types/User";
import { authedFetch } from "../../util/reauth";
import Member from "./Member";

interface MemberListProps {
	/**
	 * group entity
	 */
	group: GroupType;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a group of users
 */
export default function MemberList({
	group,
	className = "",
}: Readonly<MemberListProps>) {
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
	);
	const { data: invitedUser } = useSWR(
		`/api/groups/${group.groupID}/members/invitations`,
		fetcher
	);

	return (
		<div id={group.groupID} className={`flex flex-col ${className}`}>
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
