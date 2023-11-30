import { Trans } from "@lingui/macro";
import { hasCookie } from "cookies-next";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { ArrowRight, LogOut } from "react-feather";

import { Logo } from "@/components/graphics/Logo";
import { Button } from "@/components/input/Button";
import { authedFetch } from "@/util/reauth";

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
	if (loggedIn == undefined) {
		return <></>;
	}
	return (
		<nav
			className={`flex items-center justify-between p-5 md:p-10 ${className}`}
		>
			<Link href={loggedIn ? "/" : "/home"}>
				<Logo />
			</Link>
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
					href="/login"
					buttonStyle="secondary"
					buttonSize="sm"
					buttonTextSize="xs"
					buttonIcon={<ArrowRight size={16} />}
					className="invisible sm:visible"
				>
					<Trans>Login</Trans>
				</Button>
			)}
		</nav>
	);
};
