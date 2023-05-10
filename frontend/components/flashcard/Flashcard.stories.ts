import type { Meta, StoryObj } from "@storybook/react";

import { Card } from "./Flashcard";

const meta: Meta<typeof Card> = {
	title: "Components/Flashcard",
	component: Card,
	tags: ["autodocs"],
	args: {
		id: "CardId",
		card: {
			front: {
				header: "Front Header",
				description: "Front Description",
			},
			back: {
				header: "Back Header",
				description: "Back Description",
			},
		},
		cardsleft: 16,
	},
};

export default meta;
type Story = StoryObj<typeof Card>;

export const Front: Story = {
	args: {},
};

export const Back: Story = {
	args: {
		turned: true,
	},
};

export const EditFront: Story = {
	args: {
		edit: true,
	},
};

export const EditBack: Story = {
	args: {
		turned: true,
		edit: true,
	},
};
