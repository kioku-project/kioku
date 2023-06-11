package converter

type FiberCardSideContent struct {
	Header      string `json:"header"`
	Description string `json:"description"`
}

type FiberGroupMember struct {
	UserID    string `json:"userID"`
	Name      string `json:"name"`
	GroupRole string `json:"groupRole"`
}

type FiberGroup struct {
	GroupID          string `json:"userID"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
	IsDefault        bool   `json:"isDefault"`
	GroupType        string `json:"groupType"`
}

type FiberGroupMemberRequest struct {
	AdmissionID string `json:"admissionID"`
	UserID      string `json:"userID"`
	Name        string `json:"name"`
}

type FiberCreateCardRequestBody struct {
	Sides []FiberCardSideContent `json:"sides"`
}

type FiberModifyCardRequestBody struct {
	Sides []FiberCardSideContent `json:"sides"`
}

type FiberCreateCardSideRequestBody struct {
	PlaceBeforeCardSideID string `json:"placeBeforeCardSideID"`
	FiberCardSideContent
}

type FiberModifyCardSideRequestBody struct {
	FiberCardSideContent
}

type FiberGetGroupMembersResponseBody struct {
	Members []FiberGroupMember `json:"users"`
}

type FiberGetUserGroupsResponseBody struct {
	Groups []FiberGroup `json:"groups"`
}

type FiberGetGroupMemberRequestsResponseBody struct {
	MemberRequests []FiberGroupMemberRequest `json:"memberRequests"`
}
