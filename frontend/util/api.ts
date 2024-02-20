import { t } from "@lingui/macro";
import { toast } from "react-hot-toast";
import { mutate } from "swr";

import { Card as CardType } from "@/types/Card";
import {
	activeDecksRoute,
	cardRoute,
	cardRoutes,
	cardsRoute,
	deckRoute,
	decksRoute,
	favoriteDecksRoute,
	groupMemberRoute,
	groupMemberRoutes,
	groupMembersRoute,
	groupRoute,
	groupsRoute,
	invitationGroupRoute,
	invitationsUserRoute,
	pushCardsRoute,
	requestGroupRoute,
	userRoute,
} from "@/util/endpoints";

import { authedFetch } from "./reauth";

export async function apiRequest(method: string, url: string, body?: string) {
	const response = await authedFetch(url, {
		method,
		headers: {
			"Content-Type": "application/json",
		},
		body,
	});
	return response;
}

export async function postRequest(url: string, body?: string) {
	return apiRequest("POST", url, body);
}
export async function putRequest(url: string, body?: string) {
	return apiRequest("PUT", url, body);
}

export async function deleteRequest(url: string, body?: string) {
	return apiRequest("DELETE", url, body);
}

export async function submitForm(
	url: string,
	inputs: HTMLInputElement[],
	request: (url: string, body: string) => Promise<Response> = postRequest,
) {
	const body: Record<string, string> = {};
	inputs.forEach((input) => {
		body[input.name] = input.value;
	});
	return request(url, JSON.stringify(body));
}

export async function modifyUser(inputs: HTMLInputElement[]) {
	const response = await submitForm(userRoute, inputs, putRequest);
	if (response?.ok) {
		mutate(userRoute);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function deleteUser() {
	const toastID = toast.loading(t`Deleting user...`);
	const response = await deleteRequest(userRoute);
	if (response.ok) {
		toast.success(t`User deleted`, { id: toastID });
	} else {
		const error = await response.text();
		toast.error(error, { id: toastID });
	}
	return response;
}

export async function createGroup(inputs: HTMLInputElement[]) {
	const response = await submitForm(groupsRoute, inputs, postRequest);
	if (response.ok) {
		mutate(groupsRoute);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function modifyGroup(
	groupID: string,
	body: {
		groupName?: string;
		groupDescription?: string;
		groupType?: string;
	},
) {
	const route = groupRoute(groupID);
	const response = await putRequest(route, JSON.stringify(body));
	if (response.ok) {
		mutate(route);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

// handle group invitations
async function groupInvitation(
	groupID: string,
	userEmail: string,
	request: (url: string, body?: string) => Promise<Response>,
) {
	const response = await request(
		invitationGroupRoute(groupID),
		JSON.stringify({
			invitedUserEmail: userEmail,
		}),
	);
	if (response.ok) {
		mutateAll(groupMemberRoutes(groupID));
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

// Invite user to group or accept group request
export async function sendGroupInvitation(groupID: string, userEmail: string) {
	return groupInvitation(groupID, userEmail, postRequest);
}

// Decline group request or revoke group invitation
export async function declineGroupRequest(groupID: string, userEmail: string) {
	return groupInvitation(groupID, userEmail, deleteRequest);
}

// handle group requests
async function groupRequest(
	groupID: string,
	request: (url: string, body?: string) => Promise<Response>,
) {
	const response = await request(requestGroupRoute(groupID));
	if (response.ok) {
		mutateAll([
			invitationsUserRoute,
			groupsRoute,
			groupRoute(groupID),
			...groupMemberRoutes(groupID),
		]);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

// send group request or accept group invitation
export async function sendGroupRequest(groupID: string) {
	return groupRequest(groupID, postRequest);
}

// decline group invitation or revoke group request
export async function declineGroupInvitation(groupID: string) {
	return groupRequest(groupID, deleteRequest);
}

export async function deleteMember(groupID: string, userID: string) {
	const response = await deleteRequest(groupMemberRoute(groupID, userID));
	if (response.ok) {
		mutate(groupMembersRoute(groupID));
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function leaveGroup(groupID: string) {
	const toastID = toast.loading(t`Leaving group...`);
	const response = await deleteRequest(groupMembersRoute(groupID));
	if (response.ok) {
		toast.success(t`Left group`, { id: toastID });
		mutate(groupsRoute);
	} else {
		const error = await response.text();
		toast.error(error, { id: toastID });
	}
	return response;
}

export async function deleteGroup(groupID: string) {
	const toastID = toast.loading(t`Deleting group...`);
	const response = await deleteRequest(groupRoute(groupID));
	if (response.ok) {
		toast.success(t`Group deleted`, { id: toastID });
		mutate(groupsRoute);
	} else {
		const error = await response.text();
		toast.error(error, { id: toastID });
	}
	return response;
}

export async function createDeck(inputs: HTMLInputElement[], groupID: string) {
	const route = decksRoute(groupID);
	const response = await submitForm(route, inputs, postRequest);
	if (response.ok) {
		mutate(route);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function modifyDeck(
	deckID: string,
	body: {
		deckName?: string;
		deckDescription?: string;
		deckType?: "PUBLIC" | "PRIVATE";
	},
) {
	const route = deckRoute(deckID);
	const response = await putRequest(route, JSON.stringify(body));
	if (response.ok) {
		mutate(route);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function deleteDeck(deckID: string, groupID: string) {
	const toastID = toast.loading(t`Deleting deck...`);
	const response = await deleteRequest(deckRoute(deckID));
	if (response.ok) {
		toast.success(t`Deck deleted`, { id: toastID });
		mutate(decksRoute(groupID));
	} else {
		const error = await response.text();
		toast.error(error, { id: toastID });
	}
	return response;
}

export async function toggleFavorite(
	deckID: string,
	groupID: string,
	isFavorite: boolean | undefined,
) {
	const response = await apiRequest(
		isFavorite ? "DELETE" : "POST",
		favoriteDecksRoute,
		JSON.stringify({
			deckID: deckID,
		}),
	);
	if (response.ok) {
		mutateAll([decksRoute(groupID), favoriteDecksRoute, activeDecksRoute]);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function createCard(deckID: string, input: HTMLInputElement) {
	if (!input.value) {
		input.focus();
		return;
	}
	const response = await postRequest(
		cardsRoute(deckID),
		JSON.stringify({
			sides: [
				{
					header: input.value,
				},
			],
		}),
	);
	if (response.ok) {
		mutateAll([cardsRoute(deckID), ...cardRoutes(deckID)]);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function modifyCard(card: CardType) {
	const response = await putRequest(
		cardRoute(card.cardID),
		JSON.stringify({
			sides: card.sides,
		}),
	);
	if (response.ok && card.deckID) {
		mutate(cardsRoute(card.deckID));
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function pushCard(deckID: string, cardID: string, rating: number) {
	const response = await postRequest(
		pushCardsRoute(deckID),
		JSON.stringify({ body: { cardID, rating } }),
	);
	if (response.ok) {
		mutateAll(cardRoutes(deckID));
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

export async function deleteCard(deckID: string, cardID: string) {
	const response = await deleteRequest(cardRoute(cardID));
	if (response.ok) {
		mutateAll([cardsRoute(deckID), ...cardRoutes(deckID)]);
	} else {
		const error = await response.text();
		toast.error(error);
	}
	return response;
}

function mutateAll(routes: string[]) {
	routes.forEach((route) => mutate(route));
}
