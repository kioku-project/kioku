import { Text } from "@/components/Text";
import { Icon, IconName } from "@/components/graphics/Icon";
import { Size } from "@/types/Size";

export type IconLabelType = {
	icon?: IconName;
	header: string;
	description?: string;
};

interface IconLabelProps {
	/**
	 * Icon, Header, Description
	 */
	iconLabel: IconLabelType;
	/**
	 * Icon size
	 */
	iconSize?: number;
	/**
	 * Text size
	 */
	textSize?: Size;
	/**
	 * Icon and Header color
	 */
	color?: string;
	/**
	 * Additional classes
	 */
	className?: string;
}

export const IconLabel = ({
	iconLabel,
	iconSize = 16,
	textSize = "5xs",
	color = "",
	className = "",
}: IconLabelProps) => {
	return (
		<div className={`flex flex-row items-center space-x-1 ${className}`}>
			{iconLabel.icon && (
				<Icon icon={iconLabel.icon} size={iconSize} className={color} />
			)}
			<Text textSize={textSize} className="flex flex-row items-center">
				<div className="space-x-1">
					<span className={color}>{iconLabel.header}</span>
					<span>{iconLabel.description}</span>
				</div>
			</Text>
		</div>
	);
};
