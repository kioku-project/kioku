import Head from "next/head";
import { PropsWithChildren, useState } from "react";
import Navbar from "./Navbar";
import GroupAsideTile from "../group/GroupAsideTile";
import "react-calendar-heatmap/dist/styles.css";
import "react-tooltip/dist/react-tooltip.css";

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
