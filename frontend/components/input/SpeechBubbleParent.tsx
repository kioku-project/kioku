import { Children, ReactNode, isValidElement } from "react";

interface SpeechBubbleParentProps {
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
export const SpeechBubbleParent = ({
	className = "",
	children,
	...props
}: SpeechBubbleParentProps) => {
	return (
		<>
			{Children.map(children, (child) => {
				if (!isValidElement(child)) return null;
				return child;
			})}
		</>
	);
};
