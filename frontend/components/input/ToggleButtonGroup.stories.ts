import type { Meta, StoryObj } from "@storybook/react";

import { ToggleButtonGroup } from "@/components/input/ToggleButtonGroup";

const meta: Meta<typeof ToggleButtonGroup> = {
	title: "Input/ToggleButtonGroup",
	component: ToggleButtonGroup,
	args: {
		id: "ButtonId",
		choices: ["ONE", "TWO", "THREE"],
	},
};

export default meta;
type Story = StoryObj<typeof ToggleButtonGroup>;

export const Primary: Story = {
	args: {
		activeButton: "ONE",
		activeButtonStyle: "primary",
	},
};

export const Warning: Story = {
	args: {
		activeButton: "TWO",
		activeButtonStyle: "warning",
	},
};

export const Error: Story = {
	args: {
		activeButton: "THREE",
		activeButtonStyle: "error",
	},
};

export const Disabled: Story = {
	args: {
		activeButton: "ONE",
		disabled: true,
	},
};
