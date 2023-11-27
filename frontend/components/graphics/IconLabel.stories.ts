import type { Meta, StoryObj } from "@storybook/react";

import { IconLabel } from "./IconLabel";

const meta: Meta<typeof IconLabel> = {
	title: "Graphics/IconLabel",
	component: IconLabel,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof IconLabel>;

export const Default: Story = {
	args: {
		iconLabel: { icon: "AlertCircle", header: "Header" },
	},
};

export const Description: Story = {
	args: {
		iconLabel: {
			icon: "AlertCircle",
			header: "Header",
			description: "description",
		},
	},
};

export const Colored: Story = {
	args: {
		iconLabel: {
			icon: "AlertCircle",
			header: "Header",
			description: "description",
		},
		className: "text-kiokuRed",
	},
};

export const DoubleColored: Story = {
	args: {
		iconLabel: {
			icon: "AlertCircle",
			header: "Header",
			description: "description",
		},
		className: "text-kiokuDarkBlue",
		color: "text-kiokuRed",
	},
};
