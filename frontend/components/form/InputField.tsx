import React, {
	ChangeEventHandler,
	FocusEventHandler,
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
	 * InputField label that will be displayed above the InputField
	 */
	label?: string;
	/**
	 * InputField value
	 */
	value?: string;
	/**
	 * InputField placeholder that will be displayed when the value is empty
	 */
	placeholder?: string;
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
	style?: keyof typeof getStyle;
	/**
	 * InputField size
	 */
	size?: keyof typeof getSize;
	/**
	 * Is the InputField responsive?
	 */
	responsive?: boolean;
	/**
	 * Is the InputField required?
	 */
	required?: boolean;
	/**
	 * Minimal input length
	 */
	minLength?: number;
	/**
	 * pattern the value has to match to be valid
	 */
	pattern?: string;
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
	/**
	 * optional onBlur handler
	 */
	onBlur?: FocusEventHandler<HTMLInputElement>;
}

const getStyle = {
	primary: "border-2 py-1 pl-1.5 pr-1 sm:py-1 sm:pr-3",
	secondary: "border-none text-kiokuDarkBlue",
	tertiary: "border-none text-kiokuLightBlue",
};

const getSize = {
	xs: "text-xs sm:text-xs md:text-xs lg:text-sm xl:text-base",
	sm: "text-xs sm:text-xs md:text-sm lg:text-base xl:text-lg",
	md: "text-xs sm:text-sm md:text-base lg:text-lg xl:text-xl",
	lg: "text-sm sm:text-base md:text-lg lg:text-xl xl:text-2xl",
	xl: "text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl",
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
			id,
			type,
			name,
			label,
			statusIcon,
			tooltipMessage = `Please enter a valid ${type}.`,
			style = "primary",
			size = "md",
			responsive = true,
			pattern,
			className = "",
			onChange = () => {},
			onBlur = () => {},
			...props
		}: InputFieldProps,
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
						className="font-bold"
						size={size}
						responsive={responsive}
					>
						{label}
					</Text>
				</label>
				<div
					className={`flex flex-row items-center rounded-md bg-eggshell ${
						getStyle[style]
					} ${
						initialised && !isValid
							? "border-kiokuRed"
							: "border-transparent focus-within:border-kiokuDarkBlue"
					}`}
				>
					<input
						id={id}
						type={type}
						name={name}
						className={`w-full border-none bg-transparent font-medium outline-none ${
							size && getSize[size]
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
					setTooltip("Please fill out this field.");
				} else {
					setTooltip(tooltipMessage);
					input.setCustomValidity(tooltipMessage);
				}
			}
		}
	}
);

InputField.displayName = "InputField";
