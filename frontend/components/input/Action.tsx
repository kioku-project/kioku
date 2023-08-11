import { ReactNode } from "react";

interface ActionProps {
	/**
	 * description
	 */
	description?: ReactNode;
	/**
	 * button
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
	className,
	...props
}: ActionProps) => {
	return (
		<>
			<div
				className={`flex w-full flex-col sm:w-2/3 md:w-3/4 lg:w-4/5 xl:w-5/6 ${
					className ?? ""
				}`}
				{...props}
			>
				{description}
			</div>
			{button}
		</>
	);
};
