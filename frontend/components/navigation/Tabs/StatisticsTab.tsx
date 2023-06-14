import { Section } from "../../layout/Section";
import { StatisticOverview } from "../../statistics/StatisticOverview";

interface StatisticsTabProps {
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the StatisticsTab
 */
export const StatisticsTab = ({ className }: StatisticsTabProps) => {
	return (
		<div className="space-y-5">
			<Section
				id="personalStatisticsSectionId"
				header="Personal Statistics"
			>
				<StatisticOverview id="personalStatisticsOverviewId"></StatisticOverview>
			</Section>
			<Section id="groupStatisticsSectionId" header="Group Statistics">
				<StatisticOverview id="groupStatisticsOverviewId"></StatisticOverview>
			</Section>
		</div>
	);
};
