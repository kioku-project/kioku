import { Action } from "./Action";
import { Button } from "./Button";

interface DangerActionProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * header
	 */
	header?: string;
	/**
	 * description
	 */
	description?: string;
	/**
	 * button
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
	 * click handler
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
						<div className="font-bold text-kiokuDarkBlue">
							{header}
						</div>
						<div className="font-medium text-kiokuLightBlue">
							{description}
						</div>
					</>
				}
				button={
					<Button
						id={`${id}ButtonId`}
						style={disabled ? "disabled" : "error"}
						size="sm"
						className="h-fit w-full sm:w-1/3 md:w-1/4 lg:w-1/5 xl:w-1/6"
						onClick={() => !disabled && onClick?.()}
					>
						{button}
					</Button>
				}
			></Action>
		</div>
	);
};
