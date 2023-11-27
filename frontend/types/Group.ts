import { GroupRole } from "./GroupRole";

export type Group = {
	groupID: string;
	groupName: string;
	groupDescription?: string;
	isDefault: boolean;
	groupType?: "CLOSED" | "REQUEST" | "OPEN";
	groupRole: keyof typeof GroupRole;
	isEmpty?: boolean;
};
