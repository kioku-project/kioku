import type { Meta, StoryObj } from "@storybook/react";

import { DangerAction } from "./DangerAction";

const meta: Meta<typeof DangerAction> = {
	title: "Input/DangerAction",
	component: DangerAction,
	tags: ["autodocs"],
	args: {
		id: "dangerActionId",
	},
};

export default meta;
type Story = StoryObj<typeof DangerAction>;

export const Default: Story = {
	args: {
		header: "DangerAction",
		description: "Danger",
		button: "Action",
	},
};

export const Disabled: Story = {
	args: {
		header: "DangerAction",
		description: "Danger",
		button: "Action",
		disabled: true,
	},
};

export const Header: Story = {
	args: {
		header: "DangerAction",
		button: "Action",
	},
};

export const Description: Story = {
	args: {
		description: "Danger",
		button: "Action",
	},
};
