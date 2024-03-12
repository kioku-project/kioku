import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { ChangeEvent, MouseEvent } from "react";

import DeckList from "@/components/deck/DeckList";
import { ActionBar } from "@/components/input/ActionBar";
import { SpeechBubble } from "@/components/input/SpeechBubble";
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
						title="Start learning!"
						description="If you start learning, your active decks will appear here."
						iconName="Activity"
					></GenericPlaceholder>
				)}
			</DeckList>
			<DeckList header={_(msg`Favorite Decks`)} decks={favoriteDecks}>
				{favoriteDecks?.length === 0 && (
					<GenericPlaceholder
						title="Find your favorites!"
						description="Click on the heart icon in the top right corner of a deck to add it to this list."
						buttonText="Show me how"
						iconName="Heart"
					></GenericPlaceholder>
				)}
			</DeckList>
		</div>
	);
};
