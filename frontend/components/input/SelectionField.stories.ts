import type { Meta, StoryObj } from "@storybook/react";

import { SelectionField } from "@/components/input/SelectionField";

const meta: Meta<typeof SelectionField> = {
	title: "Input/SelectionField",
	component: SelectionField,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof SelectionField>;

export const Default: Story = {
	args: {
		list: [
			{
				title: "ONE",
				description: "Description 1",
				isSelected: true,
				icon: "Square",
			},
			{
				title: "TWO",
				description: "Description 2",
				isSelected: false,
				icon: "Circle",
			},
		],
	},
};
export const NoSelected: Story = {
	args: {
		list: [
			{
				title: "ONE",
				description: "Description 1",
				isSelected: false,
				icon: "Square",
			},
			{
				title: "TWO",
				description: "Description 2",
				isSelected: false,
				icon: "Circle",
			},
		],
	},
};
