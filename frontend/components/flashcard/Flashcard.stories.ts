import type { Meta, StoryObj } from "@storybook/react";

import { Flashcard } from "./Flashcard";

const meta: Meta<typeof Flashcard> = {
	title: "Components/Flashcard",
	component: Flashcard,
	tags: ["autodocs"],
	args: {
		id: "CardId",
		sides: [
			{
				header: "Front Header",
				description: "Front Description",
			},
			{
				header: "Middle Header",
				description: "Middle Description",
			},
			{
				header: "Back Header",
				description: "Back Description",
			},
		],
	},
};

export default meta;
type Story = StoryObj<typeof Flashcard>;

export const Front: Story = {
	args: {},
};

export const Middle: Story = {
	args: {
		cardSide: 1,
	},
};

export const Back: Story = {
	args: {
		cardSide: 2,
	},
};

export const EditFront: Story = {
	args: {
		isEdit: true,
	},
};

export const EditMiddle: Story = {
	args: {
		cardSide: 1,
		isEdit: true,
	},
};

export const EditBack: Story = {
	args: {
		cardSide: 2,
		isEdit: true,
	},
};

export const FullSize: Story = {
	args: {
		fullSize: true,
	},
};
