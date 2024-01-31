import Link from "next/link";

import DeckList from "@/components/deck/DeckList";
import { Group as GroupType } from "@/types/Group";

interface GroupListProps {
	/*
	 * Groups
	 */
	groups: GroupType[];
	/**
	 * Filter groups
	 */
	filter?: string;
	/**
	 * Reverse group order
	 */
	reverse?: boolean;
	/*
	 * Additional classes
	 */
	className?: string;
}

export default function GroupList({
	groups,
	filter = "",
	reverse = false,
	className,
}: Readonly<GroupListProps>) {
	const filteredGroups = groups?.filter(
		(group) => !group.isDefault && group.groupName.includes(filter)
	);
	const sortedGroups = reverse ? filteredGroups?.reverse() : filteredGroups;

	return (
		<div className={`space-y-3 ${className}`}>
			{sortedGroups.map((group: GroupType) => {
				return (
					<Link key={group.groupID} href={`/group/${group.groupID}`}>
						<DeckList
							header={group.groupName}
							key={group.groupID}
						/>
					</Link>
				);
			})}
			<DeckList />
		</div>
	);
}
