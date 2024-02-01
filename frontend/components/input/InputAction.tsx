import { ChangeEventHandler, InputHTMLAttributes } from "react";

import { InputField } from "@/components/form/InputField";
import { Action } from "@/components/input/Action";
import { Button } from "@/components/input/Button";

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
}: InputActionProps & InputHTMLAttributes<HTMLInputElement>) => {
	return (
		<form
			id={id}
			onSubmit={(e) => e.preventDefault()}
			className={`flex flex-col justify-between space-y-1 p-3 sm:flex-row sm:items-center sm:space-x-3 ${className}`}
		>
			<Action
				description={
					<InputField
						id={`${id}InputFieldId`}
						name="actionInput"
						label={header}
						value={value}
						readOnly={disabled}
						inputFieldStyle="secondary"
						inputFieldSize="3xs"
						onChange={onChange}
						{...props}
					/>
				}
				button={
					<Button
						id={`${id}ButtonId`}
						buttonStyle={disabled ? "disabled" : "primary"}
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
