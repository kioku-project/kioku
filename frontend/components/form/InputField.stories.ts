import type { Meta, StoryObj } from "@storybook/react";
import { userEvent, within } from "@storybook/testing-library";

import { InputField } from "@/components/form/InputField";

const meta: Meta<typeof InputField> = {
	title: "Form/InputField",
	component: InputField,
	tags: ["autodocs"],
	args: {
		id: "InputFieldId",
	},
};

export default meta;
type Story = StoryObj<typeof InputField>;

export const Primary: Story = {
	args: {
		label: "Primary",
		value: "Test",
		inputFieldStyle: "primary",
	},
};

export const Secondary: Story = {
	args: {
		label: "Secondary",
		value: "Test",
		inputFieldStyle: "secondary",
	},
};

export const XS: Story = {
	args: {
		label: "extra small",
		value: "Test",
		inputFieldSize: "xs",
	},
};

export const SM: Story = {
	args: {
		label: "small",
		value: "Test",
		inputFieldSize: "sm",
	},
};

export const MD: Story = {
	args: {
		label: "medium",
		value: "Test",
		inputFieldSize: "md",
	},
};

export const LG: Story = {
	args: {
		label: "large",
		value: "Test",
		inputFieldSize: "lg",
	},
};

export const XL: Story = {
	args: {
		label: "extra large",
		value: "Test",
		inputFieldSize: "xl",
	},
};

export const Text: Story = {
	args: {
		type: "text",
		value: "Test",
	},
};

export const Email: Story = {
	args: {
		type: "email",
		value: "test@example.com",
	},
};

export const Password: Story = {
	args: {
		type: "password",
		value: "superSecret!",
	},
};

export const PasswordShown: Story = {
	args: {
		type: "password",
		value: "superSecret!",
	},
	play: async ({ canvasElement }) => {
		const canvas = within(canvasElement);
		const icon = await canvas.getByTestId("inputFieldIconId");
		await userEvent.click(icon);
	},
};

export const Icon: Story = {
	args: {
		value: "Test",
		inputFieldIcon: "Check",
		tooltip: "This is a tooltip!",
	},
};

export const Placeholder: Story = {
	args: {
		placeholder: "Placeholder",
	},
};
