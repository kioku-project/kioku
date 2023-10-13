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
export const XXXXXS: Story = {
	args: {
		children: "extra extreme extreme small",
		size: "5xs",
	},
};
export const XXXXS: Story = {
	args: {
		children: "extreme extreme small",
		size: "4xs",
	},
};
export const XXXS: Story = {
	args: {
		children: "extra extreme small",
		size: "3xs",
	},
};
export const XXS: Story = {
	args: {
		children: "extreme small",
		size: "2xs",
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
export const XXL: Story = {
	args: {
		children: "extreme large",
		size: "2xl",
	},
};
export const XXXL: Story = {
	args: {
		children: "extra extreme large",
		size: "3xl",
	},
};
export const XXXXL: Story = {
	args: {
		children: "extreme extreme large",
		size: "4xl",
	},
};
export const XXXXXL: Story = {
	args: {
		children: "extra extreme extreme large",
		size: "5xl",
	},
};
