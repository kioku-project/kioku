import React from "react";
import Image from "next/image";
import { Inter } from "next/font/google";
import router from "next/router";

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
}

/**
 * UI component for displaying the Kioku Logo
 */
export const Logo = ({ className, text = true }: LogoProps) => {
	return (
		<div
			className={`flex flex-row items-center hover:cursor-pointer ${
				className ?? ""
			}`}
			onClick={() => {
				router.push("/home");
			}}
		>
			<Image
				src="./kioku-logo.svg"
				alt="Koiku"
				height={0}
				width={0}
				className="w-16 hover:cursor-pointer md:w-20 lg:w-24"
			/>
			{text && (
				<p
					className={`ml-3 text-lg font-extralight tracking-[0.5em] sm:text-xl md:text-2xl lg:ml-5 lg:text-3xl ${inter.className} `}
				>
					Kioku
				</p>
			)}
		</div>
	);
};
