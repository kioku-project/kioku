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
