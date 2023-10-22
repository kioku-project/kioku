import React, { ChangeEventHandler, Ref, forwardRef } from "react";

interface TextAreaProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * TextArea name
	 */
	name: string;
	/**
	 * TextArea value
	 */
	value?: string;
	/**
	 * TextArea placeholder
	 */
	placeholder?: string;
	/**
	 * Is the TextArea read only?
	 */
	readOnly?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional change handler
	 */
	onChange?: ChangeEventHandler<HTMLTextAreaElement>;
}

/**
 * UI component for multiline text inputs
 */
export const TextArea = forwardRef(
	(
		{ name, className = "", ...props }: TextAreaProps,
		ref: Ref<HTMLTextAreaElement>
	) => (
		<div className={`flex h-full w-full flex-col ${className}`}>
			<textarea
				name={name}
				className={` h-full w-full resize-none rounded-md border-eggshell bg-eggshell font-medium text-kiokuLightBlue outline-none`}
				ref={ref}
				{...props}
			/>
		</div>
	)
);
TextArea.displayName = "TextArea";
