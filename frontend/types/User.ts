type User = {
	userID: string;
	userName: string;
	groupRole?: "ADMIN" | "WRITE" | "READ";
	status?: "requested" | "invited";
	admissionID?: string;
};
