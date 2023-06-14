import useSWR from "swr";
import { authedFetch } from "../../util/reauth";
import { Card } from "./Card";

interface CardListProps {
	/**
	 * deckID
	 */
	deckID: string;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * click handler
	 */
	setCard: (card: Card) => void;
}

/**
 * UI component for displaying a list of cards
 */
export const CardList = ({ deckID, setCard, className }: CardListProps) => {
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: cards } = useSWR(`/api/decks/${deckID}/cards`, fetcher);
	return (
		<div id="cardListId" className={`flex flex-col ${className ?? ""}`}>
			<div className="snap-y overflow-y-auto">
				{cards?.cards &&
					cards.cards.map((card: Card) => (
						<Card
							className="snap-center"
							key={card.cardID}
							setCard={setCard}
							card={{ ...card, deckID: deckID }}
						/>
					))}
			</div>
			<Card card={{ cardID: "", sides: [], deckID: deckID }}></Card>
		</div>
	);
};
