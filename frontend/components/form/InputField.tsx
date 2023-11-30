import React, {
	InputHTMLAttributes,
	MouseEventHandler,
	Ref,
	forwardRef,
	useState,
} from "react";
import { Tooltip } from "react-tooltip";

import { Text } from "@/components/Text";
import { Icon, IconName } from "@/components/graphics/Icon";

interface InputFieldProps {
	/**
	 * InputField label that will be displayed above the InputField
	 */
	label?: string;
	/**
	 * InputField styling
	 */
	inputFieldStyle?: keyof typeof getStyle;
	/**
	 * InputField size
	 */
	inputFieldSize?: keyof typeof getSize;
	/**
	 * Icon that will be displayed on the right side of the InputField.
	 */
	inputFieldIcon?: IconName;
	/**
	 * Icon style
	 */
	inputFieldIconStyle?: string;
	/**
	 * Icon size
	 */
	inputFieldIconSize?: number;
	/**
	 * Message that will be displayed as tooltip on the icon
	 */
	tooltip?: string;
	/**
	 * Icon onClick handler
	 */
	onClickIcon?: MouseEventHandler;
}

const getStyle = {
	primary: "border-none text-kiokuDarkBlue",
	secondary: "border-none text-kiokuLightBlue",
} as const;

const getSize = {
	"5xs": "text-xs sm:text-xs md:text-xs lg:text-xs xl:text-xs",
	"4xs": "text-xs sm:text-xs md:text-xs lg:text-xs xl:text-sm",
	"3xs": "text-xs sm:text-xs md:text-xs lg:text-sm xl:text-base",
	"2xs": "text-xs sm:text-xs md:text-sm lg:text-base xl:text-lg",
	xs: "text-xs sm:text-sm md:text-base lg:text-lg xl:text-xl",
	sm: "text-sm sm:text-base md:text-lg lg:text-xl xl:text-2xl",
	md: "text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl",
	lg: "text-lg sm:text-xl md:text-2xl lg:text-3xl xl:text-4xl",
	xl: "text-xl sm:text-2xl md:text-3xl lg:text-4xl xl:text-5xl",
	"2xl": "text-2xl sm:text-3xl md:text-4xl lg:text-5xl xl:text-6xl",
	"3xl": "text-3xl sm:text-4xl md:text-5xl lg:text-6xl xl:text-7xl",
	"4xl": "text-4xl sm:text-5xl md:text-6xl lg:text-7xl xl:text-8xl",
	"5xl": "text-5xl sm:text-6xl md:text-7xl lg:text-8xl xl:text-9xl",
} as const;

/**
 * UI component for text inputs
 */
export const InputField = forwardRef(
	(
		{
			id = "inputFieldId",
			type,
			label,
			inputFieldStyle,
			inputFieldSize,
			inputFieldIcon,
			inputFieldIconStyle,
			inputFieldIconSize = 16,
			tooltip,
			className = "",
			onClickIcon,
			...props
		}: InputFieldProps & InputHTMLAttributes<HTMLInputElement>,
		ref: Ref<HTMLInputElement>
	) => {
		const [inputType, setInputType] = useState(type);
		const [statusIcon, setStatusIcon] = useState(
			type === "password" && !inputFieldIcon ? "EyeOff" : inputFieldIcon
		);

		return (
			<div className={`flex w-full flex-col rounded-md ${className}`}>
				<label htmlFor={id}>
					<Text
						textStyle="primary"
						textSize={inputFieldSize}
						className="font-bold"
					>
						{label}
					</Text>
				</label>
				<div
					className={`flex flex-row items-center rounded-md ${
						inputFieldStyle ? getStyle[inputFieldStyle] : ""
					} `}
				>
					<input
						id={id}
						type={inputType}
						className={`w-full border-none bg-transparent font-medium outline-none ${
							inputFieldSize ? getSize[inputFieldSize] : ""
						}`}
						ref={ref}
						{...props}
					/>
					{statusIcon && (
						<Icon
							icon={statusIcon}
							size={inputFieldIconSize}
							className={`${inputFieldIconStyle} ${
								onClickIcon ? "hover:cursor-pointer" : ""
							}`}
							data-tooltip-id={`tooltip-${id}`}
							data-testid={`inputFieldIconId`}
							onClick={(event) => {
								if (type === "password") {
									setInputType((prev) =>
										prev === "password"
											? "text"
											: "password"
									);
									setStatusIcon((prev) => {
										if (prev === "EyeOff") {
											return "Eye";
										}
										if (prev === "Eye") {
											return "EyeOff";
										}
									});
								}
								onClickIcon?.(event);
							}}
							onKeyUp={(event) => {
								if (event.key === "Enter") {
									event.target.dispatchEvent(
										new Event("click", { bubbles: true })
									);
								}
							}}
						/>
					)}
					<Tooltip id={`tooltip-${id}`} content={tooltip} />
				</div>
			</div>
		);
	}
);

InputField.displayName = "InputField";
