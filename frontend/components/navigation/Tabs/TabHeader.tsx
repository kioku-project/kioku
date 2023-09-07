import { BarChart2, Copy, Layers, Mail, Settings, Users } from "react-feather";

interface TabHeaderProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * name
	 */
	name: string;
	/**
	 * style
	 */
	style: keyof typeof getIcon;
	/**
	 * Text that should be displayed as notification
	 */
	notification?: string;
	/**
	 * Additional classes
	 */
	className?: string;
}

const getIcon = {
	cards: <Copy size={20}></Copy>,
	decks: <Layers size={20}></Layers>,
	groups: <Users size={20}></Users>,
	invitations: <Mail size={20}></Mail>,
	settings: <Settings size={20}></Settings>,
	statistics: <BarChart2 size={20}></BarChart2>,
	user: <Users size={20}></Users>,
} as const;

/**
 * UI component for a TabHeader
 */
export const TabHeader = ({
	name,
	style,
	notification = "",
	className = "",
	...props
}: TabHeaderProps) => {
	return (
		<div
			className={`flex flex-row items-center space-x-2 ${className}`}
			{...props}
		>
			{getIcon[style]}
			<div>{name}</div>
			{notification && (
				<div className="relative flex h-full text-sm text-eggshell">
					<div className="absolute inline-flex h-full w-full animate-ping rounded-full bg-kiokuRed opacity-75"></div>
					<div className="relative flex h-full w-full justify-center rounded-full bg-kiokuRed px-2">
						{notification}
					</div>
				</div>
			)}
		</div>
	);
};
