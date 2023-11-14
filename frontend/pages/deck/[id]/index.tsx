import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { GetStaticProps } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import React, { ReactNode, useState } from "react";
import "react-toastify/dist/ReactToastify.css";
import useSWR from "swr";

import { loadCatalog } from "@/pages/_app";

import Authenticated from "../../../components/accessControl/Authenticated";
import { FetchHeader } from "../../../components/layout/Header";
import { CardsTab } from "../../../components/navigation/Tabs/CardsTab";
import { DeckSettingsTab } from "../../../components/navigation/Tabs/DeckSettingsTab";
import { StatisticsTab } from "../../../components/navigation/Tabs/StatisticsTab";
import { TabBar } from "../../../components/navigation/Tabs/TabBar";
import { TabHeader } from "../../../components/navigation/Tabs/TabHeader";
import { authedFetch } from "../../../util/reauth";

export const getStaticProps: GetStaticProps = async (ctx) => {
	const translation = await loadCatalog(ctx.locale!);
	return {
		props: {
			translation,
		},
	};
};

export default function Page() {
	const router = useRouter();

	const { _ } = useLingui();

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
				name={_(msg`Cards`)}
				style="cards"
			></TabHeader>
		),
		statistics: (
			<TabHeader
				id="StatisticsTabHeaderId"
				name={_(msg`Statistics`)}
				style="statistics"
			></TabHeader>
		),
		settings: (
			<TabHeader
				id="SettingsTabHeaderId"
				name={_(msg`Settings`)}
				style="settings"
			></TabHeader>
		),
	};

	const [currentTab, setCurrentTab] = useState("cards");

	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
				<link
					rel="alternate"
					hrefLang="en"
					href={`https://app.kioku.dev/deck/${deckID}`}
				/>
				<link
					rel="alternate"
					hrefLang="de"
					href={`https://app.kioku.dev/de/deck/${deckID}`}
				/>
			</Head>
			<Authenticated>
				<div className="min-w-screen flex flex-1 flex-col bg-eggshell">
					{group && deck && (
						<div className="flex h-full flex-col space-y-3 overflow-y-auto p-5 md:space-y-5 md:p-10">
							<FetchHeader
								id={"deckPageHeaderId"}
								group={group}
								deck={deck}
							/>
							<TabBar
								id="deckTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setCurrentTab}
							></TabBar>
							<div className="h-full overflow-y-auto">
								{{
									cards: (
										<CardsTab
											deck={{
												...deck,
												groupRole: group.groupRole,
											}}
										></CardsTab>
									),
									settings: (
										<DeckSettingsTab
											group={group}
											deck={deck}
										></DeckSettingsTab>
									),
									statistics: <StatisticsTab></StatisticsTab>,
								}[currentTab] ?? <div>Error</div>}
							</div>
						</div>
					)}
				</div>
			</Authenticated>
		</div>
	);
}
