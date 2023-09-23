import type { Meta, StoryObj } from "@storybook/react";
import { userEvent, within } from "@storybook/testing-library";

import Member from "./Member";

const meta: Meta<typeof Member> = {
	title: "Components/Member",
	component: Member,
	tags: ["autodocs"],
	args: {},
};

export default meta;
type Story = StoryObj<typeof Member>;

export const Read: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
			groupID: "G-12345678",
			groupRole: "READ",
		},
	},
};

export const Write: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
			groupID: "G-12345678",
			groupRole: "WRITE",
		},
	},
};

export const Admin: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
			groupID: "G-12345678",
			groupRole: "ADMIN",
		},
	},
};

export const Delete: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
			groupID: "G-12345678",
			groupRole: "READ",
		},
	},
	play: async ({ canvasElement }) => {
		const canvas = within(canvasElement);
		const deleteButton = await canvas.getByTestId("deleteUserButtonId");
		await userEvent.click(deleteButton);
	},
};

export const Requested: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
			groupID: "G-12345678",
			groupRole: "REQUESTED",
		},
	},
};

export const Invited: Story = {
	args: {
		user: {
			userID: "U-12345678",
			userName: "Test User",
			groupID: "G-12345678",
			groupRole: "INVITED",
		},
	},
};

export const Placeholder: Story = {
	args: {
		user: {
			userID: "",
			userName: "",
			groupID: "G-12345678",
		},
	},
};
