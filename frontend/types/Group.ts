import { GroupRole } from "./GroupRole";

export type Group = {
	groupID: string;
	groupName: string;
	groupDescription?: string;
	groupType?: "CLOSED" | "REQUEST" | "OPEN";
	groupRole?: keyof typeof GroupRole;
	isDefault?: boolean;
	isEmpty?: boolean;
};
