import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { GetStaticProps } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import React, { ReactNode, useEffect, useState } from "react";

import { FetchHeader } from "@/components/layout/Header";
import { DecksTab } from "@/components/navigation/Tabs/DecksTab";
import { GroupSettingsTab } from "@/components/navigation/Tabs/GroupSettingsTab";
import { MembersTab } from "@/components/navigation/Tabs/MembersTab";
import { StatisticsTab } from "@/components/navigation/Tabs/StatisticsTab";
import { TabBar } from "@/components/navigation/Tabs/TabBar";
import { TabHeader } from "@/components/navigation/Tabs/TabHeader";
import { loadCatalog } from "@/pages/_app";
import { useGroup } from "@/util/swr";

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

	const [groupID, setGroupID] = useState<string>();
	useEffect(() => {
		setGroupID(router.query.id as string);
	}, [groupID, router]);
	const { group } = useGroup(groupID);

	const tabs: { [tab: string]: ReactNode } = {
		decks: (
			<TabHeader
				id="DecksTabHeaderId"
				name={_(msg`Decks`)}
				icon="Layers"
			/>
		),
		user: (
			<TabHeader id="UserTabHeaderId" name={_(msg`User`)} icon="Users" />
		),
		settings: (
			<TabHeader
				id="SettingsTabHeaderId"
				name={_(msg`Settings`)}
				icon="Settings"
			/>
		),
	};

	const [currentTab, setCurrentTab] = useState("decks");

	return (
		<div className="flex flex-1 overflow-auto">
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
				<link
					rel="alternate"
					hrefLang="en"
					href={`https://app.kioku.dev/group/${groupID}`}
				/>
				<link
					rel="alternate"
					hrefLang="de"
					href={`https://app.kioku.dev/de/group/${groupID}`}
				/>
			</Head>

			<div className="min-w-screen flex flex-1 flex-col bg-eggshell">
				{group && (
					<div className="flex h-full flex-col md:space-y-5">
						<FetchHeader id="groupPageHeaderId" group={group} />
						<div className="flex h-full flex-1 flex-col-reverse justify-between space-y-5 overflow-auto md:flex-col md:justify-normal">
							<TabBar
								id="groupTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setCurrentTab}
							/>
							<div className="overflow-auto px-5 md:px-10">
								{{
									decks: <DecksTab group={group} />,
									user: <MembersTab group={group} />,
									settings: (
										<GroupSettingsTab group={group} />
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
