import React from "react";

import { Statistic } from "./Statistic";

interface StatisticOverviewProps {
	/**
	 * Unique identifier
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
export const StatisticOverview = ({
	id,
	className,
}: StatisticOverviewProps) => {
	return (
		<div id={id} className="flex flex-row justify-between rounded-lg">
			<Statistic
				id={"statisticId"}
				header="Cards learned"
				value={"176"}
				separator="from"
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
				separator="from"
				reference={"100"}
				change={-3}
				className="border-l-2 border-kiokuLightBlue"
			></Statistic>
		</div>
	);
};
