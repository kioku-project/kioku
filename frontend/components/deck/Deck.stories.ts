import type { Meta, StoryObj } from "@storybook/react";

import { Deck } from "@/components/deck/Deck";

const meta: Meta<typeof Deck> = {
	title: "Components/Deck",
	component: Deck,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof Deck>;

export const PrivateDeck: Story = {
	args: {
		deck: {
			deckID: "D-12345678",
			deckName: "Example Deck",
			deckType: "PRIVATE",
			groupID: "G-12345678",
		},
	},
};

export const PublicDeck: Story = {
	args: {
		deck: {
			deckID: "D-12345678",
			deckName: "Example Deck",
			deckType: "PUBLIC",
			groupID: "G-12345678",
		},
	},
};

export const CardsDue: Story = {
	args: {
		deck: {
			deckID: "D-12345678",
			deckName: "Example Deck",
			deckType: "PUBLIC",
			groupID: "G-12345678",
			dueCards: 8,
		},
	},
};

export const MoreStats: Story = {
	args: {
		deck: {
			deckID: "D-12345678",
			deckName: "Example Deck",
			deckType: "PUBLIC",
			groupID: "G-12345678",
			dueCards: 8,
		},
		stats: [
			{
				icon: "Copy",
				header: "34 Cards",
			},
			{
				icon: "PieChart",
				header: "75%",
			},
		],
	},
};

export const WithNotification: Story = {
	args: {
		deck: {
			deckID: "D-12345678",
			deckName: "Example Deck",
			deckType: "PUBLIC",
			groupID: "G-12345678",
		},
		deckNotification: {
			icon: "Award",
			header: "You're close!",
			description: "You've almost completed this deck! (75%)",
		},
	},
};
