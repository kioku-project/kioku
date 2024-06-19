import type { Meta, StoryObj } from "@storybook/react";
import { userEvent, within } from "@storybook/test";

import { Card } from "@/components/flashcard/Card";

const meta: Meta<typeof Card> = {
	title: "Components/Card",
	component: Card,
	args: {
		card: {
			cardID: "C-12345678",
			sides: [
				{
					cardSideID: "S-12345678",
					header: "Header",
					description: "Description",
				},
			],
			deckID: "D-1234567",
		},
	},
};

export default meta;
type Story = StoryObj<typeof Card>;

export const Default: Story = {};

export const Editable: Story = {
	args: {
		editable: true,
	},
};

export const Delete: Story = {
	args: {
		editable: true,
	},
	play: async ({ canvasElement }) => {
		const canvas = within(canvasElement);
		const deleteButton = canvas.getByTestId("deleteCardButtonId");
		await userEvent.click(deleteButton);
	},
};
