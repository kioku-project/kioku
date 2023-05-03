import type { Meta, StoryObj } from "@storybook/react";

import { Logo } from "./Logo";

const meta: Meta<typeof Logo> = {
	title: "Graphics/Logo",
	component: Logo,
	tags: ["autodocs"],
	args: {},
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
