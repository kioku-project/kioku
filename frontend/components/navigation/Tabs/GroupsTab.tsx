import { Group as GroupType } from "../../../types/Group";
import DeckOverview from "../../deck/DeckOverview";

interface GroupsTabProps {
	/**
	 * groups
	 */
	groups: GroupType[];
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the GroupsTab
 */
export const GroupsTab = ({ groups, className = "" }: GroupsTabProps) => {
	return (
		<div className={`${className}`}>
			{groups
				?.filter((group: GroupType) => !group.isDefault)
				.map((group: GroupType) => {
					return (
						<DeckOverview
							key={group.groupID}
							group={group}
						></DeckOverview>
					);
				})}
			<DeckOverview />
		</div>
	);
};
