import React, { Dispatch, ReactNode, SetStateAction } from "react";
import { X } from "react-feather";

import { Text } from "../Text";

export interface ModalProps {
	/**
	 * Unique identifier
	 */
	id?: string;
	/**
	 * Is the modal visible?
	 */
	visible: boolean;
	/**
	 * Change visibility
	 */
	setVisible: Dispatch<SetStateAction<boolean>>;
	/**
	 * Header
	 */
	header?: string;
	/**
	 * Content
	 */
	children?: ReactNode;
	/**
	 * Additional classNamees
	 */
	className?: string;
}

/**
 * UI component for text
 */
export const Modal = ({
	visible,
	setVisible,
	header,
	className = "",
	children,
	...props
}: ModalProps) => {
	return (
		<div className={`${className}`} {...props}>
			<div className="relative z-10">
				{visible && (
					<>
						<div className="fixed inset-0 bg-gray-500 bg-opacity-75"></div>
						<div className="fixed inset-0 w-screen overflow-hidden">
							<div className="flex min-h-full items-end justify-center text-center sm:items-center sm:p-0">
								<div className="relative w-full space-y-5  rounded-t-lg bg-white p-5 text-left shadow-xl sm:max-w-xl sm:rounded-lg">
									<div className="flex w-full justify-between">
										<Text
											textSize="xs"
											className="font-bold"
										>
											{header}
										</Text>
										<X
											size={16}
											className="cursor-pointer text-gray-400"
											onClick={() => {
												setVisible(false);
											}}
										></X>
									</div>
									{children}
								</div>
							</div>
						</div>
					</>
				)}
			</div>
		</div>
	);
};
