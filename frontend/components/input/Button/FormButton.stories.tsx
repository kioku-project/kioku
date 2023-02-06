import React from 'react'

import { ComponentStory, ComponentMeta } from '@storybook/react'

import {FormButton} from './FormButton'

export default {
    title: 'Input/FormButton',
    component: FormButton
} as ComponentMeta<typeof FormButton>;

const Template: ComponentStory<typeof FormButton> = (args) => <FormButton {...args} />;

export const Default = Template.bind({});
Default.args = {
    label: 'FormButton'
}