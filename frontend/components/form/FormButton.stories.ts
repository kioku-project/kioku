import type { Meta, StoryObj } from '@storybook/react';

import { FormButton } from "./FormButton";

const meta: Meta<typeof FormButton> = {
    title: 'Form/FormButton',
    component: FormButton,
    tags: ['autodocs'],
    argTypes: {
     
    },
  };
  
  export default meta;
  type Story = StoryObj<typeof FormButton>;
  
  export const Default: Story = {
    args: {
      id: 'ButtonId',
      value: 'Button',
    },
  };  