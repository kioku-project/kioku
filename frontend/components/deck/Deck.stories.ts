import type { Meta, StoryObj } from "@storybook/react";
import Deck from "./Deck";

const meta: Meta<typeof Deck> = {
	title: "Components/Deck",
	component: Deck,
	tags: ["autodocs"],
	args: {
		group: { groupID: "G-12345678", groupName: "Example Group" },
	},
};

export default meta;
type Story = StoryObj<typeof Deck>;

export const NoCardsDue: Story = {
	args: {
		deck: { deckID: "D-12345678", deckName: "Example Deck", dueCards: 0 },
	},
};

export const WithCardsDue: Story = {
	args: {
		deck: { deckID: "D-12345678", deckName: "Example Deck", dueCards: 1 },
	},
};

export const ManyCardsDue: Story = {
	args: {
		deck: { deckID: "D-12345678", deckName: "Example Deck", dueCards: 100 },
	},
};

export const CreateDeck: Story = {
	args: {},
};
