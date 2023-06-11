package converter

type FiberCardSideContent struct {
	Header      string `json:"header"`
	Description string `json:"description"`
}

type FiberGroupMember struct {
	UserID    string `json:"userID"`
	Name      string `json:"userName"`
	GroupRole string `json:"groupRole"`
}

type FiberGroup struct {
	GroupID          string `json:"groupID"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
	IsDefault        bool   `json:"isDefault"`
	GroupType        string `json:"groupType"`
	GroupRole        string `json:"groupRole"`
}

type FiberGroupMemberAdmission struct {
	AdmissionID string `json:"admissionID"`
	UserID      string `json:"userID"`
	Name        string `json:"userName"`
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

type FiberPushCardRequestBody struct {
	CardID string `json:"cardID"`
	Rating int64  `json:"rating"`
}

type FiberModifyCardSideRequestBody struct {
	FiberCardSideContent
}

type FiberGetGroupMembersResponseBody struct {
	Members []FiberGroupMember `json:"users"`
}

type FiberGetUserDueCardsResponseBody struct {
	DueCards int64 `json:"dueCards"`
	DueDecks int64 `json:"dueDecks"`
}

type FiberGetUserGroupsResponseBody struct {
	Groups []FiberGroup `json:"groups"`
}

type FiberGetGroupMemberRequestsResponseBody struct {
	MemberRequests []FiberGroupMemberAdmission `json:"memberRequests"`
}

type FiberGetInvitationsForGroupResponseBody struct {
	MemberRequests []FiberGroupMemberAdmission `json:"groupInvitations"`
}
