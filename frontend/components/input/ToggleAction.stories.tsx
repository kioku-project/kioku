import type { Meta, StoryObj } from "@storybook/react";

import { ToggleAction } from "@/components/input/ToggleAction";

const meta: Meta<typeof ToggleAction> = {
	title: "Input/ToggleAction",
	component: ToggleAction,
	args: {
		id: "toggleActionId",
		choices: ["ONE", "TWO", "THREE"],
	},
};

export default meta;
type Story = StoryObj<typeof ToggleAction>;

export const Default: Story = {
	args: {
		header: "ToggleAction",
		description: "Choose one option",
		activeButton: "ONE",
	},
};

export const Disabled: Story = {
	args: {
		header: "ToggleAction",
		description: "This one is disabled",
		disabled: true,
		activeButton: "TWO",
	},
};

export const Header: Story = {
	args: {
		header: "ToggleAction",
		activeButton: "THREE",
	},
};

export const Description: Story = {
	args: {
		description: "Choose one option",
	},
};
