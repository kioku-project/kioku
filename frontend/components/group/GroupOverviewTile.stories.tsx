import type { Meta, StoryObj } from "@storybook/react";
import GroupOverviewTile from "./GroupOverviewTile";

const meta: Meta<typeof GroupOverviewTile> = {
	title: "Group/OverviewTile",
	component: GroupOverviewTile,
	tags: ["autodocs"],
	args: {
		id: "GroupId",
	},
};

export default meta;
type Story = StoryObj<typeof GroupOverviewTile>;

export const Default: Story = {
	args: {
		name: "Example Group",
		decks: [
			{ name: "Deck1", count: 1 },
			{ name: "Deck2", count: 2 },
			{ name: "Deck3", count: 3 },
		],
	},
};

export const NoDecks: Story = {
	args: {
		name: "Example Group",
		decks: [],
	},
};
