import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { GetStaticProps } from "next";
import Head from "next/head";
import React, { ReactNode, useState } from "react";

import { FetchHeader } from "@/components/layout/Header";
import { DashboardTab } from "@/components/navigation/Tabs/DashboardTab";
import { DecksTab } from "@/components/navigation/Tabs/DecksTab";
import { GroupsTab } from "@/components/navigation/Tabs/GroupsTab";
import { StatisticsTab } from "@/components/navigation/Tabs/StatisticsTab";
import { TabBar } from "@/components/navigation/Tabs/TabBar";
import { TabHeader } from "@/components/navigation/Tabs/TabHeader";
import { UserSettingsTab } from "@/components/navigation/Tabs/UserSettingsTab";
import { loadCatalog } from "@/pages/_app";
import { useGroups, useUser, useUserDue } from "@/util/swr";

export const getStaticProps: GetStaticProps = async (ctx) => {
	const translation = await loadCatalog(ctx.locale!);
	return {
		props: {
			translation,
		},
	};
};

export default function Home() {
	const { user } = useUser();
	const { due } = useUserDue();
	const { groups } = useGroups();

	const homeGroup = groups?.filter((group) => group.isDefault)[0];

	const { _ } = useLingui();

	const tabs: { [tab: string]: ReactNode } = {
		dashboard: (
			<TabHeader
				id="dashboardTabHeaderId"
				name={_(msg`Dashboard`)}
				icon="Home"
			/>
		),
		decks: (
			<TabHeader
				id="decksTabHeaderId"
				name={_(msg`Decks`)}
				icon="Layers"
			/>
		),
		groups: (
			<TabHeader
				id="groupTabHeaderId"
				name={_(msg`Groups`)}
				icon="Users"
			/>
		),
		settings: (
			<TabHeader
				id="SettingsTabHeaderId"
				name={_(msg`Settings`)}
				icon="Settings"
			/>
		),
	};

	const [currentTab, setCurrentTab] = useState("dashboard");

	return (
		<div className="flex flex-1 overflow-auto">
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
				<link
					rel="alternate"
					hrefLang="en"
					href="https://app.kioku.dev/"
				/>
				<link
					rel="alternate"
					hrefLang="de"
					href="https://app.kioku.dev/de"
				/>
			</Head>
			<div className="min-w-screen flex flex-1 flex-col bg-eggshell">
				{user && groups && (
					<div className="flex h-full flex-col md:space-y-5">
						<FetchHeader
							id="userPageHeaderId"
							user={{ ...user, ...due }}
						/>
						<div className="flex h-full flex-1 flex-col-reverse justify-between space-y-5 overflow-auto md:flex-col md:justify-normal">
							<TabBar
								id="deckTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setCurrentTab}
							/>
							<div className="h-full overflow-auto px-5 md:px-10">
								{{
									decks: homeGroup && (
										<DecksTab group={homeGroup} />
									),
									groups: <GroupsTab groups={groups} />,
									settings: <UserSettingsTab user={user} />,
									statistics: <StatisticsTab />,
									dashboard: homeGroup && <DashboardTab />,
								}[currentTab] ?? <div>Error</div>}
							</div>
						</div>
					</div>
				)}
			</div>
		</div>
	);
}
