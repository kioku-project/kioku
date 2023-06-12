import { ReactNode } from "react";
import { BarChart2, Copy, Layers, Mail, Settings, Users } from "react-feather";
import { Text } from "../../Text";

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
	style:
		| "cards"
		| "decks"
		| "groups"
		| "invitations"
		| "settings"
		| "statistics"
		| "user";
	/**
	 * Additional classes
	 */
	className?: string;
}

function getIcon(style: string): ReactNode {
	const getIcon: { [style: string]: ReactNode } = {
		cards: <Copy size={20}></Copy>,
		decks: <Layers size={20}></Layers>,
		groups: <Users size={20}></Users>,
		invitations: <Mail size={20}></Mail>,
		settings: <Settings size={20}></Settings>,
		statistics: <BarChart2 size={20}></BarChart2>,
		user: <Users size={20}></Users>,
	};
	return getIcon[style] ?? "";
}

/**
 * UI component for a TabHeader
 */
export const TabHeader = ({
	name,
	style,
	className,
	...props
}: TabHeaderProps) => {
	return (
		<div
			className={`flex flex-row items-center space-x-1 ${
				className ?? ""
			}`}
			{...props}
		>
			{getIcon(style)}
			<div>{name}</div>
		</div>
	);
};
