import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";

import { useActiveDecks, useFavoriteDecks } from "../../../util/swr";
import DeckList from "../../deck/DeckList";

interface DashboardTabProps {
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the CardsTab
 */
export const DashboardTab = ({ className = "" }: DashboardTabProps) => {
	const { decks: activeDecks } = useActiveDecks();
	const { decks: favoriteDecks } = useFavoriteDecks();

	const { _ } = useLingui();

	return (
		<div className={`space-y-5 ${className}`}>
			<DeckList header={_(msg`Active decks`)} decks={activeDecks} />
			<DeckList header={_(msg`Favorite Decks`)} decks={favoriteDecks} />
		</div>
	);
};
