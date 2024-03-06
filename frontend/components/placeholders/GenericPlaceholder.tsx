import { MouseEventHandler, ReactNode } from "react";
import { Heart, Plus } from "react-feather";

import { Text } from "@/components/Text";
import { Icon, IconName } from "@/components/graphics/Icon";
import { clickOnEnter } from "@/util/utils";

interface GenericPlaceholderProps {
	/**
	 * Title
	 */
	title: string;
	/**
	 * Description
	 */
	description: string;
	/**
	 * Button text
	 */
	buttonText?: string;
	/**
	 * Icon name
	 */
	iconName: IconName;
	/**
	 * Onclick function
	 */
	onClick?: MouseEventHandler;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a DangerAction
 */
export const GenericPlaceholder = ({
	title,
	description,
	buttonText,
	iconName,
	onClick,
	className = "",
	...props
}: GenericPlaceholderProps) => {
	return (
		<div {...props} className={className}>
			<div className="flex h-[8rem] w-full items-center justify-between space-x-3	 rounded-md border-2 border-dashed border-neutral-300 px-4 py-3 ">
				<div className="flex space-x-4">
					<Icon
						icon={iconName}
						className="size-14 self-center stroke-neutral-300 stroke-2"
					/>
					<div className="flex w-fit flex-col">
						<Text className="pb-1 text-base font-bold">
							{title}
						</Text>
						<Text className="text-sm font-light leading-snug">
							{description}
						</Text>
					</div>
				</div>
				{buttonText && (
					<button
						onClick={onClick}
						onKeyUp={clickOnEnter}
						tabIndex={0}
						className="rounded bg-black px-3 py-2 text-sm leading-snug text-white hover:scale-105"
					>
						{buttonText}
					</button>
				)}
			</div>
		</div>
	);
};
