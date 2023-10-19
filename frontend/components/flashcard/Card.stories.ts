import type { Meta, StoryObj } from "@storybook/react";
import { userEvent, within } from "@storybook/testing-library";

import { Card } from "./Card";

const meta: Meta<typeof Card> = {
	title: "Components/Card",
	component: Card,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof Card>;

export const Default: Story = {
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

export const Editable: Story = {
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
		editable: true
	},
};

export const Delete: Story = {
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
		editable: true
	},
	play: async ({ canvasElement }) => {
		const canvas = within(canvasElement);
		const deleteButton = await canvas.getByTestId("deleteCardButtonId");
		await userEvent.click(deleteButton);
	},
};

export const Placeholder: Story = {
	args: {
		card: {
			cardID: "",
			sides: [],
			deckID: "D-1234567",
		},
	},
};
