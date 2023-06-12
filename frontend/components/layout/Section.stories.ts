import type { Meta, StoryObj } from "@storybook/react";

import { Section } from "./Section";

const meta: Meta<typeof Section> = {
	title: "Layout/Section",
	component: Section,
	tags: ["autodocs"],
	args: {
		id: "sectionId",
	},
};

export default meta;
type Story = StoryObj<typeof Section>;

export const Primary: Story = {
	args: {
		header: "Primary",
		style: "primary",
	},
};

export const Secondary: Story = {
	args: {
		header: "Secondary",
		style: "secondary",
	},
};

export const Error: Story = {
	args: {
		header: "Error",
		style: "error",
	},
};

export const NoBorder: Story = {
	args: {
		header: "No Border",
		style: "noBorder",
	},
};
