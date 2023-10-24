import { Group as GroupType } from "../../../types/Group";
import DeckOverview from "../../deck/DeckOverview";

interface DecksTabProps {
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
 * UI component for the DecksTab
 */
export const DecksTab = ({ group, className = "" }: DecksTabProps) => {
	return (
		<div className={`${className}`}>
			<DeckOverview group={{ ...group, groupName: "" }}></DeckOverview>
		</div>
	);
};
