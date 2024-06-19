import type { Meta, StoryObj } from "@storybook/react";

import { SelectionField } from "@/components/input/SelectionField";

const meta: Meta<typeof SelectionField> = {
	title: "Input/SelectionField",
	component: SelectionField,
	args: {},
};

export default meta;
type Story = StoryObj<typeof SelectionField>;

export const Default: Story = {
	args: {
		label: "SelectionField",
		placeholder: "Select an option",
	},
};
