import type { Meta, StoryObj } from "@storybook/react";

import { Button } from "./Button";

const meta: Meta<typeof Button> = {
	title: "Input/Button",
	component: Button,
	tags: ["autodocs"],
	args: {
		id: "ButtonId",
	},
};

export default meta;
type Story = StoryObj<typeof Button>;

export const Primary: Story = {
	args: {
		children: "Primary",
		style: "primary",
	},
};

export const Secondary: Story = {
	args: {
		children: "Secondary",
		style: "secondary",
	},
};
