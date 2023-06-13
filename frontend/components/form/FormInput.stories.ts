import type { Meta, StoryObj } from "@storybook/react";

import { FormInput } from "./FormInput";

const meta: Meta<typeof FormInput> = {
	title: "Form/FormInput",
	component: FormInput,
	tags: ["autodocs"],
	args: {
		id: "InputId",
	},
};

export default meta;
type Story = StoryObj<typeof FormInput>;

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

export const Placeholder: Story = {
	args: {
		label: "Placeholder",
		placeholder: "Placeholder",
	},
};
