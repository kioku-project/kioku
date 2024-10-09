import type { Meta, StoryObj } from "@storybook/react";

import { Text } from "@/components/Text";

const meta: Meta<typeof Text> = {
	title: "Components/Text",
	component: Text,
	args: {},
};

export default meta;
type Story = StoryObj<typeof Text>;

export const Primary: Story = {
	args: {
		children: "Primary",
		textStyle: "primary",
	},
};

export const Secondary: Story = {
	args: {
		children: "Secondary",
		textStyle: "secondary",
	},
};
export const XXXXXS: Story = {
	args: {
		children: "extra extreme extreme small",
		textSize: "5xs",
	},
};
export const XXXXS: Story = {
	args: {
		children: "extreme extreme small",
		textSize: "4xs",
	},
};
export const XXXS: Story = {
	args: {
		children: "extra extreme small",
		textSize: "3xs",
	},
};
export const XXS: Story = {
	args: {
		children: "extreme small",
		textSize: "2xs",
	},
};
export const XS: Story = {
	args: {
		children: "extra small",
		textSize: "xs",
	},
};
export const SM: Story = {
	args: {
		children: "small",
		textSize: "sm",
	},
};
export const MD: Story = {
	args: {
		children: "medium",
		textSize: "md",
	},
};
export const LG: Story = {
	args: {
		children: "large",
		textSize: "lg",
	},
};
export const XL: Story = {
	args: {
		children: "extra large",
		textSize: "xl",
	},
};
export const XXL: Story = {
	args: {
		children: "extreme large",
		textSize: "2xl",
	},
};
export const XXXL: Story = {
	args: {
		children: "extra extreme large",
		textSize: "3xl",
	},
};
export const XXXXL: Story = {
	args: {
		children: "extreme extreme large",
		textSize: "4xl",
	},
};
export const XXXXXL: Story = {
	args: {
		children: "extra extreme extreme large",
		textSize: "5xl",
	},
};
