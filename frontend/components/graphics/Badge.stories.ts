import type { Meta, StoryObj } from "@storybook/react";

import { Badge } from "@/components/graphics/Badge";

const meta: Meta<typeof Badge> = {
	title: "Components/Badge",
	component: Badge,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof Badge>;

export const Primary: Story = {
	args: {
		label: "Primary",
		style: "primary",
	},
};

export const Secondary: Story = {
	args: {
		label: "Secondary",
		style: "secondary",
	},
};

export const Tertiary: Story = {
	args: {
		label: "Tertiary",
		style: "tertiary",
	},
};
