import { useRouter } from "next/router";
import React from "react";
import "react-calendar-heatmap/dist/styles.css";
import "react-tooltip/dist/react-tooltip.css";
import Authenticated from "../components/accessControl/Authenticated";
import DeckOverview from "../components/deck/DeckOverview";
import { Header } from "../components/navigation/Header";

export default function Home() {
	const router = useRouter();
	return (
		<Authenticated>
			<div className="min-w-screen flex h-screen select-none flex-col bg-eggshell">
				<Header login={true}></Header>
				<div className="space-y-10 p-10">
					<DeckOverview
						name="Your personal Decks"
						id="personalDecks"
						decks={[
							{ name: "Japan", count: 0 },
							{ name: "BWL Grundlagen", count: 1 },
							{ name: "Marketing", count: 2 },
							{ name: "Mathe", count: 3 },
							{ name: "English", count: 4 },
							{ name: "Deck 1", count: 5 },
							{ name: "Deck 2", count: 6 },
							{ name: "Deck 3", count: 7 },
							{ name: "Deck 4", count: 8 },
							{ name: "Deck 5", count: 9 },
							{ name: "Deck 6", count: 10 },
							{ name: "Deck 7", count: 11 },
						]}
					></DeckOverview>
					<DeckOverview
						name="Your Groups"
						id="groupDecks"
						decks={[
							{ name: "Japan Class", count: 0 },
							{ name: "DHBW", count: 1 },
							{ name: "BWL", count: 2 },
						]}
					></DeckOverview>
				</div>
			</div>
		</Authenticated>
	);
}
