import { Trans } from "@lingui/macro";
import { hasCookie } from "cookies-next";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { LogOut } from "react-feather";

import { authedFetch } from "../../util/reauth";
import { Logo } from "../graphics/Logo";
import { Button } from "../input/Button";

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
		if (router.pathname == "/login") {
			setLoggedIn(undefined);
		} else {
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
			/>
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
				/>
			)}
			{loggedIn == false && (
				<Button
					id="loginButton"
					buttonStyle="secondary"
					buttonSize="sm"
					buttonTextSize="xs"
					buttonIcon="ArrowRight"
					className="invisible sm:visible"
					onClick={() => router.push("/login")}
				>
					<Trans>Login</Trans>
				</Button>
			)}
		</nav>
	);
};
