import React, { ReactNode, useState } from "react";
import { Text } from "../../Text";

interface TabBarProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * tabs to display
	 */
	tabs: { [tab: string]: ReactNode };
	/**
	 * currentTab
	 */
	currentTab: string;
	/**
	 * setTab
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
	className,
	...props
}: TabBarProps) => {
	return (
		<div
			className={`flex flex-row justify-between ${className ?? ""}`}
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
							currentTab == tab
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
