import { t } from "@lingui/macro";
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
import { handleWithToast } from "@/util/toasts";

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
	request: (url: string, body: string) => Promise<Response> = postRequest
) {
	const body: Record<string, string> = {};
	inputs.forEach((input) => {
		body[input.name] = input.value;
	});
	return request(url, JSON.stringify(body));
}

export async function modifyUser(inputs: HTMLInputElement[]) {
	const response = await handleWithToast(
		submitForm(userRoute, inputs, putRequest),
		"modifyUserToastID"
	);
	if (response.ok) mutate(userRoute);
	return response;
}

export async function deleteUser() {
	return await handleWithToast(
		deleteRequest(userRoute),
		"deleteUserToastID",
		t`User deleted`,
		t`Deleting user...`
	);
}

export async function createGroup(inputs: HTMLInputElement[]) {
	const response = await handleWithToast(
		submitForm(groupsRoute, inputs, postRequest),
		"createGroupToastID"
	);
	if (response.ok) mutate(groupsRoute);
	return response;
}

export async function modifyGroup(
	groupID: string,
	body: {
		groupName?: string;
		groupDescription?: string;
		groupType?: string;
	}
) {
	const route = groupRoute(groupID);
	const response = await handleWithToast(
		putRequest(route, JSON.stringify(body)),
		"modifyGroupToastID"
	);
	if (response.ok) mutate(route);
	return response;
}

// handle group invitations
async function groupInvitation(
	groupID: string,
	userEmail: string,
	request: (url: string, body?: string) => Promise<Response>
) {
	const response = await handleWithToast(
		request(
			invitationGroupRoute(groupID),
			JSON.stringify({
				invitedUserEmail: userEmail,
			})
		),
		"groupInvitationToastID"
	);
	if (response.ok) mutateAll(groupMemberRoutes(groupID));
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
	request: (url: string, body?: string) => Promise<Response>
) {
	const response = await handleWithToast(
		request(requestGroupRoute(groupID)),
		"groupRequestToastID"
	);
	if (response.ok)
		mutateAll([
			invitationsUserRoute,
			groupsRoute,
			groupRoute(groupID),
			...groupMemberRoutes(groupID),
		]);
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
	const response = await handleWithToast(
		deleteRequest(groupMemberRoute(groupID, userID)),
		"deleteMemberToastID"
	);
	if (response.ok) mutate(groupMembersRoute(groupID));
	return response;
}

export async function leaveGroup(groupID: string) {
	const response = await handleWithToast(
		deleteRequest(groupMembersRoute(groupID)),
		"leaveGroupToastID",
		t`Left group`,
		t`Leaving group...`
	);
	if (response.ok) mutate(groupsRoute);
	return response;
}

export async function deleteGroup(groupID: string) {
	const response = await handleWithToast(
		deleteRequest(groupRoute(groupID)),
		"deleteGroupToastID",
		t`Group deleted`,
		t`Deleting group...`
	);
	if (response.ok) mutate(groupsRoute);
	return response;
}

export async function createDeck(inputs: HTMLInputElement[], groupID: string) {
	const route = decksRoute(groupID);
	const response = await handleWithToast(
		submitForm(route, inputs, postRequest),
		"createDeckToastID"
	);
	if (response.ok) mutate(route);
	return response;
}

export async function modifyDeck(
	deckID: string,
	body: {
		deckName?: string;
		deckDescription?: string;
		deckType?: "PUBLIC" | "PRIVATE";
	}
) {
	const route = deckRoute(deckID);
	const response = await handleWithToast(
		putRequest(route, JSON.stringify(body)),
		"modifyDeckToastID"
	);
	if (response.ok) mutate(route);
	return response;
}

export async function deleteDeck(deckID: string, groupID: string) {
	const response = await handleWithToast(
		deleteRequest(deckRoute(deckID)),
		"deleteDeckToastID",
		t`Deck deleted`,
		t`Deleting deck...`
	);
	if (response.ok) mutate(decksRoute(groupID));
	return response;
}

export async function toggleFavorite(
	deckID: string,
	groupID: string,
	isFavorite: boolean | undefined
) {
	const response = await handleWithToast(
		apiRequest(
			isFavorite ? "DELETE" : "POST",
			favoriteDecksRoute,
			JSON.stringify({
				deckID: deckID,
			})
		),
		"toggleFavoriteToastID"
	);
	if (response.ok)
		mutateAll([decksRoute(groupID), favoriteDecksRoute, activeDecksRoute]);
	return response;
}

export async function createCard(
	deckID: string,
	front: HTMLInputElement,
	back: HTMLInputElement
) {
	if (!front.value) {
		front.focus();
		return;
	}
	const response = await handleWithToast(
		postRequest(
			cardsRoute(deckID),
			JSON.stringify({
				sides: [
					{
						header: front.value,
					},
					{
						header: back.value,
					},
				],
			})
		),
		"createCardToastID"
	);
	if (response.ok) mutateAll([cardsRoute(deckID), ...cardRoutes(deckID)]);
	return response;
}

export async function modifyCard(card: CardType) {
	const response = await handleWithToast(
		putRequest(
			cardRoute(card.cardID),
			JSON.stringify({
				sides: card.sides,
			})
		),
		"modifyCardToastID"
	);
	if (response.ok && card.deckID) mutate(cardsRoute(card.deckID));
	return response;
}

export async function pushCard(deckID: string, cardID: string, rating: number) {
	const response = await handleWithToast(
		postRequest(pushCardsRoute(deckID), JSON.stringify({ cardID, rating })),
		"pushCardToastID"
	);
	if (response.ok) mutateAll(cardRoutes(deckID));
	return response;
}

export async function deleteCard(deckID: string, cardID: string) {
	const response = await handleWithToast(
		deleteRequest(cardRoute(cardID)),
		"deleteCardToastID"
	);
	if (response.ok) mutateAll([cardsRoute(deckID), ...cardRoutes(deckID)]);
	return response;
}

function mutateAll(routes: string[]) {
	routes.forEach((route) => mutate(route));
}
