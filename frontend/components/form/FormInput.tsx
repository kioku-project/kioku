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
	style?: string;
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

function getStyle(style: string): string {
	const getStyle: { [style: string]: string } = {
		primary:
			"invalid:border-red px-1.5 py-1.5 font-medium text-kiokuDarkBlue focus:border-kiokuDarkBlue",
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
	style = "",
	className,
	...props
}: FormInputProps) => {
	return (
		<div className="flex w-full flex-col text-kiokuDarkBlue">
			<label htmlFor={name} className="">
				{label}
			</label>
			<input
				name={name}
				className={`w-full rounded-md border-2 border-eggshell bg-eggshell outline-none ${getStyle(
					style
				)} ${className ?? ""}`}
				{...props}
			/>
		</div>
	);
};
