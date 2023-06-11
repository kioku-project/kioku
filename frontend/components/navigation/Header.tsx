import React from "react";
import { Logo } from "../graphics/Logo";
import { Inter } from "next/font/google";
import { Button } from "../input/Button";
import router from "next/router";
import { ArrowRight, LogOut } from "react-feather";
import { authedFetch } from "../../util/reauth"

const inter = Inter({
	weight: ["200", "400"],
	subsets: ["latin"],
});

interface HeaderProps {
	/**
	 * show login or logout button
	 */
	login?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for diplaying the Header
 */
export const Header = ({ login, className }: HeaderProps) => {
	return (
		<div
			className={`flex items-center justify-between p-5 md:p-10 ${
				className ?? ""
			}`}
		>
			<Logo></Logo>
			{login ? (
				<div className="flex flex-row">
					<LogOut
						className="text-kiokuDarkBlue hover:cursor-pointer"
						onClick={
							async () => {
								const response = await authedFetch("/api/logout", {
									method: "POST"
								});
								if (response?.ok) {
									location.reload()
								}
							}
						}
					></LogOut>
				</div>
			) : (
				<Button
					id="loginButton"
					style="secondary"
					className="invisible h-full justify-end sm:visible"
					onClick={() => router.push("/")}
				>
					Login <ArrowRight className="ml-1 h-2/3"></ArrowRight>
				</Button>
			)}
		</div>
	);
};
