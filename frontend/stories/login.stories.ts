import type { Meta, StoryObj } from "@storybook/react";
import { userEvent, within } from "@storybook/testing-library";

import Page from "../pages/login";

const meta: Meta<typeof Page> = {
	title: "Pages/Access",
	component: Page,
	parameters: {
		layout: "fullscreen",
	},
};

export default meta;
type Story = StoryObj<typeof Page>;

export const Login: Story = {};

export const Register: Story = {
	play: async ({ canvasElement }) => {
		const canvas = within(canvasElement);
		const registerButton = await canvas.getByText("Sign up now!");
		await userEvent.click(registerButton);
	},
};
