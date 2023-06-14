import { Group } from "../../../types/Group";
import MemberList from "../../group/MemberList";

interface MembersTabProps {
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
 * UI component for the MembersTab
 */
export const MembersTab = ({ group, className }: MembersTabProps) => {
	return (
		<div>
			<MemberList group={group}></MemberList>
		</div>
	);
};
