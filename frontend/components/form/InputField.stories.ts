import type { Meta, StoryObj } from "@storybook/react";

import { InputField } from "./InputField";

const meta: Meta<typeof InputField> = {
	title: "Form/InputField",
	component: InputField,
	tags: ["autodocs"],
	args: {
		id: "InputFieldId",
		tooltipMessage: "This is a tooltip!",
	},
};

export default meta;
type Story = StoryObj<typeof InputField>;

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

export const XS: Story = {
	args: {
		label: "extra small",
		value: "Test",
		size: "xs",
	},
};

export const SM: Story = {
	args: {
		label: "small",
		value: "Test",
		size: "sm",
	},
};

export const MD: Story = {
	args: {
		label: "medium",
		value: "Test",
		size: "md",
	},
};

export const LG: Story = {
	args: {
		label: "large",
		value: "Test",
		size: "lg",
	},
};

export const XL: Story = {
	args: {
		label: "extra large",
		value: "Test",
		size: "xl",
	},
};

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

export const Success: Story = {
	args: {
		label: "Success",
		value: "Test",
		statusIcon: "success",
	},
};

export const Error: Story = {
	args: {
		label: "Error",
		value: "Test",
		statusIcon: "error",
	},
};

export const Warning: Story = {
	args: {
		label: "Warning",
		value: "Test",
		statusIcon: "warning",
	},
};

export const Info: Story = {
	args: {
		label: "Info",
		value: "Test",
		statusIcon: "info",
	},
};

export const Placeholder: Story = {
	args: {
		label: "Placeholder",
		placeholder: "Placeholder",
	},
};
