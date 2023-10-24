import type { Meta, StoryObj } from "@storybook/react";

import LoadingSpinner from "./LoadingSpinner";

const meta: Meta<typeof LoadingSpinner> = {
	title: "Graphics/LoadingSpinner",
	component: LoadingSpinner,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof LoadingSpinner>;

export const Default: Story = {
	args: {
		className: "w-16",
	},
};

export const NoDelay: Story = {
	args: {
		className: "w-16",
		delay: 0,
	},
};
