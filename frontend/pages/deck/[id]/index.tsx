import Head from "next/head";
import "react-toastify/dist/ReactToastify.css";
import Authenticated from "../../../components/accessControl/Authenticated";
import { Navbar } from "../../../components/navigation/Navbar";
import { useRouter } from "next/router";
import React, { ReactNode, useState } from "react";
import useSWR from "swr";
import { authedFetch } from "../../../util/reauth";
import { StatisticsTab } from "../../../components/navigation/Tabs/StatisticsTab";
import { TabHeader } from "../../../components/navigation/Tabs/TabHeader";
import { Header } from "../../../components/layout/Header";
import { CardsTab } from "../../../components/navigation/Tabs/CardsTab";
import { TabBar } from "../../../components/navigation/Tabs/TabBar";
import { DeckSettingsTab } from "../../../components/navigation/Tabs/DeckSettingsTab";

export default function Page() {
	const router = useRouter();

	const deckID = router.query.id as string;

	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: deck } = useSWR(
		deckID ? `/api/decks/${deckID}` : null,
		fetcher
	);
	const { data: group } = useSWR(
		deck?.groupID ? `/api/groups/${deck.groupID}` : null,
		fetcher
	);

	const tabs: { [tab: string]: ReactNode } = {
		cards: (
			<TabHeader
				id="CardsTabHeaderId"
				name="Cards"
				style="cards"
			></TabHeader>
		),
		statistics: (
			<TabHeader
				id="StatisticsTabHeaderId"
				name="Statistics"
				style="statistics"
			></TabHeader>
		),
		settings: (
			<TabHeader
				id="SettingsTabHeaderId"
				name="Settings"
				style="settings"
			></TabHeader>
		),
	};

	const [currentTab, setTab] = useState("cards");

	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<Authenticated>
				<div className="min-w-screen flex h-screen flex-col bg-eggshell">
					<Navbar login={true}></Navbar>
					{group && deck && (
						<div className="flex flex-col space-y-3 p-5 md:space-y-5 md:p-10">
							<Header
								id={"deckPageHeaderId"}
								group={group}
								deck={{ ...deck, deckID: deck.deckID }}
							></Header>
							<TabBar
								id="deckTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setTab}
							></TabBar>
							<div>
								{{
									cards: (
										<CardsTab
											deckID={deck?.deckID}
										></CardsTab>
									),
									settings: (
										<DeckSettingsTab
											groupID={deck?.groupID}
											deck={deck}
										></DeckSettingsTab>
									),
									statistics: <StatisticsTab></StatisticsTab>,
								}[currentTab] || <div>Error</div>}
							</div>
						</div>
					)}
				</div>
			</Authenticated>
		</div>
	);
}
