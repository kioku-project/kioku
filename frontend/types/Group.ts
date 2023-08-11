export type Group = {
	groupID: string;
	groupName: string;
	groupDescription?: string;
	isDefault?: boolean;
	groupType?: string;
	groupRole?: "ADMIN" | "WRITE" | "READ" | "INVITED" | "REQUESTED";
	isEmpty?: boolean;
};

export enum groupRole {
	EXTERNAL,
	REQUESTED,
	INVITED,
	READ,
	WRITE,
	ADMIN,
}
