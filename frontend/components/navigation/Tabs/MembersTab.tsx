import MemberList from "../../group/MemberList";

interface MembersTabProps {
	/**
	 * groupID
	 */
	groupID: string;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the MembersTab
 */
export const MembersTab = ({ groupID, className }: MembersTabProps) => {
	return (
		<div>
			<MemberList groupID={groupID}></MemberList>
		</div>
	);
};
