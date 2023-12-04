import { Inter } from "next/font/google";
import Image from "next/image";
import Link, { LinkProps } from "next/link";

import { Text } from "@/components/Text";
import kiokuLogo from "@/public/kioku-logo.svg";
import { Size } from "@/types/Size";

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
	 * Text size
	 */
	textSize?: Size;
	/**
	 * Logo size
	 */
	logoSize?: keyof typeof getSize;
	/**
	 * Additional classes
	 */
	className?: string;
}

const getSize = {
	sm: "w-12 sm:w-14 md:w-16 lg:w-20",
	md: "w-14 sm:w-16 md:w-20 lg:w-24",
	lg: "w-16 sm:w-20 md:w-24 lg:w-28",
} as const;

/**
 * UI component for displaying the Kioku Logo
 */
export const Logo = ({
	text = true,
	textSize = "lg",
	logoSize = "md",
	className = "",
	...props
}: LogoProps & LinkProps) => {
	return (
		<Link
			className={`flex flex-row items-center hover:cursor-pointer ${className}`}
			{...props}
		>
			<Image
				src={kiokuLogo}
				alt="Kioku"
				className={`${getSize[logoSize]}`}
			/>
			{text && (
				<Text
					textSize={textSize}
					className={`ml-3 font-extralight tracking-[0.5em] ${inter.className}`}
				>
					Kioku
				</Text>
			)}
		</Link>
	);
};
