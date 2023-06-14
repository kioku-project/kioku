import useSWR from "swr";
import { authedFetch } from "../../../util/reauth";
import DeckOverview from "../../deck/DeckOverview";

interface GroupsTabProps {
	/**
	 * groups
	 */
	groups: Group[];
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component the GroupsTab
 */
export const GroupsTab = ({ groups, className }: GroupsTabProps) => {
	return (
		<div className="">
			{groups &&
				groups
					.filter((group: Group) => !group.isDefault)
					.map((group: Group) => {
						return (
							<DeckOverview
								key={group.groupID}
								group={group}
							></DeckOverview>
						);
					})}
			<DeckOverview></DeckOverview>
		</div>
	);
};
