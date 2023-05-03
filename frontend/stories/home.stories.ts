import type { Meta, StoryObj } from '@storybook/react';

import Page from '../pages/home';

const meta: Meta<typeof Page> = {
  title: 'Pages/Home',
  component: Page,
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj<typeof Page>;

export const Default: Story = {};
