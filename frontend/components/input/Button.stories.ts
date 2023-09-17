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

export const Error: Story = {
	args: {
		children: "Error",
		style: "error",
	},
};

export const Warning: Story = {
	args: {
		children: "Warning",
		style: "warning",
	},
};

export const Disabled: Story = {
	args: {
		children: "Disabled",
		style: "disabled",
	},
};

export const Small: Story = {
	args: {
		children: "Small",
		size: "sm",
	},
};

export const Medium: Story = {
	args: {
		children: "Medium",
		size: "md",
	},
};

export const Large: Story = {
	args: {
		children: "Large",
		size: "lg",
	},
};
