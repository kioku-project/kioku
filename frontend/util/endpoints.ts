// User routes
export const userRoute = "/api/user";
export const registerRoute = "/api/register";
export const loginRoute = "/api/login";
export const logoutRoute = "/api/logout";
export const reauthRoute = "/api/reauth";
export const dueCardsUserRoute = "/api/user/dueCards";
export const invitationsUserRoute = "/api/user/invitations";
export const notificationsRoute = "/api/user/notification";
export const notificationRoute = (notificationID: string) =>
	`/api/user/notification/${notificationID}`;

// Group routes
export const groupsRoute = "/api/groups";
export const groupRoute = (groupID: string) => `/api/groups/${groupID}`;
export const decksRoute = (groupID: string) => `/api/groups/${groupID}/decks`;
export const groupMembersRoute = (groupID: string) =>
	`/api/groups/${groupID}/members`;
export const groupMemberRoute = (groupID: string, userID: string) =>
	`/api/groups/${groupID}/members/${userID}`;
export const invitationsGroupRoute = (groupID: string) =>
	`/api/groups/${groupID}/members/invitation`;
export const requestsGroupRoute = (groupID: string) =>
	`/api/groups/${groupID}/members/request`;
export const groupMemberRoutes = (groupID: string) => [
	groupMembersRoute(groupID),
	invitationsGroupRoute(groupID),
	`/api/groups/${groupID}/members/invitations`,
	requestsGroupRoute(groupID),
	`/api/groups/${groupID}/members/requests`,
];

// Deck routes
export const activeDecksRoute = "/api/decks/active";
export const favoriteDecksRoute = "/api/decks/favorites";
export const deckRoute = (deckID: string) => `/api/decks/${deckID}`;
export const cardsRoute = (deckID: string) => `/api/decks/${deckID}/cards`;
export const pullCardsRoute = (deckID: string) => `/api/decks/${deckID}/pull`;
export const pushCardsRoute = (deckID: string) => `/api/decks/${deckID}/push`;
export const dueCardsDeckRoute = (deckID: string) =>
	`/api/decks/${deckID}/dueCards`;

// Card routes
export const cardRoute = (cardID: string) => `/api/cards/${cardID}`;
export const cardRoutes = (deckID: string) => [
	pullCardsRoute(deckID),
	dueCardsDeckRoute(deckID),
];
