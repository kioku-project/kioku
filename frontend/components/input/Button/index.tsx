import React from 'react';

export interface ButtonProps {
	/**
	 * id of the button
	 */
	id?: string
	/**
	 * label of the button
	 */
	label?: string;
	/**
	  * background color of the button
	  */
	backgroundColor?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

export const Button = ({ id, label, backgroundColor, ...props }: ButtonProps) => {
	return (
		<button
			id={id}
			type="button"
			className={`p-3 rounded-md border border-black bg-white`}
			style={{ backgroundColor }}
			{...props}
		>
			{label}
		</button>
	)
}