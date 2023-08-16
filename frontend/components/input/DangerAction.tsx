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
	 * Additional classes
	 */
	className?: string;
	/**
	 * Additional classes
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
						style="error"
						size="small"
						className="h-fit w-full sm:w-1/3 md:w-1/4 lg:w-1/5 xl:w-1/6"
						onClick={onClick}
					>
						{button}
					</Button>
				}
			></Action>
		</div>
	);
};
