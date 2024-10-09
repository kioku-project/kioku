import type { Meta, StoryObj } from "@storybook/react";

import { Header } from "@/components/layout/Header";

const meta: Meta<typeof Header> = {
	title: "Layout/Header",
	component: Header,
	args: {},
};

export default meta;
type Story = StoryObj<typeof Header>;

export const User: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
		},
	},
};

export const UserWithDescription: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
			dueCards: 7,
			dueDecks: 2,
		},
	},
};

export const PublicDeck: Story = {
	args: {
		deck: {
			deckID: "D-12345678",
			deckName: "Test Deck",
			deckDescription: "Deck Description",
			deckType: "PUBLIC",
			deckRole: "ADMIN",
			groupID: "G-12345678",
		},
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupDescription: "Group Description",
			isDefault: false,
			groupType: "REQUEST",
			groupRole: "ADMIN",
		},
	},
};

export const PrivateDeck: Story = {
	args: {
		deck: {
			deckID: "D-12345678",
			deckName: "Test Deck",
			deckDescription: "Deck Description",
			deckType: "PRIVATE",
			deckRole: "ADMIN",
			groupID: "G-12345678",
		},
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupDescription: "Group Description",
			isDefault: false,
			groupType: "REQUEST",
			groupRole: "ADMIN",
		},
	},
};

export const OpenGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupDescription: "Group Description",
			isDefault: false,
			groupType: "OPEN",
			groupRole: "ADMIN",
		},
	},
};

export const RequestGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupDescription: "Group Description",
			isDefault: false,
			groupType: "REQUEST",
			groupRole: "ADMIN",
		},
	},
};

export const ClosedGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupDescription: "Group Description",
			isDefault: false,
			groupType: "CLOSED",
			groupRole: "ADMIN",
		},
	},
};

export const RequestedGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupDescription: "Group Description",
			isDefault: false,
			groupType: "REQUEST",
			groupRole: "REQUESTED",
		},
	},
};

export const InvitedGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupDescription: "Group Description",
			isDefault: false,
			groupType: "REQUEST",
			groupRole: "INVITED",
		},
	},
};
