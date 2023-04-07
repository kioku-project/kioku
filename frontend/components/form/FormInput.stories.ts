import type { Meta, StoryObj } from '@storybook/react';

import { FormInput } from "./FormInput";

const meta: Meta<typeof FormInput> = {
    title: 'Form/FormInput',
    component: FormInput,
    tags: ['autodocs'],
    argTypes: {
        id: 'InputId'
    },
};

export default meta;
type Story = StoryObj<typeof FormInput>;

export const TextInput: Story = {
    args: {
        type: 'text',
        name: 'text',
        label: 'Text Input',
    },
};  

export const EmailInput: Story = {
    args: {
        type: 'email',
        name: 'email',
        label: 'Email Input',
    },
};  

export const PasswordInput: Story = {
    args: {
        type: 'password',
        name: 'password',
        label: 'Password Input',
    },
};  