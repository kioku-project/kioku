import React, { ReactNode } from "react";

import { Text } from "../../Text";

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
	 * CurrentTab
	 */
	currentTab: string;
	/**
	 * SetTab
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
			<div className="flex flex-row">
				{Object.keys(tabs).map((tab) => (
					<Text
						size="xs"
						onClick={() => {
							setTab(tab);
						}}
						key={tab}
						className={`border-kiokuDarkBlue p-3 font-bold transition hover:cursor-pointer ${
							currentTab === tab
								? "border-b-2 text-kiokuDarkBlue"
								: "border-none text-kiokuLightBlue"
						}`}
					>
						{tabs[tab]}
					</Text>
				))}
			</div>
		</div>
	);
};
