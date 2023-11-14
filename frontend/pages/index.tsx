import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { GetStaticProps } from "next";
import Head from "next/head";
import React, { ReactNode, useState } from "react";

import Authenticated from "../components/accessControl/Authenticated";
import { FetchHeader } from "../components/layout/Header";
import { DecksTab } from "../components/navigation/Tabs/DecksTab";
import { GroupsTab } from "../components/navigation/Tabs/GroupsTab";
import { InvitationsTab } from "../components/navigation/Tabs/InvitationsTabs";
import { StatisticsTab } from "../components/navigation/Tabs/StatisticsTab";
import { TabBar } from "../components/navigation/Tabs/TabBar";
import { TabHeader } from "../components/navigation/Tabs/TabHeader";
import { UserSettingsTab } from "../components/navigation/Tabs/UserSettingsTab";
import { useGroups, useInvitations, useUser, useUserDue } from "../util/swr";
import { loadCatalog } from "./_app";

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
	const { invitations } = useInvitations();
	const { groups } = useGroups();

	const { _ } = useLingui();

	const tabs: { [tab: string]: ReactNode } = {
		decks: (
			<TabHeader
				id="decksTabHeaderId"
				name={_(msg`Decks`)}
				style="decks"
			></TabHeader>
		),
		groups: (
			<TabHeader
				id="groupTabHeaderId"
				name={_(msg`Groups`)}
				style="groups"
			></TabHeader>
		),
		invitations: (
			<TabHeader
				id="invitationTabHeaderId"
				name={_(msg`Invitations`)}
				style="invitations"
				notificationBadgeContent={`${invitations?.length ?? ""}`}
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
												groups.filter(
													(group) => group.isDefault
												)[0]
											}
										></DecksTab>
									),
									groups: (
										<GroupsTab groups={groups}></GroupsTab>
									),
									invitations: invitations && (
										<InvitationsTab
											invitations={invitations}
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
