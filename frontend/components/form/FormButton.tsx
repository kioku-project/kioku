import React from "react";

interface FormButtonProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Button contents
	 */
	value: string;
	/**
	 * Button styling
	 */
	style?: string;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

function getStyle(style: string): string {
	const getStyle: { [style: string]: string } = {
		primary: "bg-darkblue border-darkblue text-eggshell",
	};
	return getStyle[style] ?? getStyle.primary;
}

/**
 * UI component for submitting forms
 */
export const FormButton = ({
	className,
	style = "",
	...props
}: FormButtonProps) => {
	return (
		<input
			type="submit"
			className={`flex justify-center rounded-md border-2 px-3 py-1.5 text-center text-sm font-semibold leading-6 shadow-sm outline-none transition hover:scale-105 hover:cursor-pointer ${getStyle(
				style
			)} ${className ?? ""}`}
			{...props}
		/>
	);
};
