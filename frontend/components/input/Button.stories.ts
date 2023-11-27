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
		buttonStyle: "primary",
		buttonSize: "md",
	},
};

export const Secondary: Story = {
	args: {
		children: "Secondary",
		buttonStyle: "secondary",
		buttonSize: "md",
	},
};

export const Error: Story = {
	args: {
		children: "Error",
		buttonStyle: "error",
		buttonSize: "md",
	},
};

export const Warning: Story = {
	args: {
		children: "Warning",
		buttonStyle: "warning",
		buttonSize: "md",
	},
};

export const Disabled: Story = {
	args: {
		children: "Disabled",
		buttonStyle: "disabled",
		buttonSize: "md",
	},
};

export const Icon: Story = {
	args: {
		children: "Let's go",
		buttonStyle: "primary",
		buttonSize: "md",
		buttonIcon: "ArrowRight",
	},
};

export const Small: Story = {
	args: {
		children: "Button",
		buttonStyle: "primary",
		buttonSize: "sm",
	},
};

export const Medium: Story = {
	args: {
		children: "Button",
		buttonStyle: "primary",
		buttonSize: "md",
	},
};

export const Large: Story = {
	args: {
		children: "Button",
		buttonStyle: "primary",
		buttonSize: "lg",
	},
};
