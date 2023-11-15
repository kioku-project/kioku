import { Group as GroupType } from "../../types/Group";
import { User } from "../../types/User";
import { useInvitedUser, useMembers, useRequestedUser } from "../../util/swr";
import Member from "./Member";

interface MemberListProps {
	/**
	 * Group entity
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
	const { members } = useMembers(group.groupID);
	const { invitedUser } = useInvitedUser(group.groupID);
	const { requestedUser } = useRequestedUser(group.groupID);

	return (
		<div id={group.groupID} className={`flex flex-col ${className}`}>
			<div className="snap-y overflow-y-auto">
				{members?.map((user: User) => (
					<Member
						className="snap-center"
						key={user.userID}
						user={{ ...user, groupID: group.groupID }}
					/>
				))}
				{requestedUser?.map(
					(requestedUser: { userID: string; userName: string }) => (
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
				{invitedUser?.map((invitedUser) => (
					<Member
						className="snap-center"
						key={invitedUser.userID}
						user={{
							...invitedUser,
							groupRole: "INVITED",
							groupID: group.groupID,
						}}
					/>
				))}
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
