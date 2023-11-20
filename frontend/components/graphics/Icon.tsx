import * as icons from "react-feather";

export type IconName = keyof typeof icons;

interface IconProps {
	/**
	 * icon
	 */
	icon: IconName;
}

export const Icon = ({ icon, ...props }: IconProps & icons.IconProps) => {
	const FeatherIcon = icons[icon];
	return <FeatherIcon {...props}></FeatherIcon>;
};
