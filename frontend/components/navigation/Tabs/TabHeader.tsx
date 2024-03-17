import { Icon, IconName } from "@/components/graphics/Icon";

interface TabHeaderProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * Name
	 */
	name: string;
	/**
	 * Style
	 */
	icon: IconName;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for a TabHeader
 */
export const TabHeader = ({
	name,
	icon,
	className = "",
	...props
}: TabHeaderProps) => {
	return (
		<div
			className={`flex flex-col items-center justify-center sm:flex-row sm:space-x-2 ${className}`}
			{...props}
		>
			<Icon icon={icon} />

			<div>{name}</div>
		</div>
	);
};
