import type { Meta, StoryObj } from "@storybook/react";
import GroupAsideTile from "./GroupAsideTile";

const meta: Meta<typeof GroupAsideTile> = {
	title: "Group/AsideTile",
	component: GroupAsideTile,
	tags: ["autodocs"],
	args: {
		id: "GroupId",
	},
};

export default meta;
type Story = StoryObj<typeof GroupAsideTile>;

export const Default: Story = {
	args: {
		name: "Example Group",
		count: 2,
	},
};

export const NoCards: Story = {
	args: {
		name: "Example Group",
		count: 0,
	},
};
