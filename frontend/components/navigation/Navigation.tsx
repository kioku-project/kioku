import Head from "next/head";
import Image from "next/image";
import { Router, useRouter } from "next/router";
import { PropsWithChildren, useContext, useEffect, useState } from "react";
import { UserContext } from "../../contexts/user";
import styles from "../styles/Home.module.css";
import Navbar from "./Navbar";
import GroupAsideTile from "../group/GroupAsideTile";
import CalendarHeatmap from "react-calendar-heatmap";
import "react-calendar-heatmap/dist/styles.css";
import React from "react";
import { Tooltip } from "react-tooltip";
import "react-tooltip/dist/react-tooltip.css";
import GroupOverviewTile from "../group/GroupOverviewTile";

export default function Navigation({ children }: PropsWithChildren) {
	const [asideOpen, toggleAside] = useState<boolean>(true);
	return (
		<div className="flex flex-col w-screen h-screen max-h-screen">
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<Navbar
				toggleMenuCallback={() => {
					toggleAside(!asideOpen);
				}}
			/>
			<div className="flex h-[calc(100%-3rem)] overflow-hidden">
				{asideOpen && (
					<aside className="bg-[#CCCCCC] h-full p-4 w-5/6 md:w-1/6 absolute md:relative z-10">
						<GroupAsideTile name="Group name" count={1} />
						<GroupAsideTile name="Group name" count={1} />
					</aside>
				)}
				<main className="p-8 w-full h-full overflow-y-scroll">
					{children}
				</main>
			</div>
		</div>
	);
}
