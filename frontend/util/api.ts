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
	groupMemberRoutes,
	groupRoute,
	groupsRoute,
	invitationsRoute,
	memberRoute,
	membersRoute,
	pushCardsRoute,
	requestsRoute,
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
	request: (url: string, body: string) => Promise<Response> = postRequest
) {
	const body: Record<string, string> = {};
	inputs.forEach((input) => {
		body[input.name] = input.value;
	});
	return request(url, JSON.stringify(body));
}

export async function modifyUser(inputs: HTMLInputElement[]) {
	const response = await submitForm(userRoute, inputs, putRequest);
	if (response?.ok) mutate(userRoute);
	return response;
}

export async function deleteUser() {
	const response = await deleteRequest(userRoute);
	return response;
}

export async function createGroup(inputs: HTMLInputElement[]) {
	const response = await submitForm(groupsRoute, inputs, postRequest);
	if (response?.ok) mutate(groupsRoute);
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
	const response = await putRequest(route, JSON.stringify(body));
	if (response?.ok) mutate(route);
	return response;
}

async function groupInvitation(
	groupID: string,
	userEmail: string,
	request: (url: string, body: string) => Promise<Response>
) {
	const response = await request(
		invitationsRoute(groupID),
		JSON.stringify({
			invitedUserEmail: userEmail,
		})
	);
	if (response?.ok) mutateAll(groupMemberRoutes(groupID));
	return response;
}

export async function sendGroupInvitation(groupID: string, userEmail: string) {
	return groupInvitation(groupID, userEmail, postRequest);
}

export async function declineGroupInvitation(
	groupID: string,
	userEmail: string
) {
	return groupInvitation(groupID, userEmail, deleteRequest);
}

export async function deleteMember(groupID: string, userID: string) {
	const response = await deleteRequest(memberRoute(groupID, userID));
	if (response?.ok) mutate(membersRoute(groupID));
	return response;
}

export async function joinGroup(groupID: string) {
	const response = await postRequest(requestsRoute(groupID));
	if (response?.ok) mutateAll([groupsRoute, ...groupMemberRoutes(groupID)]);
	return response;
}

export async function leaveGroup(groupID: string) {
	const response = await deleteRequest(membersRoute(groupID));
	if (response?.ok) mutate(groupsRoute);
	return response;
}

export async function deleteGroup(groupID: string) {
	const response = await deleteRequest(groupRoute(groupID));
	if (response?.ok) mutate(groupsRoute);
	return response;
}

export async function createDeck(inputs: HTMLInputElement[], groupID: string) {
	const route = decksRoute(groupID);
	const response = await submitForm(route, inputs, postRequest);
	if (response?.ok) mutate(route);
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
	const response = await putRequest(route, JSON.stringify(body));
	if (response?.ok) mutate(route);
	return response;
}

export async function deleteDeck(deckID: string, groupID: string) {
	const response = await deleteRequest(deckRoute(deckID));
	if (response?.ok) mutate(decksRoute(groupID));
	return response;
}

export async function toggleFavorite(
	deckID: string,
	groupID: string,
	isFavorite: boolean | undefined
) {
	const response = await apiRequest(
		isFavorite ? "DELETE" : "POST",
		favoriteDecksRoute,
		JSON.stringify({
			deckID: deckID,
		})
	);
	if (response?.ok)
		mutateAll([decksRoute(groupID), favoriteDecksRoute, activeDecksRoute]);
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
		})
	);
	if (response?.ok) mutateAll([cardsRoute(deckID), ...cardRoutes(deckID)]);
	return response;
}

export async function modifyCard(card: CardType) {
	const response = await putRequest(
		cardRoute(card.cardID),
		JSON.stringify({
			sides: card.sides,
		})
	);
	if (response?.ok && card.deckID) mutate(cardsRoute(card.deckID));
	return response;
}

export async function pushCard(deckID: string, cardID: string, rating: number) {
	const response = await postRequest(
		pushCardsRoute(deckID),
		JSON.stringify({ body: { cardID, rating } })
	);
	if (response?.ok) mutateAll(cardRoutes(deckID));
	return response;
}

export async function deleteCard(deckID: string, cardID: string) {
	const response = await deleteRequest(cardRoute(cardID));
	if (response?.ok) mutateAll([cardsRoute(deckID), ...cardRoutes(deckID)]);
	return response;
}

function mutateAll(routes: string[]) {
	routes.forEach((route) => mutate(route));
}
