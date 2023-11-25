import { t } from "@lingui/macro";
import React, {
	InputHTMLAttributes,
	ReactNode,
	Ref,
	forwardRef,
	useEffect,
	useState,
} from "react";
import { AlertCircle, AlertTriangle, Check, Info } from "react-feather";
import { Tooltip } from "react-tooltip";

import { Text } from "../Text";

interface InputFieldProps {
	/**
	 * InputField label that will be displayed above the InputField
	 */
	label?: string;
	/**
	 * Icon that will be displayed on the right side of the InputField. If Icon is undefined, it will be set dynamically according to InputField validity.
	 */
	statusIcon?: keyof typeof statusIcon;
	/**
	 * Message that will be displayed as tooltip on the icon
	 */
	tooltipMessage?: string;
	/**
	 * InputField styling
	 */
	inputFieldStyle?: keyof typeof getStyle;
	/**
	 * InputField size
	 */
	inputFieldSize?: keyof typeof getSize;
}

const getStyle = {
	primary: "border-2 py-1 pl-1.5 pr-1 sm:py-1 sm:pr-3",
	secondary: "border-none text-kiokuDarkBlue",
	tertiary: "border-none text-kiokuLightBlue",
};

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

const statusIcon = {
	none: <></>,
	success: <Check className="text-kiokuDarkBlue outline-none" />,
	error: <AlertCircle className="text-kiokuRed outline-none" />,
	warning: <AlertTriangle className="text-kiokuYellow outline-none" />,
	info: <Info className="text-kiokuDarkBlue outline-none" />,
};
function getIcon(status: keyof typeof statusIcon, id: string): ReactNode {
	return <span data-tooltip-id={`tooltip-${id}`}>{statusIcon[status]}</span>;
}

/**
 * UI component for text inputs
 */
export const InputField = forwardRef(
	(
		{
			id = "inputFieldId",
			type,
			label,
			statusIcon,
			tooltipMessage = t`Please enter a valid ${type}.`,
			inputFieldStyle = "primary",
			inputFieldSize = "md",
			pattern,
			className = "",
			onChange = () => {},
			onBlur = () => {},
			...props
		}: InputFieldProps & InputHTMLAttributes<HTMLInputElement>,
		ref: Ref<HTMLInputElement>
	) => {
		const [initialised, setInitialised] = useState(false);
		const [icon, setIcon] = useState(statusIcon);
		const [tooltip, setTooltip] = useState("");
		const [inputPattern, setInputPattern] = useState(pattern);

		const isValid =
			typeof ref !== "function" && ref?.current?.validity.valid;

		useEffect(() => {
			setInputPattern(pattern);
			if (typeof ref !== "function" && initialised && ref?.current) {
				checkValidity(ref.current);
			}
		});

		return (
			<div
				className={`flex w-full flex-col text-kiokuDarkBlue ${className}`}
			>
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
					className={`flex flex-row items-center rounded-md bg-eggshell ${
						getStyle[inputFieldStyle]
					} ${
						initialised && !isValid
							? "border-kiokuRed"
							: "border-transparent focus-within:border-kiokuDarkBlue"
					}`}
				>
					<input
						id={id}
						type={type}
						className={`w-full border-none bg-transparent font-medium outline-none ${
							inputFieldSize && getSize[inputFieldSize]
						}`}
						ref={ref}
						onChange={(event) => {
							onChange(event);
							if (initialised) {
								checkValidity(event.target);
							}
						}}
						onKeyDown={(event) => {
							if (event.key === "Enter") {
								setInitialised(true);
								checkValidity(event?.currentTarget);
							}
						}}
						onBlur={(event) => {
							onBlur(event);
							setInitialised(true);
							checkValidity(event.target);
						}}
						pattern={inputPattern}
						{...props}
					/>
					{icon && getIcon(icon, id)}
					<Tooltip id={`tooltip-${id}`} content={tooltip} />
				</div>
			</div>
		);

		function checkValidity(input: HTMLInputElement) {
			if (statusIcon) {
				return;
			}
			input.setCustomValidity("");
			if (input.validity.valid) {
				setIcon("success");
				setTooltip("");
			} else {
				setIcon("error");
				if (input.validity.valueMissing) {
					setTooltip(t`Please fill out this field.`);
				} else {
					setTooltip(tooltipMessage);
					input.setCustomValidity(tooltipMessage);
				}
			}
		}
	}
);

InputField.displayName = "InputField";
