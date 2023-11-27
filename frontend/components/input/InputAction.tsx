import { ChangeEventHandler } from "react";

import { InputField } from "../form/InputField";
import { Action } from "./Action";
import { Button } from "./Button";

interface InputActionProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * Header
	 */
	header?: string;
	/**
	 * Input value
	 */
	value?: string;
	/**
	 * Button content
	 */
	button: string;
	/**
	 * Is the InputAction disabled?
	 */
	disabled?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Change handler
	 */
	onChange: ChangeEventHandler<HTMLInputElement>;
	/**
	 * Click handler
	 */
	onClick?: () => void;
}

/**
 * UI component for dislpaying a InputAction
 */
export const InputAction = ({
	id,
	header,
	value,
	button,
	disabled = false,
	className = "",
	onChange,
	onClick,
	...props
}: InputActionProps) => {
	return (
		<form
			id={id}
			onSubmit={(e) => e.preventDefault()}
			className={`flex flex-col justify-between space-y-1 p-3 sm:flex-row sm:items-center sm:space-x-3 ${className}`}
			{...props}
		>
			<Action
				description={
					<InputField
						id={`${id}InputFieldId`}
						type="text"
						name="actionInput"
						label={header}
						value={value}
						statusIcon="none"
						inputFieldStyle="tertiary"
						inputFieldSize="3xs"
						readOnly={disabled}
						onChange={onChange}
					/>
				}
				button={
					<Button
						id={`${id}ButtonId`}
						buttonStyle={disabled ? "disabled" : "primary"}
						buttonSize="sm"
						buttonTextSize="3xs"
						className="w-full justify-center"
						onClick={() => !disabled && onClick?.()}
					>
						{button}
					</Button>
				}
			/>
		</form>
	);
};
