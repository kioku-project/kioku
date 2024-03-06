import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";

import DeckList from "@/components/deck/DeckList";
import { Placeholder } from "@/components/flashcard/Card.stories";
import { GenericPlaceholder } from "@/components/placeholders/GenericPlaceholder";
import { useActiveDecks, useFavoriteDecks } from "@/util/swr";

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
			<DeckList
				header={_(msg`Active Decks`)}
				decks={activeDecks}
				title="No decks yet"
				description="Your active decks will appear here, once you start learning"
				iconName="Activity"
			/>

			<DeckList
				header={_(msg`Favorite Decks`)}
				decks={favoriteDecks}
				title="Choose your favorite deck!"
				description="Click on the heart icon in the top right corner of a deck to add it to your favorites."
				iconName="Heart"
				buttonText="Find Favorites"
			/>
		</div>
	);
};
