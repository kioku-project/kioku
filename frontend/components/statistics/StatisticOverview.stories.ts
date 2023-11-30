import type { Meta, StoryObj } from "@storybook/react";

import { StatisticOverview } from "@/components/statistics/StatisticOverview";

const meta: Meta<typeof StatisticOverview> = {
	title: "Statistics/StatisticOverview",
	component: StatisticOverview,
	tags: ["autodocs"],
	args: {
		id: "StatisticId",
	},
};

export default meta;
type Story = StoryObj<typeof StatisticOverview>;

export const Default: Story = {
	args: {},
};
