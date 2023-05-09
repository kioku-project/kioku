import CalendarHeatmap from "react-calendar-heatmap";
import "react-calendar-heatmap/dist/styles.css";
import React from "react";
import { Tooltip } from "react-tooltip";
import "react-tooltip/dist/react-tooltip.css";
import GroupOverviewTile from "../components/group/GroupOverviewTile";
import Navigation from "../components/navigation/Navigation";
import Authenticated from "../components/accessControl/Authenticated";

export default function Home() {
	return (
		<Authenticated>
			<Navigation>
				<h1
					className="text-xl font-bold"
					data-tooltip-id="my-tooltip"
					data-tooltip-content="Hello world!"
				>
					Hello 👋
				</h1>
				<p>You have 10 cards in 2 different decks to learn</p>
				<div className="m-4 mx-24 flex justify-center">
					<CalendarHeatmap
						startDate={new Date("2023-01-01")}
						endDate={new Date("2023-12-31")}
						showWeekdayLabels={true}
						transformDayElement={(element, value, index) =>
							React.cloneElement(element, {
								rx: 2,
								ry: 2,
								"data-tooltip-id": "heatmap-tooltip",
								"data-tooltip-content": `${
									value?.count || 0
								} cards reviewed`,
								style: {
									stroke:
										value?.date ==
											new Date()
												.toISOString()
												.split("T")[0] && "black",
								},
							})
						}
						values={[
							{ date: "2023-01-01", count: 12 },
							{ date: "2023-01-22", count: 122 },
							{ date: "2023-01-30", count: 38 },
							{ date: "2023-04-06", count: 38 },
							// ...and so on
						]}
					/>
				</div>
				<GroupOverviewTile
					name="Group1"
					decks={[
						{ name: "Deck1", count: 1 },
						{ name: "Deck2", count: 2 },
					]}
				/>
				<GroupOverviewTile
					name="Group2"
					decks={[
						{ name: "Deck1", count: 1 },
						{ name: "Deck2", count: 2 },
					]}
				/>
				<GroupOverviewTile
					name="Group3"
					decks={[
						{ name: "Deck1", count: 1 },
						{ name: "Deck2", count: 2 },
					]}
				/>
				<Tooltip id="heatmap-tooltip" />
			</Navigation>
		</Authenticated>
	);
}
