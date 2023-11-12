import Head from "next/head";
import { useRouter } from "next/router";
import React, { ReactNode, useEffect, useState } from "react";
import "react-toastify/dist/ReactToastify.css";
import useSWR from "swr";

import Authenticated from "../../../components/accessControl/Authenticated";
import { FetchHeader } from "../../../components/layout/Header";
import { DecksTab } from "../../../components/navigation/Tabs/DecksTab";
import { GroupSettingsTab } from "../../../components/navigation/Tabs/GroupSettingsTab";
import { MembersTab } from "../../../components/navigation/Tabs/MembersTab";
import { StatisticsTab } from "../../../components/navigation/Tabs/StatisticsTab";
import { TabBar } from "../../../components/navigation/Tabs/TabBar";
import { TabHeader } from "../../../components/navigation/Tabs/TabHeader";
import { authedFetch } from "../../../util/reauth";

export default function Page() {
	const router = useRouter();

	const [groupId, setGroupId] = useState<string>();
	useEffect(() => {
		setGroupId(router.query.id as string);
	}, [groupId, router]);

	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: group } = useSWR(
		groupId ? `/api/groups/${groupId}` : null,
		fetcher
	);

	const tabs: { [tab: string]: ReactNode } = {
		decks: (
			<TabHeader
				id="DecksTabHeaderId"
				name="Decks"
				style="decks"
			></TabHeader>
		),
		user: (
			<TabHeader
				id="UserTabHeaderId"
				name="User"
				style="user"
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

	const [currentTab, setCurrentTab] = useState("decks");

	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<Authenticated>
				<div className="min-w-screen flex flex-1 flex-col bg-eggshell">
					{group && (
						<div className="space-y-5 px-5 py-1 md:px-10 md:py-3">
							<FetchHeader id="groupPageHeaderId" group={group} />
							<TabBar
								id="groupTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setCurrentTab}
							></TabBar>
							<div>
								{{
									decks: <DecksTab group={group}></DecksTab>,
									user: (
										<MembersTab group={group}></MembersTab>
									),
									settings: (
										<GroupSettingsTab
											group={group}
										></GroupSettingsTab>
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
