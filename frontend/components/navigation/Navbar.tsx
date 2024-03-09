import { Trans, t } from "@lingui/macro";
import { hasCookie } from "cookies-next";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { ArrowRight, LogOut } from "react-feather";
import toast from "react-hot-toast";

import { Logo } from "@/components/graphics/Logo";
import { Button } from "@/components/input/Button";
import { logoutRoute } from "@/util/endpoints";
import { authedFetch } from "@/util/reauth";
import { handleWithToast } from "@/util/toasts";

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
	}, [router.pathname]);
	if (loggedIn == undefined) {
		return <></>;
	}
	return (
		<nav
			className={`flex items-center justify-between p-5 md:p-10 ${className}`}
		>
			<Logo href={loggedIn ? "/" : "/home"} />
			{loggedIn == true && (
				<LogOut
					className="cursor-pointer text-kiokuDarkBlue"
					onClick={async () => {
						const response = await handleWithToast(
							authedFetch(logoutRoute, {
								method: "POST",
							}),
							"logoutToastID",
							t`Logged out`
						);
						if (response.ok) router.replace("/home");
					}}
				/>
			)}
			{loggedIn == false && (
				<Button
					id="loginButton"
					href="/login"
					buttonStyle="tertiary"
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
