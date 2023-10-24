import { hasCookie } from "cookies-next";
import { Inter } from "next/font/google";
import router from "next/router";
import React, { useEffect, useState } from "react";
import { ArrowRight, LogOut } from "react-feather";

import { authedFetch } from "../../util/reauth";
import { Logo } from "../graphics/Logo";
import { Button } from "../input/Button";

const inter = Inter({
	weight: ["200", "400"],
	subsets: ["latin"],
});

interface NavbarProps {
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for diplaying the Navbar
 */
export const Navbar = ({ className = "" }: NavbarProps) => {
	const [loggedIn, setLoggedIn] = useState(false);
	useEffect(() => {
		setLoggedIn(hasCookie("access_token"));
	}, []);
	return (
		<div
			className={`flex items-center justify-between p-5 md:p-10 ${className}`}
		>
			{loggedIn ? (
				<>
					<Logo onClick={() => router.push("/")}></Logo>
					<div className="flex flex-row">
						<LogOut
							className="text-kiokuDarkBlue hover:cursor-pointer"
							onClick={async () => {
								const response = await authedFetch(
									"/api/logout",
									{
										method: "POST",
									}
								);
								if (response?.ok) {
									location.reload();
								}
							}}
						></LogOut>
					</div>
				</>
			) : (
				<>
					<Logo onClick={() => router.push("/home")}></Logo>
					<Button
						id="loginButton"
						style="secondary"
						className="invisible h-full justify-end sm:visible"
						onClick={() => router.push("/")}
					>
						Login <ArrowRight className="ml-1 h-2/3"></ArrowRight>
					</Button>
				</>
			)}
		</div>
	);
};
