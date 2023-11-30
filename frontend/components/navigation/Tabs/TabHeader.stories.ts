import type { Meta, StoryObj } from "@storybook/react";

import { TabHeader } from "@/components/navigation/Tabs/TabHeader";

const meta: Meta<typeof TabHeader> = {
	title: "Navigation/Tabs/TabHeader",
	component: TabHeader,
	tags: ["autodocs"],
	args: {
		id: "TabHeaderId",
	},
};

export default meta;
type Story = StoryObj<typeof TabHeader>;

export const Cards: Story = {
	args: {
		name: "Cards",
		style: "cards",
	},
};

export const Decks: Story = {
	args: {
		name: "Decks",
		style: "decks",
	},
};

export const Groups: Story = {
	args: {
		name: "Groups",
		style: "groups",
	},
};

export const Invitations: Story = {
	args: {
		name: "Invitations",
		style: "invitations",
	},
};

export const Settings: Story = {
	args: {
		name: "Settings",
		style: "settings",
	},
};

export const Statistics: Story = {
	args: {
		name: "Statistics",
		style: "statistics",
	},
};

export const User: Story = {
	args: {
		name: "User",
		style: "user",
	},
};
