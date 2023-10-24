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
	 * Function to execute on Shift Enter
	 */
	save: () => void;
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
		{ name, className = "", save, ...props }: TextAreaProps,
		ref: Ref<HTMLTextAreaElement>
	) => (
		<div className={`flex h-full w-full flex-col ${className}`}>
			<textarea
				name={name}
				className="h-full w-full resize-none rounded-md bg-transparent font-medium text-kiokuLightBlue outline-none"
				ref={ref}
				{...props}
				onKeyDown={(event) => {
					if (event.key === "Enter" && !event.shiftKey) {
						save();
						event.preventDefault();
					}
				}}
			/>
		</div>
	)
);
TextArea.displayName = "TextArea";
