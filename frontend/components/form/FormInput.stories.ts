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
		style: "primary",
		value: "Test",
	},
};

export const Placeholder: Story = {
	args: {
		placeholder: "Placeholder",
	},
};
