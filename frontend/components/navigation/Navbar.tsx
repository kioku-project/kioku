import { hasCookie } from "cookies-next";
import { Inter } from "next/font/google";
import router, { useRouter } from "next/router";
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
	const router = useRouter();
	const [loggedIn, setLoggedIn] = useState<boolean>();
	useEffect(() => {
		if (router.pathname != "/login") {
			setLoggedIn(hasCookie("access_token"));
		}
	});
	return (
		<nav
			className={`flex items-center justify-between p-5 md:p-10 ${className}`}
		>
			<Logo
				onClick={() =>
					loggedIn ? router.push("/") : router.push("/home")
				}
			></Logo>
			{loggedIn == true && (
				<LogOut
					className="text-kiokuDarkBlue hover:cursor-pointer"
					onClick={async () => {
						const response = await authedFetch("/api/logout", {
							method: "POST",
						});
						if (response?.ok) {
							router.replace("/home");
						}
					}}
				></LogOut>
			)}
			{loggedIn == false && (
				<Button
					id="loginButton"
					style="secondary"
					className="invisible h-full justify-end sm:visible"
					onClick={() => router.push("/")}
				>
					Login <ArrowRight className="ml-1 h-2/3"></ArrowRight>
				</Button>
			)}
		</nav>
	);
};
