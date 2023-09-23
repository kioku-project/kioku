import React from "react";

import { Statistic } from "./Statistic";

interface StatisticOverviewProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for grouping statistics
 */
export const StatisticOverview = ({ className }: StatisticOverviewProps) => {
	return (
		<div className="flex flex-row justify-between rounded-lg bg-gray-200">
			<Statistic
				id={"statisticId"}
				header="Cards learned"
				value={"176"}
				seperator="from"
				reference={"200"}
				change={12}
				className="border-r-2 border-kiokuLightBlue"
			></Statistic>
			<Statistic
				id={"statisticId"}
				header="Hit Rate"
				value={"34%"}
				change={2}
			></Statistic>
			<Statistic
				id={"statisticId"}
				header="Test"
				value={"0"}
				seperator="from"
				reference={"100"}
				change={-3}
				className="border-l-2 border-kiokuLightBlue"
			></Statistic>
		</div>
	);
};
