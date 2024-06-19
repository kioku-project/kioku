import type { Meta, StoryObj } from "@storybook/react";

import { InputAction } from "@/components/input/InputAction";

const meta: Meta<typeof InputAction> = {
	title: "Input/InputAction",
	component: InputAction,
	args: {
		id: "inputActionId",
	},
};

export default meta;
type Story = StoryObj<typeof InputAction>;

export const Default: Story = {
	args: {
		header: "InputAction",
		value: "Input",
		button: "Action",
	},
};

export const Disabled: Story = {
	args: {
		header: "InputAction",
		value: "Input",
		button: "Action",
		disabled: true,
	},
};

export const WithoutValue: Story = {
	args: {
		header: "InputAction",
		button: "Action",
	},
};
