import clsx from "clsx";
import React, { ReactNode } from "react";

import { Text } from "@/components/Text";
import { OperatingSystem, getOperatingSystem } from "@/util/client";

interface TabBarProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * Tabs to display
	 */
	tabs: { [tab: string]: ReactNode };
	/**
	 * Current tab
	 */
	currentTab: string;
	/**
	 * Set tab
	 */
	setTab: (tab: string) => void;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a TabBar
 */
export const TabBar = ({
	tabs,
	currentTab,
	setTab,
	className = "",
	...props
}: TabBarProps) => {
	const os = getOperatingSystem(navigator.userAgent);

	return (
		<div
			className={`flex flex-row justify-between md:px-10 ${className}`}
			{...props}
		>
			<div
				className={clsx(
					"flex w-full flex-row md:relative md:border-0 md:pb-0",
					os === OperatingSystem.IOS && "pb-3"
				)}
			>
				{Object.keys(tabs).map((tab) => (
					<Text
						textSize="xs"
						key={tab}
						className={clsx(
							"flex-1 cursor-pointer border-t-2 p-3 font-bold transition md:flex-initial md:border-t-0",
							currentTab === tab
								? "border-kiokuDarkBlue text-kiokuDarkBlue md:border-b-2"
								: "border-kiokuLightBlue text-kiokuLightBlue"
						)}
						onClick={() => {
							setTab(tab);
						}}
					>
						{tabs[tab]}
					</Text>
				))}
			</div>
		</div>
	);
};
