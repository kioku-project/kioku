import React, { ChangeEventHandler } from "react";
import { EventsType } from "react-tooltip";

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
	 * FormInput placeholder
	 */
	placeholder?: string;
	/**
	 * FormInput styling
	 */
	style?: "primary" | "secondary";
	/**
	 * Is the FormInput required?
	 */
	required?: boolean;
	/**
	 * Is the FormInput read only?
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
	};
	return getStyle[style] ?? "";
}

function getInputStyle(style: string): string {
	const getStyle: { [style: string]: string } = {
		primary:
			"border-2 border-eggshell bg-eggshell invalid:border-red px-1.5 py-1.5 font-medium text-kiokuDarkBlue focus:border-kiokuDarkBlue",
		secondary: "text-kiokuLightBlue font-medium bg-transparent",
	};
	return getStyle[style] ?? "";
}

/**
 * UI component for text inputs
 */
export const FormInput = ({
	name,
	label,
	required = true,
	style,
	className,
	...props
}: FormInputProps) => {
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
