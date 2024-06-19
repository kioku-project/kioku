import type { Meta, StoryObj } from "@storybook/react";

import { TabHeader } from "@/components/navigation/Tabs/TabHeader";

const meta: Meta<typeof TabHeader> = {
	title: "Navigation/Tabs/TabHeader",
	component: TabHeader,
	args: {
		id: "TabHeaderId",
	},
};

export default meta;
type Story = StoryObj<typeof TabHeader>;

export const Cards: Story = {
	args: {
		name: "Dashboard",
		icon: "Home",
	},
};
