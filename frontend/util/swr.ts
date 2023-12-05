import useSWR from "swr";

import { Card } from "@/types/Card";
import { Deck } from "@/types/Deck";
import { Group } from "@/types/Group";
import { Invitation } from "@/types/Invitation";
import { User } from "@/types/User";

import { authedFetch } from "./reauth";

export const fetcher = (url: RequestInfo | URL) =>
	authedFetch(url, {
		method: "GET",
	}).then((res) => res?.json());

export function useUser() {
	const { data, error, isLoading, isValidating } = useSWR<User>(`/api/user`, fetcher);
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
	}>(`/api/user/dueCards`, fetcher);
	return {
		due: data,
		error,
		isLoading,
		isValidating,
	};
}

export function useInvitations() {
	const { data, error, isLoading, isValidating } = useSWR<{
		groupInvitation: Invitation[];
	}>(`/api/user/invitations`, fetcher);
	return {
		invitations: data?.groupInvitation,
		error,
		isLoading,
		isValidating,
	};
}

export function useGroups() {
	const { data, error, isLoading, isValidating } = useSWR<{
		groups: Group[];
	}>(`/api/groups`, fetcher);
	return {
		groups: data?.groups,
		error,
		isLoading,
		isValidating,
	};
}

export function useGroup(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<Group>(
		groupID ? `/api/groups/${groupID}` : null,
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
	}>(groupID ? `/api/groups/${groupID}/members` : null, fetcher);
	return {
		members: data?.users,
		error,
		isLoading,
		isValidating,
	};
}

export function useRequestedUser(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		memberRequests: User[];
	}>(groupID ? `/api/groups/${groupID}/members/requests` : null, fetcher);
	return {
		requestedUser: data?.memberRequests,
		error,
		isLoading,
		isValidating,
	};
}

export function useInvitedUser(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		groupInvitations: User[];
	}>(groupID ? `/api/groups/${groupID}/members/invitations` : null, fetcher);
	return {
		invitedUser: data?.groupInvitations,
		error,
		isLoading,
		isValidating,
	};
}

export function useDecks(groupID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{
		decks: Deck[];
	}>(groupID ? `/api/groups/${groupID}/decks` : null, fetcher);
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
	}>(`/api/decks/favorites`, fetcher);
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
	}>(`/api/decks/active`, fetcher);
	return {
		decks: data?.decks,
		error,
		isLoading,
		isValidating,
	};
}

export function useDeck(deckID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<Deck>(
		deckID ? `/api/decks/${deckID}` : null,
		fetcher
	);
	return {
		deck: data,
		error,
		isLoading,
		isValidating,
	};
}

export function useDueCards(deckID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<{ dueCards: number }>(
		deckID ? `/api/decks/${deckID}/dueCards` : null,
		fetcher
	);
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
	}>(deckID ? `/api/decks/${deckID}/cards` : null, fetcher);
	return {
		cards: data?.cards,
		error,
		isLoading,
		isValidating,
	};
}

export function usePullCard(deckID?: string) {
	const { data, error, isLoading, isValidating } = useSWR<Card>(
		deckID ? `/api/decks/${deckID}/pull` : null,
		fetcher
	);
	return {
		card: data,
		error,
		isLoading,
		isValidating,
	};
}
