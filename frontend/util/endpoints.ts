// User routes
export const userRoute = "/api/user";
export const notificationsRoute = "/api/user/notification";
export const notificationRoute = (notificationID: string) =>
	`/api/user/notification/${notificationID}`;
export const registerRoute = "/api/register";
export const loginRoute = "/api/login";
export const reauthRoute = "/api/reauth";
export const logoutRoute = "/api/logout";

// Group routes
export const groupsRoute = "/api/groups";
export const groupRoute = (groupID: string) => `/api/groups/${groupID}`;
export const decksRoute = (groupID: string) => `/api/groups/${groupID}/decks`;
export const membersRoute = (groupID: string) =>
	`/api/groups/${groupID}/members`;
export const memberRoute = (groupID: string, userID: string) =>
	`/api/groups/${groupID}/members/${userID}`;
export const invitationsRoute = (groupID: string) =>
	`/api/groups/${groupID}/members/invitation`;
export const requestsRoute = (groupID: string) =>
	`/api/groups/${groupID}/members/request`;
export const groupMemberRoutes = (groupID: string) => [
	membersRoute(groupID),
	invitationsRoute(groupID),
	`/api/groups/${groupID}/members/invitations`,
	requestsRoute(groupID),
	`/api/groups/${groupID}/members/requests`,
];

// Deck routes
export const activeDecksRoute = "/api/decks/active";
export const favoriteDecksRoute = "/api/decks/favorites";
export const deckRoute = (deckID: string) => `/api/decks/${deckID}`;
export const cardsRoute = (deckID: string) => `/api/decks/${deckID}/cards`;
export const pullCardsRoute = (deckID: string) => `/api/decks/${deckID}/pull`;
export const pushCardsRoute = (deckID: string) => `/api/decks/${deckID}/push`;
export const dueCardsRoute = (deckID: string) =>
	`/api/decks/${deckID}/dueCards`;

// Card routes
export const cardRoute = (cardID: string) => `/api/cards/${cardID}`;
export const cardRoutes = (deckID: string) => [
	pullCardsRoute(deckID),
	dueCardsRoute(deckID),
];
