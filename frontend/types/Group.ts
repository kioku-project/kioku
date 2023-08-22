import { groupRole } from "./GroupRole";

export type Group = {
	groupID: string;
	groupName: string;
	groupDescription?: string;
	isDefault?: boolean;
	groupType?: string;
	groupRole?: keyof typeof groupRole;
	isEmpty?: boolean;
};
