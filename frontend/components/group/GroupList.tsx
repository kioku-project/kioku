import Link from "next/link";
import { Children, ReactNode, isValidElement, useMemo } from "react";

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
	/**
	 * SelectionField options
	 */
	children?: ReactNode;
}

/**
 * UI component for displaying a list of groups
 */
export default function GroupList({
	groups,
	filter = "",
	reverse = false,
	className,
	children,
}: Readonly<GroupListProps>) {
	const filteredGroups = useMemo(() => {
		const filteredGroups = groups?.filter(
			(group) =>
				!group.isDefault &&
				(group.groupName.toUpperCase().includes(filter) ||
					group.groupDescription?.toUpperCase().includes(filter))
		);
		return reverse ? filteredGroups?.toReversed() : filteredGroups;
	}, [groups, filter, reverse]);

	return (
		<div className={`space-y-3 ${className}`}>
			{filteredGroups.map((group: GroupType) => {
				return (
					<Link key={group.groupID} href={`/group/${group.groupID}`}>
						<DeckList
							header={group.groupName}
							key={group.groupID}
						/>
					</Link>
				);
			})}
			{Children.map(children, (child) => {
				if (!isValidElement(child)) return null;
				return child;
			})}
		</div>
	);
}
