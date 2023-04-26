import React from "react";

interface FormInputProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * FormInput type
	 */
	type: string;
	/**
	 * FormInput name
	 */
	name: string;
	/**
	 * optional FormInput label
	 */
	label?: string;
	/**
	 * FormInput value
	 */
	value?: string;
	/**
	 * FormInput styling
	 */
	style?: string;
	/**
	 * Is the FormInput required?
	 */
	required?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
}

function getStyle(style: string): string {
	const getStyle: { [style: string]: string } = {
		primary:
			"border-eggshell bg-eggshell text-gray-900 invalid:border-red focus:border-darkblue",
	};
	return getStyle[style] ?? getStyle.primary;
}

/**
 * UI component for text inputs
 */
export const FormInput = ({
	name,
	label,
	required = true,
	style = "",
	className,
	...props
}: FormInputProps) => {
	return (
		<div className="flex w-full flex-col">
			<label
				htmlFor={name}
				className="block text-sm font-medium leading-6 text-gray-600"
			>
				{label}
			</label>
			<input
				name={name}
				className={`block w-full rounded-md border-2 px-1.5 py-1.5 outline-none sm:text-sm sm:leading-6 ${getStyle(
					style
				)}`}
				{...props}
			/>
		</div>
	);
};
