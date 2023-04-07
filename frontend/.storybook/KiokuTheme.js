import { create } from '@storybook/theming/create';

import image from '../public/kioku-logo-horizontal.svg'

export default create({
  base: 'light',
  brandTitle: 'Kioku',
  brandUrl: 'https://app.kioku.dev',
  brandImage: image,
  brandTarget: '_self',
});