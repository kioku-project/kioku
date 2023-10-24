import type { Meta, StoryObj } from "@storybook/react";

import { Header } from "./Header";

const meta: Meta<typeof Header> = {
	title: "Layout/Header",
	component: Header,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof Header>;

export const Deck: Story = {
	args: {
		deck: {
			deckID: "D-12345678",
			deckName: "Test Deck",
		},
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
		},
	},
};

export const Group: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
		},
	},
};

export const GroupWithDescription: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupDescription: "Group Description",
		},
	},
};

export const PrivateGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupType: "PRIVATE",
		},
	},
};

export const PublicGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupType: "PUBLIC",
		},
	},
};

export const RequestedGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupRole: "REQUESTED",
		},
	},
};

export const InvitedGroup: Story = {
	args: {
		group: {
			groupID: "G-12345678",
			groupName: "Test Group",
			groupRole: "INVITED",
		},
	},
};

export const User: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
		},
	},
};
