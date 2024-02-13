import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { GetStaticProps } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import React, { ReactNode, useState } from "react";
import "react-toastify/dist/ReactToastify.css";

import { FetchHeader } from "@/components/layout/Header";
import { CardsTab } from "@/components/navigation/Tabs/CardsTab";
import { DeckSettingsTab } from "@/components/navigation/Tabs/DeckSettingsTab";
import { StatisticsTab } from "@/components/navigation/Tabs/StatisticsTab";
import { TabBar } from "@/components/navigation/Tabs/TabBar";
import { TabHeader } from "@/components/navigation/Tabs/TabHeader";
import { loadCatalog } from "@/pages/_app";
import { useDeck, useGroup } from "@/util/swr";

export const getServerSideProps: GetStaticProps = async (ctx) => {
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
	const { deck } = useDeck(deckID);
	const { group } = useGroup(deck?.groupID);

	const tabs: { [tab: string]: ReactNode } = {
		cards: (
			<TabHeader id="CardsTabHeaderId" name={_(msg`Cards`)} icon="Copy" />
		),
		settings: (
			<TabHeader
				id="SettingsTabHeaderId"
				name={_(msg`Settings`)}
				icon="Settings"
			/>
		),
	};

	const [currentTab, setCurrentTab] = useState("cards");

	return (
		<div className="flex flex-1 overflow-auto">
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

			<div className="min-w-screen flex flex-1 flex-col overflow-auto bg-eggshell">
				{group && deck && (
					<div className="flex h-full flex-col px-5 py-1 md:space-y-5 md:px-10 md:py-3">
						<FetchHeader
							id={"deckPageHeaderId"}
							group={group}
							deck={deck}
						/>
						<div className="flex h-full flex-1 flex-col-reverse justify-between space-y-5 overflow-auto md:flex-col md:justify-normal">
							<TabBar
								id="deckTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setCurrentTab}
							/>
							<div className="overflow-auto">
								{{
									cards: (
										<CardsTab
											deck={{
												...deck,
											}}
										/>
									),
									settings: (
										<DeckSettingsTab
											group={group}
											deck={deck}
										/>
									),
									statistics: <StatisticsTab />,
								}[currentTab] ?? <div>Error</div>}
							</div>
						</div>
					</div>
				)}
			</div>
		</div>
	);
}
