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
	const { data, error, isLoading } = useSWR<User>(`/api/user`, fetcher);
	return {
		user: data,
		error,
		isLoading,
	};
}

export function useUserDue() {
	const { data, error, isLoading } = useSWR<{
		dueCards: number;
		dueDecks: number;
	}>(`/api/user/dueCards`, fetcher);
	return {
		due: data,
		error,
		isLoading,
	};
}

export function useInvitations() {
	const { data, error, isLoading } = useSWR<{
		groupInvitation: Invitation[];
	}>(`/api/user/invitations`, fetcher);
	return {
		invitations: data?.groupInvitation,
		error,
		isLoading,
	};
}

export function useGroups() {
	const { data, error, isLoading } = useSWR<{
		groups: Group[];
	}>(`/api/groups`, fetcher);
	return {
		groups: data?.groups,
		error,
		isLoading,
	};
}

export function useGroup(groupID?: string) {
	const { data, error, isLoading } = useSWR<Group>(
		groupID ? `/api/groups/${groupID}` : null,
		fetcher
	);
	return {
		group: data,
		error,
		isLoading,
	};
}

export function useMembers(groupID?: string) {
	const { data, error, isLoading } = useSWR<{
		users: User[];
	}>(groupID ? `/api/groups/${groupID}/members` : null, fetcher);
	return {
		members: data?.users,
		error,
		isLoading,
	};
}

export function useRequestedUser(groupID?: string) {
	const { data, error, isLoading } = useSWR<{
		memberRequests: User[];
	}>(groupID ? `/api/groups/${groupID}/members/requests` : null, fetcher);
	return {
		requestedUser: data?.memberRequests,
		error,
		isLoading,
	};
}

export function useInvitedUser(groupID?: string) {
	const { data, error, isLoading } = useSWR<{
		groupInvitations: User[];
	}>(groupID ? `/api/groups/${groupID}/members/invitations` : null, fetcher);
	return {
		invitedUser: data?.groupInvitations,
		error,
		isLoading,
	};
}

export function useDecks(groupID?: string) {
	const { data, error, isLoading } = useSWR<{
		decks: Deck[];
	}>(groupID ? `/api/groups/${groupID}/decks` : null, fetcher);
	return {
		decks: data?.decks,
		error,
		isLoading,
	};
}

export function useFavoriteDecks() {
	const { data, error, isLoading } = useSWR<{
		decks: Deck[];
	}>(`/api/decks/favorites`, fetcher);
	return {
		decks: data?.decks,
		error,
		isLoading,
	};
}

export function useActiveDecks() {
	const { data, error, isLoading } = useSWR<{
		decks: Deck[];
	}>(`/api/decks/active`, fetcher);
	return {
		decks: data?.decks,
		error,
		isLoading,
	};
}

export function useDeck(deckID?: string) {
	const { data, error, isLoading } = useSWR<Deck>(
		deckID ? `/api/decks/${deckID}` : null,
		fetcher
	);
	return {
		deck: data,
		error,
		isLoading,
	};
}

export function useDueCards(deckID?: string) {
	const { data, error, isLoading } = useSWR<{ dueCards: number }>(
		deckID ? `/api/decks/${deckID}/dueCards` : null,
		fetcher
	);
	return {
		dueCards: data?.dueCards,
		error,
		isLoading,
	};
}

export function useCards(deckID?: string) {
	const { data, error, isLoading } = useSWR<{
		cards: Card[];
	}>(deckID ? `/api/decks/${deckID}/cards` : null, fetcher);
	return {
		cards: data?.cards,
		error,
		isLoading,
	};
}

export function usePullCard(deckID?: string) {
	const { data, error, isLoading } = useSWR<Card>(
		deckID ? `/api/decks/${deckID}/pull` : null,
		fetcher
	);
	return {
		card: data,
		error,
		isLoading,
	};
}
