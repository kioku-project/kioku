import useSWR from "swr";

import { Card } from "@/types/Card";
import { Deck } from "@/types/Deck";
import { Group } from "@/types/Group";
import { Invitation } from "@/types/Invitation";
import { User } from "@/types/User";
import {
	activeDecksRoute,
	cardsRoute,
	deckRoute,
	decksRoute,
	dueCardsDeckRoute,
	dueCardsUserRoute,
	favoriteDecksRoute,
	groupMembersRoute,
	groupRoute,
	groupsRoute,
	invitationsGroupRoute,
	invitationsUserRoute,
	notificationsRoute,
	pullCardsRoute,
	requestsGroupRoute,
	userRoute,
} from "@/util/endpoints";

import { authedFetch } from "./reauth";

export const fetcher = (url: RequestInfo | URL) =>
	authedFetch(url, {
		method: "GET",
	}).then((res) => res?.json());

export function useUser() {
	const { data, error, isLoading, isValidating } = useSWR<User>(
		userRoute,
		fetcher
	);
	return {
		user: data,
		error,
		isLoading,
		isValidating,
	};
}

export function useUserDue() {
	const { data, error, isLoading, isValidating } = useSWR<{
		dueCards: number;
		dueDecks: number;
	}>(dueCardsUserRoute, fetcher);
	return {
		due: data,
		error,
		isLoading,
		isValidating,
	};
}

export function useInvitations(condition: boolean = true) {
	const { data, error, isLoading, isValidating } = useSWR<{
		groups: Invitation[];
	}>(condition ? invitationsUserRoute : null, fetcher);
	return {
		invitations: data?.groups,
		error,
		isLoading,
		isValidating,
	};
}

export function useNotifications() {
	const { data, error, isLoading, isValidating } = useSWR(
		notificationsRoute,
		fetcher
	);
	return {
		subscriptions: data?.userSubscriptions,
		error,
		isLoading,
		isValidating,
	};
}

export function useGroups() {
	const { data, error, isLoading, isValidating } = useSWR<{
		groups: Group[];
	}>(groupsRoute, fetcher);
	return {
		groups: data?.groups,
		error,
		isLoading,
		isValidating,
	};
}

export function useGroup(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<Group>(
		groupID ? groupRoute(groupID) : null,
		fetcher
	);
	return {
		group: data,
		error,
		isLoading,
		isValidating,
	};
}

export function useMembers(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		users: User[];
	}>(groupID ? groupMembersRoute(groupID) : null, fetcher);
	return {
		members: data?.users,
		error,
		isLoading,
		isValidating,
	};
}

export function useInvitedUser(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		groupInvitations: User[];
	}>(groupID ? invitationsGroupRoute(groupID) : null, fetcher);
	return {
		invitedUser: data?.groupInvitations,
		error,
		isLoading,
		isValidating,
	};
}

export function useRequestedUser(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		memberRequests: User[];
	}>(groupID ? requestsGroupRoute(groupID) : null, fetcher);
	return {
		requestedUser: data?.memberRequests,
		error,
		isLoading,
		isValidating,
	};
}

export function useDecks(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		decks: Deck[];
	}>(groupID ? decksRoute(groupID) : null, fetcher);
	return {
		decks: data?.decks,
		error,
		isLoading,
		isValidating,
	};
}

export function useFavoriteDecks() {
	const { data, error, isLoading, isValidating } = useSWR<{
		decks: Deck[];
	}>(favoriteDecksRoute, fetcher);
	return {
		decks: data?.decks,
		error,
		isLoading,
		isValidating,
	};
}

export function useActiveDecks() {
	const { data, error, isLoading, isValidating } = useSWR<{
		decks: Deck[];
	}>(activeDecksRoute, fetcher);
	return {
		decks: data?.decks,
		error,
		isLoading,
		isValidating,
	};
}

export function useDeck(deckID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<Deck>(
		deckID ? deckRoute(deckID) : null,
		fetcher
	);
	return {
		deck: data,
		error,
		isLoading,
		isValidating,
	};
}

export function useDeckDueCards(deckID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		dueCards: number;
	}>(deckID ? dueCardsDeckRoute(deckID) : null, fetcher);
	return {
		dueCards: data?.dueCards,
		error,
		isLoading,
		isValidating,
	};
}

export function useCards(deckID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		cards: Card[];
	}>(deckID ? cardsRoute(deckID) : null, fetcher);
	return {
		cards: data?.cards,
		error,
		isLoading,
		isValidating,
	};
}

export function usePullCard(deckID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<Card>(
		deckID ? pullCardsRoute(deckID) : null,
		fetcher
	);
	return {
		card: data,
		error,
		isLoading,
		isValidating,
	};
}
