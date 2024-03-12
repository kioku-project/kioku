import clsx from "clsx";
import React, { Children, MouseEventHandler, isValidElement } from "react";
import { JSXElementConstructor, ReactElement, ReactNode } from "react";
import { ChevronUp } from "react-feather";

interface SpeechBubbleProps {
	/**
	 * Value
	 */
	// value: ReactNode;
	// /**
	//  * Parent
	//  */
	// parent: ReactNode;
	/**
	 * Align
	 */
	align: "left" | "center" | "right" | string;
	/**
	 * Show
	 */
	show: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * SelectionField options
	 */
	children: ReactNode;
	/**
	 *
	 */
	onHide: MouseEventHandler;
}

/**
 * UI component for displaying a Speech Bubble
 */
export const SpeechBubble = ({
	// value,
	// parent,
	align,
	className = "",
	show,
	onHide,
	children,
	...props
}: SpeechBubbleProps) => {
	const parent = Children.map(children, (child) => {
		if (
			isValidElement(child) &&
			typeof child.type !== "string" &&
			child.type.name == "SpeechBubbleParent"
		) {
			return child;
		}
	});
	const content = Children.map(children, (child) => {
		if (
			isValidElement(child) &&
			typeof child.type !== "string" &&
			child.type.name == "SpeechBubbleContent"
		) {
			console.log("SpeechBubble");
			console.log(child);
			return child;
		}
	});

	return show ? (
		<div {...props} className={clsx(className, "relative")}>
			<button onClick={onHide} className="fixed inset-0 cursor-default" />
			<div className="flex h-fit w-fit flex-col">
				{parent}

				<div className="absolute left-1/2 top-[100%] z-10 h-0 w-0 -translate-x-1/2 self-center border-b-[1rem] border-l-[0.75rem] border-r-[0.75rem] border-t-0 border-transparent border-b-black">
					<div
						className={clsx(
							" absolute h-0 w-[50vw] translate-y-[0.75rem] sm:w-[50vw] md:w-[35vw] lg:w-[25vw] xl:w-[25vw] 2xl:w-[20vw]",
							align === "center" && "left-1/2 -translate-x-1/2",
							align === "left" && "-left-8 ",
							align === "right" && "-right-8"
						)}
					>
						<div
							className={clsx(
								"  w-fit min-w-[30%] max-w-full space-y-2 rounded-lg border-b border-l border-r border-neutral-100  bg-black  p-3 text-sm text-white",
								align === "center" && "mx-auto",
								align === "left" && "mr-auto",
								align === "right" && "ml-auto"
							)}
						>
							{content}
						</div>
					</div>
				</div>
			</div>
		</div>
	) : (
		<div {...props} className={clsx(className, "relative")}>
			{parent}
		</div>
	);
};
