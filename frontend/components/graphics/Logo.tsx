import { Inter } from "next/font/google";
import Image from "next/image";
import router from "next/router";
import React from "react";

import kiokuLogo from "../../public/kioku-logo.svg";
import { Text } from "../Text";

const inter = Inter({
	weight: ["200", "400"],
	subsets: ["latin"],
});

interface LogoProps {
	/**
	 * Should the text be displayed
	 */
	text?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * alternative click handler
	 */
	onClick?: () => void;
}

/**
 * UI component for displaying the Kioku Logo
 */
export const Logo = ({ className = "", text = true, onClick }: LogoProps) => {
	return (
		<div
			className={`flex flex-row items-center hover:cursor-pointer ${className}`}
			onClick={() => {
				if (onClick) {
					onClick();
				} else {
					router.push("/");
				}
			}}
		>
			<Image
				src={kiokuLogo}
				alt="Kioku"
				height={0}
				width={0}
				className="w-16 hover:cursor-pointer md:w-20 lg:w-28"
			/>
			{text && (
				<Text
					size="lg"
					className={`ml-3 font-extralight tracking-[0.5em] ${inter.className}`}
				>
					Kioku
				</Text>
			)}
		</div>
	);
};
