import type { Meta, StoryObj } from "@storybook/react";

import { Navbar } from "@/components/navigation/Navbar";

const meta: Meta<typeof Navbar> = {
	title: "Navigation/NavBar",
	component: Navbar,
	args: {},
};

export default meta;
type Story = StoryObj<typeof Navbar>;

export const Default: Story = {
	args: {},
};
