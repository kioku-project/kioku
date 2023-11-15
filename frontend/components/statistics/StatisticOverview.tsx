import React from "react";

import { Statistic } from "./Statistic";
import { useLingui } from "@lingui/react";
import { msg } from "@lingui/macro";

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
	const { _ } = useLingui()
	return (
		<div id={id} className="flex flex-row justify-between rounded-lg">
			<Statistic
				id={"statisticId"}
				header={_(msg`Cards learned`)}
				value={"176"}
				separator={_(msg`from`)}
				reference={"200"}
				change={12}
				className="border-r-2 border-kiokuLightBlue"
			></Statistic>
			<Statistic
				id={"statisticId"}
				header={_(msg`Hit Rate`)}
				value={"34%"}
				change={2}
			></Statistic>
			<Statistic
				id={"statisticId"}
				header={_(msg`Test`)}
				value={"0"}
				separator={_(msg`from`)}
				reference={"100"}
				change={-3}
				className="border-l-2 border-kiokuLightBlue"
			></Statistic>
		</div>
	);
};
