import type { Meta, StoryObj } from "@storybook/react";

import { TextArea } from "./TextArea";

const meta: Meta<typeof TextArea> = {
	title: "Form/TextArea",
	component: TextArea,
	tags: ["autodocs"],
	args: {
		id: "TextAreaId",
	},
};

export default meta;
type Story = StoryObj<typeof TextArea>;

export const Default: Story = {
	args: {
		name: "text",
		value: "Test",
	},
};

