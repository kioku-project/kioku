import React from "react";

interface ButtonProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Button contents
	 */
	value?: string;
	/**
	 * Button styling
	 */
	style?: "primary" | "secondary";
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
		primary:
			"bg-darkblue border-darkblue text-eggshell hover:bg-lightblue shadow-sm",
		secondary:
			"bg-transparent border-transparent hover:bg-lightblue text-darkblue",
	};
	return getStyle[style] ?? getStyle.primary;
}

/**
 * UI component for user interactions
 */
export const Button = ({ className, value, style, ...props }: ButtonProps) => {
	return (
		<button
			className={`flex justify-center rounded-md border-2 px-3 py-1.5 text-center text-xs font-bold leading-6 outline-none transition hover:cursor-pointer sm:text-sm md:text-base lg:px-5 lg:py-3 lg:text-lg xl:text-xl ${getStyle(
				style || ""
			)} ${className || ""}`}
			{...props}
		>
			{value}
		</button>
	);
};
