import type { Meta, StoryObj } from "@storybook/react";
import Deck from "./Deck";

const meta: Meta<typeof Deck> = {
	title: "Components/Deck",
	component: Deck,
	tags: ["autodocs"],
	args: {
		id: "DeckId",
	},
};

export default meta;
type Story = StoryObj<typeof Deck>;

export const NoCardsDue: Story = {
	args: {
		deck: { name: "Example Deck", count: 0 },
	},
};

export const WithCardsDue: Story = {
	args: {
		deck: { name: "Example Deck", count: 1 },
	},
};

export const ManyCardsDue: Story = {
	args: {
		deck: { name: "Example Deck", count: 100 },
	},
};

export const CreateDeck: Story = {
	args: {},
};
