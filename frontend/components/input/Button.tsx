import Link, { LinkProps } from "next/link";
import React, { ButtonHTMLAttributes, ReactNode } from "react";

import { Text } from "@/components/Text";
import { Icon, IconName } from "@/components/graphics/Icon";
import { Size } from "@/types/Size";

export interface ButtonProps {
	/**
	 * If href is set, a Link will be returned
	 */
	href?: string;
	/**
	 * Button styling
	 */
	buttonStyle?: keyof typeof getStyle;
	/**
	 * Button size
	 */
	buttonSize?: string;
	/**
	 * Text size
	 */
	buttonTextSize?: Size;
	/**
	 * Icon that will be displayed in the Button
	 */
	buttonIcon?: IconName | ReactNode;
	/**
	 * Icon size
	 */
	buttonIconSize?: number;
}

const getStyle = {
	primary: "bg-kiokuDarkBlue font-medium text-white hover:scale-[1.02]",
	secondary:
		"bg-black font-medium text-white hover:scale-[1.02] hover:bg-neutral-900",
	tertiary:
		"bg-transparent font-medium text-kiokuDarkBlue hover:scale-105 hover:bg-gray-100",
	cancel: "bg-transparent font-normal text-gray-400 hover:bg-gray-100",
	error: "bg-kiokuRed font-medium text-white hover:scale-105",
	warning: "bg-kiokuYellow font-medium text-white hover:scale-105",
	disabled: "bg-gray-200 font-medium text-gray-400 hover:cursor-not-allowed",
} as const;

/**
 * UI component for user interactions
 */
export const Button = ({
	href,
	replace,
	scroll,
	buttonStyle,
	buttonSize = "px-3 py-1.5 lg:px-3 lg:py-2",
	buttonTextSize,
	buttonIcon,
	buttonIconSize = 16,
	className = "",
	children,
	...props
}: ButtonProps &
	ButtonHTMLAttributes<HTMLButtonElement> &
	Pick<LinkProps, "replace" | "scroll">) => {
	const innerButton = (
		<>
			<Text textSize={buttonTextSize}>{children}</Text>
			{buttonIcon &&
				(typeof buttonIcon === "string" ? (
					<Icon
						icon={buttonIcon as IconName}
						size={buttonIconSize}
						className="flex-none"
					/>
				) : (
					buttonIcon
				))}
		</>
	);
	const classNames = [
		"flex items-center space-x-1 rounded-md outline-none transition",
		className,
	];
	if (buttonStyle) {
		classNames.push(getStyle[buttonStyle]);
	}
	classNames.push(buttonSize);

	return href ? (
		<Link
			href={href}
			replace={replace}
			scroll={scroll}
			className={classNames.join(" ")}
		>
			{innerButton}
		</Link>
	) : (
		<button className={classNames.join(" ")} {...props}>
			{innerButton}
		</button>
	);
};
