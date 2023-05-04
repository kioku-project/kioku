import React from "react";
import { Logo } from "../graphics/Logo";
import { Inter } from "next/font/google";
import { Button } from "../input/Button";
import router from "next/router";

const inter = Inter({
	weight: ["200", "400"],
	subsets: ["latin"],
});

interface HeaderProps {
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for diplaying the Header
 */
export const Header = ({ className }: HeaderProps) => {
	return (
		<div
			className={`flex items-center justify-between p-5 md:p-10 ${
				className ?? ""
			}`}
		>
			<Logo></Logo>
			<Button
				id="loginButton"
				value="Login &rarr;"
				style="secondary"
				className="invisible justify-end sm:visible"
				onClick={() => {
					router.push("/login");
				}}
			></Button>
		</div>
	);
};
