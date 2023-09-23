import type { Meta, StoryObj } from "@storybook/react";

import DeckOverview from "./DeckOverview";

const meta: Meta<typeof DeckOverview> = {
	title: "Components/DeckOverview",
	component: DeckOverview,
	tags: ["autodocs"],
	args: {
		group: { groupID: "G-12345678", groupName: "Example Group" },
	},
};

export default meta;
type Story = StoryObj<typeof DeckOverview>;

export const Default: Story = {
	args: {},
};

export const NoDecks: Story = {
	args: {},
};
