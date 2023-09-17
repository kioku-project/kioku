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

export const Disabled: Story = {
	args: {
		value: "Disabled",
		style: "disabled",
	},
};

export const Small: Story = {
	args: {
		value: "Small",
		size: "sm",
	},
};

export const Medium: Story = {
	args: {
		value: "Medium",
		size: "md",
	},
};

export const Large: Story = {
	args: {
		value: "Large",
		size: "lg",
	},
};
