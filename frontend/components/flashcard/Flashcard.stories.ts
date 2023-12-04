import type { Meta, StoryObj } from "@storybook/react";

import { Flashcard } from "@/components/flashcard/Flashcard";

const meta: Meta<typeof Flashcard> = {
	title: "Components/Flashcard",
	component: Flashcard,
	tags: ["autodocs"],
	args: {
		card: {
			cardID: "C-12345678",
			sides: [
				{
					cardSideID: "S-12345678",
					header: "Front Header",
					description: "Front Description",
				},
				{
					cardSideID: "S-12345678",
					header: "Middle Header",
					description: "Middle Description",
				},
				{
					cardSideID: "S-12345678",
					header: "Back Header",
					description: "Back Description",
				},
			],
		},
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

export const Editable: Story = {
	args: {
		editable: true,
	},
};

export const EditFront: Story = {
	args: {
		isEdit: true,
		editable: true,
	},
};

export const EditMiddle: Story = {
	args: {
		cardSide: 1,
		isEdit: true,
		editable: true,
	},
};

export const EditBack: Story = {
	args: {
		cardSide: 2,
		isEdit: true,
		editable: true,
	},
};

export const FullSize: Story = {
	args: {
		fullSize: true,
	},
};
