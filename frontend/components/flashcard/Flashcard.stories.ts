import type { Meta, StoryObj } from "@storybook/react";

import { Card } from "./Flashcard";

const meta: Meta<typeof Card> = {
	title: "Components/Flashcard",
	component: Card,
	tags: ["autodocs"],
	args: {
		id: "CardId",
		card: [
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
		cardsleft: 16,
	},
};

export default meta;
type Story = StoryObj<typeof Card>;

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
