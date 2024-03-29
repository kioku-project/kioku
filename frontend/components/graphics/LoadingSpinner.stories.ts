import type { Meta, StoryObj } from "@storybook/react";

import LoadingSpinner from "@/components/graphics/LoadingSpinner";

const meta: Meta<typeof LoadingSpinner> = {
	title: "Graphics/LoadingSpinner",
	component: LoadingSpinner,
	tags: ["autodocs"],
	args: {
		className: "w-16",
	},
};

export default meta;
type Story = StoryObj<typeof LoadingSpinner>;

export const Default: Story = {
	args: {},
};

export const NoDelay: Story = {
	args: {
		delay: 0,
	},
};

export const SimpleTheme: Story = {
	args: {
		delay: 0,
		theme: "simple",
	},
};
