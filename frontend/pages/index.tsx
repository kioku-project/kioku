import Head from "next/head";
import React, { ReactNode, useState } from "react";
import useSWR from "swr";

import Authenticated from "../components/accessControl/Authenticated";
import { FetchHeader } from "../components/layout/Header";
import { DecksTab } from "../components/navigation/Tabs/DecksTab";
import { GroupsTab } from "../components/navigation/Tabs/GroupsTab";
import { InvitationsTab } from "../components/navigation/Tabs/InvitationsTabs";
import { StatisticsTab } from "../components/navigation/Tabs/StatisticsTab";
import { TabBar } from "../components/navigation/Tabs/TabBar";
import { TabHeader } from "../components/navigation/Tabs/TabHeader";
import { UserSettingsTab } from "../components/navigation/Tabs/UserSettingsTab";
import { Group as GroupType } from "../types/Group";
import { Invitation } from "../types/Invitation";
import { authedFetch } from "../util/reauth";

export default function Home() {
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: groups } = useSWR("/api/groups", fetcher);
	const { data: user } = useSWR("/api/user", fetcher);
	const { data: due } = useSWR("/api/user/dueCards", fetcher);
	const { data: invitations } = useSWR<{
		groupInvitation: Invitation[];
	}>(`/api/user/invitations`, fetcher);

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
				notificationBadgeContent={`${
					invitations?.groupInvitation?.length || ""
				}`}
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
					{user && groups && (
						<div className="flex flex-col space-y-3 p-5 md:space-y-5 md:p-10">
							<FetchHeader
								id="userPageHeaderId"
								user={{ ...user, ...due }}
							/>
							<TabBar
								id="deckTabBarId"
								tabs={tabs}
								currentTab={currentTab}
								setTab={setCurrentTab}
							></TabBar>
							<div>
								{{
									decks: (
										<DecksTab
											group={
												groups.groups.filter(
													(group: GroupType) =>
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
									invitations: invitations && (
										<InvitationsTab
											invitations={
												invitations.groupInvitation
											}
										></InvitationsTab>
									),
									statistics: <StatisticsTab></StatisticsTab>,
									settings: (
										<UserSettingsTab
											user={user}
										></UserSettingsTab>
									),
								}[currentTab] ?? <div>Error</div>}
							</div>
						</div>
					)}
				</div>
			</Authenticated>
		</div>
	);
}
