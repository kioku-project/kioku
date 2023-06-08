import type { Meta, StoryObj } from "@storybook/react";

import { FormButton } from "./FormButton";

const meta: Meta<typeof FormButton> = {
	title: "Form/FormButton",
	component: FormButton,
	tags: ["autodocs"],
	args: {
		id: "ButtonId",
	},
};

export default meta;
type Story = StoryObj<typeof FormButton>;

export const Primary: Story = {
	args: {
		value: "Primary",
		style: "primary",
	},
};

export const Small: Story = {
	args: {
		value: "Small",
		size: "small",
	},
};

export const Medium: Story = {
	args: {
		value: "Medium",
		size: "medium",
	},
};

export const Large: Story = {
	args: {
		value: "Large",
		size: "large",
	},
};
