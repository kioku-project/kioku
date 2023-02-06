import React from 'react';

import { ButtonProps } from '.';

export const FormButton = ({ id, label, backgroundColor, ...props }: ButtonProps) => {
	return (
		<input
			type="submit"
			id={id}
			value={label}
			className={`p-2 rounded-md bg-white hover:cursor-pointer transition ease-in-out delay-100 hover:-translate-y-0.5 hover:scale-105 hover:bg-gray-100 duration-200`}
			{...props}
		/>
	)
}