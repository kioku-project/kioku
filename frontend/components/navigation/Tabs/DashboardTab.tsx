import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";

import DeckList from "@/components/deck/DeckList";
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
			<DeckList header={_(msg`Active Decks`)} decks={activeDecks}>
				{activeDecks?.length === 0 && (
					<GenericPlaceholder
						title={_(msg`Start learning!`)}
						description={_(
							msg`If you start learning, your active decks will appear here.`
						)}
						iconName="Activity"
					></GenericPlaceholder>
				)}
			</DeckList>
			<DeckList header={_(msg`Favorite Decks`)} decks={favoriteDecks}>
				{favoriteDecks?.length === 0 && (
					<GenericPlaceholder
						title={_(msg`Find your favorites!`)}
						description={_(
							msg`Click on the heart icon in the top right corner of a deck to add it to this list.`
						)}
						iconName="Heart"
					></GenericPlaceholder>
				)}
			</DeckList>
		</div>
	);
};
