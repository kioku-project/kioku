import type { Meta, StoryObj } from "@storybook/react";

import { InputField } from "./InputField";

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

export const TextInput: Story = {
	args: {
		type: "text",
		name: "text",
		label: "Text Input",
		value: "Test",
	},
};

export const EmailInput: Story = {
	args: {
		type: "email",
		name: "email",
		label: "Email Input",
		value: "test@example.com",
	},
};

export const PasswordInput: Story = {
	args: {
		type: "password",
		name: "password",
		label: "Password Input",
		value: "superSecret!",
	},
};

export const Primary: Story = {
	args: {
		label: "Primary",
		value: "Test",
		style: "primary",
	},
};

export const Secondary: Story = {
	args: {
		label: "Secondary",
		value: "Test",
		style: "secondary",
	},
};

export const Tertiary: Story = {
	args: {
		label: "Tertiary",
		value: "Test",
		style: "tertiary",
	},
};

export const Placeholder: Story = {
	args: {
		label: "Placeholder",
		placeholder: "Placeholder",
	},
};
