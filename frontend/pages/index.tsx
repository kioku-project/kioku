import React from "react";
import Authenticated from "../components/accessControl/Authenticated";
import DeckOverview from "../components/deck/DeckOverview";
import { Navbar } from "../components/navigation/Navbar";
import Head from "next/head";

export default function Home() {
	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<Authenticated>
				<div className="min-w-screen flex h-screen select-none flex-col bg-eggshell">
					<Navbar login={true}></Navbar>
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
		</div>
	);
}
