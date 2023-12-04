import type { Meta, StoryObj } from "@storybook/react";

import { Statistic } from "@/components/statistics/Statistic";

const meta: Meta<typeof Statistic> = {
	title: "Statistics/Statistic",
	component: Statistic,
	tags: ["autodocs"],
	args: {
		id: "StatisticId",
	},
};

export default meta;
type Story = StoryObj<typeof Statistic>;

export const Up: Story = {
	args: {
		header: "Cards learned",
		value: "176",
		separator: "from",
		reference: "200",
		change: 12,
	},
};

export const UpRight: Story = {
	args: {
		header: "Cards learned",
		value: "176",
		separator: "from",
		reference: "200",
		change: 7,
	},
};

export const Neutral: Story = {
	args: {
		header: "Cards learned",
		value: "176",
		separator: "from",
		reference: "200",
		change: 0,
	},
};

export const DownRight: Story = {
	args: {
		header: "Cards learned",
		value: "176",
		separator: "from",
		reference: "200",
		change: -3,
	},
};

export const Down: Story = {
	args: {
		header: "Cards learned",
		value: "176",
		separator: "from",
		reference: "200",
		change: -17,
	},
};
