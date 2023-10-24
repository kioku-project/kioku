import type { Meta, StoryObj } from "@storybook/react";

import { Deck } from "./Deck";

const meta: Meta<typeof Deck> = {
	title: "Components/Deck",
	component: Deck,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof Deck>;

export const Default: Story = {
	args: {
		deck: { deckID: "D-12345678", deckName: "Example Deck" },
		group: { groupID: "G-12345678", groupName: "Example Group" },
	},
};

export const CreateDeck: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Example Group",
			groupRole: "ADMIN",
		},
	},
};

export const EmptyGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Example Group",
			groupRole: "READ",
			isEmpty: true,
		},
	},
};
