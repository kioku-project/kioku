import type { Meta, StoryObj } from "@storybook/react";

import { Button } from "@/components/input/Button";

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
		buttonStyle: "primary",
	},
};

export const Secondary: Story = {
	args: {
		children: "Secondary",
		buttonStyle: "secondary",
	},
};

export const Tertiary: Story = {
	args: {
		children: "Tertiary",
		buttonStyle: "tertiary",
	},
};

export const Cancel: Story = {
	args: {
		children: "Cancel",
		buttonStyle: "cancel",
	},
};

export const Error: Story = {
	args: {
		children: "Error",
		buttonStyle: "error",
	},
};

export const Warning: Story = {
	args: {
		children: "Warning",
		buttonStyle: "warning",
	},
};

export const Disabled: Story = {
	args: {
		children: "Disabled",
		buttonStyle: "disabled",
	},
};

export const Icon: Story = {
	args: {
		children: "Let's go",
		buttonStyle: "primary",
		buttonIcon: "ArrowRight",
	},
};
