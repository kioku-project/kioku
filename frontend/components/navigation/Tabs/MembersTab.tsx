import { Group as GroupType } from "../../../types/Group";
import MemberList from "../../group/MemberList";

interface MembersTabProps {
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
 * UI component for the MembersTab
 */
export const MembersTab = ({ group, className = "" }: MembersTabProps) => {
	return (
		<div className={`${className}`}>
			<MemberList group={group}></MemberList>
		</div>
	);
};
