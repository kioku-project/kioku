import React, { ReactNode } from "react";

interface ButtonProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Button contents
	 */
	children: ReactNode;
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
		primary: "bg-darkblue border-darkblue text-eggshell shadow-sm",
		secondary: "bg-transparent border-transparent text-darkblue",
	};
	return getStyle[style] ?? getStyle.primary;
}

/**
 * UI component for user interactions
 */
export const Button = ({
	className,
	style,
	children,
	...props
}: ButtonProps) => {
	return (
		<button
			className={`flex items-center justify-center rounded-md border-2 px-3 py-1.5 text-center text-xs font-bold leading-6 outline-none transition hover:scale-105 hover:cursor-pointer sm:text-sm md:text-base lg:px-5 lg:py-3 lg:text-lg xl:text-xl ${getStyle(
				style ?? ""
			)} ${className ?? ""}`}
			{...props}
		>
			{children}
		</button>
	);
};
