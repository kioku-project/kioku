import { MouseEventHandler } from "react";

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
			<div
				className="flex min-h-[6.5rem] w-full flex-row gap-x-4 rounded-md border-2 border-dashed border-neutral-300 bg-gradient-to-l to-60% p-3	  hover:from-neutral-50 sm:min-h-28 md:min-h-32 lg:min-h-32
			 "
			>
				<div className="flex self-center rounded-md bg-neutral-100 p-4">
					<Icon
						icon={iconName}
						className=" mx-auto size-10
					self-center stroke-neutral-600 stroke-2  p-0"
					/>
				</div>
				<div className="flex w-full flex-col gap-1 self-center text-neutral-800">
					<Text className=" text-base font-bold ">{title}</Text>
					<Text className=" text-sm font-light">{description}</Text>
					{buttonText && (
						<button
							onClick={onClick}
							onKeyUp={clickOnEnter}
							tabIndex={0}
							className="h-fit w-fit rounded bg-black p-1 px-2 text-xs font-light leading-snug text-white hover:scale-105"
						>
							{buttonText}
						</button>
					)}
				</div>
			</div>
		</div>
	);
};
