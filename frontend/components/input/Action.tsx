import { ReactNode } from "react";

interface ActionProps {
	/**
	 * Description
	 */
	description?: ReactNode;
	/**
	 * Button
	 */
	button?: ReactNode;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for dislpaying a DangerAction
 */
export const Action = ({
	description,
	button,
	className = "",
	...props
}: ActionProps) => {
	// 1/6 1/5 1/4 1/3 2/5 1/2 3/5 2/3 3/4 4/5 5/6
	return (
		<>
			<div
				className={`flex w-full flex-col sm:w-3/5 md:w-3/5 lg:w-2/3 xl:w-3/4 ${className}`}
				{...props}
			>
				{description}
			</div>
			<div className="w-full sm:w-2/5 md:w-2/5 lg:w-1/3 xl:w-1/4">
				{button}
			</div>
		</>
	);
};
