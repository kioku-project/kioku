import { Children, ReactNode, isValidElement } from "react";

interface SpeechBubbleContentProps {
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Children
	 */
	children: ReactNode;
}

/**
 * UI component for displaying a DangerAction
 */
export const SpeechBubbleContent = ({
	className = "",
	children,
	...props
}: SpeechBubbleContentProps) => {
	return (
		<>
			{Children.map(children, (child) => {
				if (!isValidElement(child)) return null;
				return child;
			})}
		</>
	);
};
