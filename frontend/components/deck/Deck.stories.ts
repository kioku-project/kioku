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

export const Default: Story = {
	args: {
		deck: { deckID: "D-12345678", deckName: "Example Deck" },
	},
};

export const CreateDeck: Story = {
	args: {},
};
