import React, { ReactNode, useState } from "react";
import Authenticated from "../components/accessControl/Authenticated";
import { Navbar } from "../components/navigation/Navbar";
import Head from "next/head";
import useSWR from "swr";
import { authedFetch } from "../util/reauth";
import { Header } from "../components/layout/Header";
import { TabBar } from "../components/navigation/Tabs/TabBar";
import { TabHeader } from "../components/navigation/Tabs/TabHeader";
import { StatisticsTab } from "../components/navigation/Tabs/StatisticsTab";
import { GroupsTab } from "../components/navigation/Tabs/GroupsTab";
import { DecksTab } from "../components/navigation/Tabs/DecksTab";
import { UserSettingsTab } from "../components/navigation/Tabs/UserSettingsTab";
import { InvitationsTab } from "../components/navigation/Tabs/InvitationsTabs";
import { Group } from "../types/Group";

export default function Home() {
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: groups } = useSWR("/api/groups", fetcher);
	const { data: user } = useSWR("/api/user", fetcher);
	const { data: due } = useSWR("/api/user/dueCards", fetcher);

	const tabs: { [tab: string]: ReactNode } = {
		decks: (
			<TabHeader
				id="decksTabHeaderId"
				name="Decks"
				style="decks"
			></TabHeader>
		),
		groups: (
			<TabHeader
				id="groupTabHeaderId"
				name="Groups"
				style="groups"
			></TabHeader>
		),
		invitations: (
			<TabHeader
				id="invitationTabHeaderId"
				name="Invitations"
				style="invitations"
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
				<div className="min-w-screen flex h-screen select-none flex-col bg-eggshell">
					<Navbar login={true}></Navbar>
					{user && groups && (
						<div className="flex flex-col space-y-3 p-5 md:space-y-5 md:p-10">
							<Header
								id="userPageHeaderId"
								user={{ ...user, ...due }}
							></Header>
							<TabBar
								id="deckTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setTab}
							></TabBar>
							<div className="">
								{{
									decks: (
										<DecksTab
											group={
												groups?.groups.filter(
													(group: Group) =>
														group.isDefault
												)[0]
											}
										></DecksTab>
									),
									groups: (
										<GroupsTab
											groups={groups.groups}
										></GroupsTab>
									),
									invitations: (
										<InvitationsTab></InvitationsTab>
									),
									statistics: <StatisticsTab></StatisticsTab>,
									settings: (
										<UserSettingsTab
											user={user}
										></UserSettingsTab>
									),
								}[currentTab] || <div>Fehler</div>}
							</div>
						</div>
					)}
				</div>
			</Authenticated>
		</div>
	);
}
