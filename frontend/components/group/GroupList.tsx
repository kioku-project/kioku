import Link from "next/link";

import DeckList from "@/components/deck/DeckList";
import { Group as GroupType } from "@/types/Group";

interface GroupListProps {
	/*
	 * Groups
	 */
	groups: GroupType[];
	/*
	 * Additional classes
	 */
	className?: string;
}

export default function GroupList({ groups, className }: GroupListProps) {
	return (
		<div className={`space-y-3 ${className}`}>
			{groups.map((group: GroupType) => {
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
