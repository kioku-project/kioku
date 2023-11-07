import { ChangeEventHandler } from "react";

import { FormButton } from "../form/FormButton";
import { InputField } from "../form/InputField";
import { Action } from "./Action";

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
					></InputField>
				}
				button={
					<FormButton
						id={`${id}ButtonId`}
						value={button}
						style={disabled ? "disabled" : "primary"}
						size="sm"
						className="h-fit w-full sm:w-1/3 md:w-1/4 lg:w-1/5 xl:w-1/6"
						onClick={() => !disabled && onClick?.()}
					></FormButton>
				}
			></Action>
		</form>
	);
};
