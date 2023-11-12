import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import Head from "next/head";
import { useRouter } from "next/router";
import React, { ReactNode, useEffect, useState } from "react";
import "react-toastify/dist/ReactToastify.css";

import Authenticated from "../../../components/accessControl/Authenticated";
import { FetchHeader } from "../../../components/layout/Header";
import { DecksTab } from "../../../components/navigation/Tabs/DecksTab";
import { GroupSettingsTab } from "../../../components/navigation/Tabs/GroupSettingsTab";
import { MembersTab } from "../../../components/navigation/Tabs/MembersTab";
import { StatisticsTab } from "../../../components/navigation/Tabs/StatisticsTab";
import { TabBar } from "../../../components/navigation/Tabs/TabBar";
import { TabHeader } from "../../../components/navigation/Tabs/TabHeader";
import { useGroup } from "../../../util/swr";

export default function Page() {
	const router = useRouter();
	const { _ } = useLingui();

	const [groupID, setGroupID] = useState<string>();
	useEffect(() => {
		setGroupID(router.query.id as string);
	}, [groupID, router]);
	const { group } = useGroup(groupID ? groupID : "");

	const tabs: { [tab: string]: ReactNode } = {
		decks: (
			<TabHeader
				id="DecksTabHeaderId"
				name={_(msg`Decks`)}
				style="decks"
			></TabHeader>
		),
		user: (
			<TabHeader
				id="UserTabHeaderId"
				name={_(msg`User`)}
				style="user"
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

	const [currentTab, setCurrentTab] = useState("decks");

	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
				<link rel="alternate" hrefLang="en" href={`https://app.kioku.dev/group/${groupId}`} />
				<link rel="alternate" hrefLang="de" href={`https://app.kioku.dev/de/group/${groupId}`} />
			</Head>

			<Authenticated>
				<div className="min-w-screen flex flex-1 flex-col bg-eggshell">
					{group && (
						<div className="space-y-5 p-10">
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
