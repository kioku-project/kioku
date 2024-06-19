import type { Meta, StoryObj } from "@storybook/react";

import { Logo } from "@/components/graphics/Logo";

const meta: Meta<typeof Logo> = {
	title: "Graphics/Logo",
	component: Logo,
	args: {
		href: "/",
	},
};

export default meta;
type Story = StoryObj<typeof Logo>;

export const WithText: Story = {
	args: {
		text: true,
	},
};

export const NoText: Story = {
	args: {
		text: false,
	},
};
