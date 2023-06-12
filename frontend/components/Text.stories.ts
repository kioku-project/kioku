import type { Meta, StoryObj } from "@storybook/react";

import { Text } from "./Text";

const meta: Meta<typeof Text> = {
	title: "Components/Text",
	component: Text,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof Text>;

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

export const XS: Story = {
	args: {
		children: "extra small",
		size: "xs",
	},
};
export const SM: Story = {
	args: {
		children: "small",
		size: "sm",
	},
};
export const MD: Story = {
	args: {
		children: "medium",
		size: "md",
	},
};
export const LG: Story = {
	args: {
		children: "large",
		size: "lg",
	},
};
export const XL: Story = {
	args: {
		children: "extra large",
		size: "xl",
	},
};
