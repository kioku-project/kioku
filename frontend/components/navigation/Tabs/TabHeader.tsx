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
	 * Text that should be displayed as notification
	 */
	notificationBadgeContent?: string;
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
	notificationBadgeContent = "",
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
			{notificationBadgeContent && (
				<div className="relative flex h-full text-sm text-eggshell">
					<div className="absolute inline-flex h-full w-full animate-[ping_1s_ease-out_3] rounded-full bg-kiokuRed opacity-75" />
					<div className="relative flex h-full w-full justify-center rounded-full bg-kiokuRed px-2">
						{notificationBadgeContent}
					</div>
				</div>
			)}
		</div>
	);
};
