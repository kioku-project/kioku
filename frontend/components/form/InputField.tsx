import React, { ChangeEventHandler } from "react";

interface InputFieldProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * InputField type
	 */
	type: string;
	/**
	 * InputField name
	 */
	name: string;
	/**
	 * optional InputField label
	 */
	label?: string;
	/**
	 * InputField value
	 */
	value?: string;
	/**
	 * InputField placeholder
	 */
	placeholder?: string;
	/**
	 * InputField styling
	 */
	style?: "primary" | "secondary" | "tertiary";
	/**
	 * Is the InputField required?
	 */
	required?: boolean;
	/**
	 * Is the InputField read only?
	 */
	readOnly?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional change handler
	 */
	onChange?: ChangeEventHandler<HTMLInputElement>;
}

function getLabelStyle(style: string): string {
	const getStyle: { [style: string]: string } = {
		primary: "text-kiokuDarkBlue",
		secondary: "text-kiokuDarkBlue font-bold",
		tertiary: "text-kiokuDarkBlue font-bold",
	};
	return getStyle[style] ?? "";
}

function getInputStyle(style: string): string {
	const getStyle: { [style: string]: string } = {
		primary:
			"border-2 border-eggshell bg-eggshell invalid:border-red px-1.5 py-1.5 font-medium text-kiokuDarkBlue focus:border-kiokuDarkBlue",
		secondary: "text-kiokuDarkBlue font-medium bg-transparent",
		tertiary: "text-kiokuLightBlue font-medium bg-transparent",
	};
	return getStyle[style] ?? "";
}

/**
 * UI component for text inputs
 */
export const InputField = ({
	name,
	label,
	required = true,
	style,
	className,
	...props
}: InputFieldProps) => {
	return (
		<div className={`flex w-full flex-col ${className ?? ""}`}>
			<label htmlFor={name} className={`${getLabelStyle(style!)}`}>
				{label}
			</label>
			<input
				name={name}
				className={`w-full rounded-md outline-none ${getInputStyle(
					style!
				)}`}
				{...props}
			/>
		</div>
	);
};
