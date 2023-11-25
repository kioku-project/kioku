import { Text } from "../Text";
import { Action } from "./Action";
import { Button } from "./Button";

interface DangerActionProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * Header
	 */
	header?: string;
	/**
	 * Description
	 */
	description?: string;
	/**
	 * Button
	 */
	button?: string;
	/**
	 * Is the DangerAction disabled?
	 */
	disabled?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Click handler
	 */
	onClick?: () => void;
}

/**
 * UI component for dislpaying a DangerAction
 */
export const DangerAction = ({
	id,
	header,
	description,
	button,
	disabled,
	className = "",
	onClick,
	...props
}: DangerActionProps) => {
	return (
		<div
			id={id}
			className={`flex flex-col justify-between space-y-1 p-3 sm:flex-row sm:items-center sm:space-x-3 ${className}`}
			{...props}
		>
			<Action
				description={
					<>
						<Text
							textStyle="primary"
							textSize="3xs"
							className="font-bold"
						>
							{header}
						</Text>
						<Text
							textStyle="secondary"
							textSize="3xs"
							className="font-medium"
						>
							{description}
						</Text>
					</>
				}
				button={
					<Button
						id={`${id}ButtonId`}
						buttonStyle={disabled ? "disabled" : "error"}
						buttonSize="sm"
						className="w-full"
						onClick={() => !disabled && onClick?.()}
					>
						{button}
					</Button>
				}
			/>
		</div>
	);
};
