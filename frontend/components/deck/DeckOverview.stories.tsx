import type { Meta, StoryObj } from "@storybook/react";
import DeckOverview from "./DeckOverview";

const meta: Meta<typeof DeckOverview> = {
	title: "Components/DeckOverview",
	component: DeckOverview,
	tags: ["autodocs"],
	args: {
		id: "DeckGroupId",
	},
};

export default meta;
type Story = StoryObj<typeof DeckOverview>;

export const Default: Story = {
	args: {
		name: "Example",
		decks: [
			{ name: "Deck 0", count: 0 },
			{ name: "Deck 1", count: 1 },
			{ name: "Deck 2", count: 2 },
		],
	},
};

export const NoDecks: Story = {
	args: {
		name: "Example Group",
		decks: [],
	},
};
