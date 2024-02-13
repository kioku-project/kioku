import React, { ReactNode } from "react";

import { Text } from "@/components/Text";

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
	return (
		<div
			className={`flex flex-row justify-between ${className}`}
			{...props}
		>
			<div className="flex w-full flex-row md:relative">
				{Object.keys(tabs).map((tab) => (
					<Text
						textSize="xs"
						key={tab}
						className={`flex-1 cursor-pointer border-kiokuDarkBlue p-3 font-bold transition md:flex-initial ${
							currentTab === tab
								? "border-b-2 text-kiokuDarkBlue"
								: "border-none text-kiokuLightBlue"
						}`}
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
