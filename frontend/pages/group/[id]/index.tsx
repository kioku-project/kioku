import Head from "next/head";
import "react-toastify/dist/ReactToastify.css";
import Authenticated from "../../../components/accessControl/Authenticated";
import { Navbar } from "../../../components/navigation/Navbar";
import { useRouter } from "next/router";
import { ReactNode, useEffect, useState } from "react";
import { StatisticsTab } from "../../../components/navigation/Tabs/StatisticsTab";
import { authedFetch } from "../../../util/reauth";
import useSWR from "swr";
import { TabHeader } from "../../../components/navigation/Tabs/TabHeader";
import { Header } from "../../../components/layout/Header";
import { MembersTab } from "../../../components/navigation/Tabs/MembersTab";
import { DecksTab } from "../../../components/navigation/Tabs/DecksTab";
import { TabBar } from "../../../components/navigation/Tabs/TabBar";
import { GroupSettingsTab } from "../../../components/navigation/Tabs/GroupSettingTab";
import React from "react";

export default function Page() {
	const router = useRouter();

	const [groupID, setGroupId] = useState<string>();
	useEffect(() => {
		setGroupId(router.query.id as string);
	}, [groupID, router]);

	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: group } = useSWR(
		groupID ? `/api/groups/${groupID}` : null,
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

	const [currentTab, setTab] = useState("decks");

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
					{group && (
						<div className="space-y-5 p-10">
							<Header
								id="groupPageHeaderId"
								group={group}
							></Header>
							<TabBar
								id="groupTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setTab}
							></TabBar>
							<div className="">
								{{
									decks: <DecksTab group={group}></DecksTab>,
									user: (
										<MembersTab
											groupID={group.groupID}
										></MembersTab>
									),
									settings: (
										<GroupSettingsTab
											group={group}
										></GroupSettingsTab>
									),
									statistics: <StatisticsTab></StatisticsTab>,
								}[currentTab] || <div>Fehler</div>}
							</div>
						</div>
					)}
				</div>
			</Authenticated>
		</div>
	);
}
