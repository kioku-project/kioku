import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";

import { Section } from "@/components/layout/Section";
import { StatisticOverview } from "@/components/statistics/StatisticOverview";

interface StatisticsTabProps {
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the StatisticsTab
 */
export const StatisticsTab = ({ className = "" }: StatisticsTabProps) => {
	const { _ } = useLingui();
	return (
		<div className={`space-y-5 ${className}`}>
			<Section
				id="personalStatisticsSectionId"
				header={_(msg`Personal Statistics`)}
			>
				<StatisticOverview id="personalStatisticsOverviewId" />
			</Section>
			<Section
				id="groupStatisticsSectionId"
				header={_(msg`Group Statistics`)}
			>
				<StatisticOverview id="groupStatisticsOverviewId" />
			</Section>
		</div>
	);
};
