import { Trans, msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { GetStaticProps } from "next";
import { Inter } from "next/font/google";
import Head from "next/head";
import Image from "next/image";
import { useRouter } from "next/router";
import { useEffect, useRef, useState } from "react";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

import { Text } from "../components/Text";
import { FormButton } from "../components/form/FormButton";
import { InputField } from "../components/form/InputField";
import { checkAccessTokenValid } from "../util/reauth";
import { loadCatalog } from "./_app";

const inter = Inter({
	weight: ["200", "400"],
	subsets: ["latin"],
});

export const getStaticProps: GetStaticProps = async (ctx) => {
	const translation = await loadCatalog(ctx.locale!);
	return {
		props: {
			translation,
		},
	};
};

export default function Page() {
	const router = useRouter();
	const [login, setLogin] = useState(true); // true = login, false = register
	const form = useRef<HTMLFormElement>(null);
	const emailInput = useRef<HTMLInputElement>(null);
	const nameInput = useRef<HTMLInputElement>(null);
	const passwordInput = useRef<HTMLInputElement>(null);
	const repeatPasswordInput = useRef<HTMLInputElement>(null);
	const [password, setPassword] = useState("");

	const { _ } = useLingui();

	useEffect(() => {
		if (checkAccessTokenValid()) {
			router.push("/");
		}
	}, []);

	return (
		<>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
				<link
					rel="alternate"
					hrefLang="en"
					href="https://app.kioku.dev/login"
				/>
				<link
					rel="alternate"
					hrefLang="de"
					href="https://app.kioku.dev/de/login"
				/>
			</Head>

			<div className="min-w-screen flex flex-1 items-center justify-center sm:p-5 md:p-10">
				<div className="flex min-h-screen w-full flex-col items-center justify-evenly bg-blue-100 p-5 align-middle sm:min-h-fit sm:rounded-3xl sm:p-10 md:flex-row xl:w-5/6">
					<div className="m-5 flex w-2/3 flex-col items-center md:m-10 md:w-1/2 md:justify-center">
						<div className="relative mb-5 h-[120px] w-full">
							<Image
								src="/kioku-logo.svg"
								alt="Kioku Logo"
								className="object-contain"
								fill
							/>
						</div>
						<p
							className={`${inter.className} text-clip indent-[0.5em] text-4xl font-extralight tracking-[0.5em] sm:text-5xl md:text-6xl`}
						>
							Kioku
						</p>
					</div>
					<div
						className={`flex w-full flex-col items-center rounded-2xl bg-kiokuLightBlue p-3 sm:w-5/6 sm:p-5 md:w-1/2 lg:w-1/3 ${inter.className}`}
					>
						<Text
							textSize="md"
							className="text-center font-bold leading-9 tracking-tight text-kiokuDarkBlue"
						>
							{login ? (
								<Trans>Sign in to your account</Trans>
							) : (
								<Trans>Create an account</Trans>
							)}
						</Text>
						{forms()}

						<Text
							textSize="3xs"
							className="text-center text-gray-500"
						>
							{login ? (
								<Trans>Not registered?</Trans>
							) : (
								<Trans>Already registered?</Trans>
							)}
							<a
								className="whitespace-nowrap font-semibold text-kiokuDarkBlue transition hover:cursor-pointer hover:text-eggshell"
								onClick={() => {
									emailInput.current?.focus();
									setLogin((prev) => !prev);
								}}
								onKeyUp={(event) => {
									if (event.key === "Enter") {
										event.target.dispatchEvent(
											new Event("click", {
												bubbles: true,
											})
										);
									}
								}}
								tabIndex={0}
							>
								<span> </span>
								{login ? (
									<Trans>Create an account</Trans>
								) : (
									<Trans>Sign in</Trans>
								)}
							</a>
						</Text>
					</div>
				</div>
			</div>
		</>
	);

	function forms() {
		return (
			<form
				onSubmit={(e) => e.preventDefault()}
				className="my-3 flex w-5/6 flex-col items-center space-y-4 sm:my-5"
				ref={form}
			>
				<InputField
					id="email"
					type="email"
					name="email"
					label={_(msg`Email`)}
					required={true}
					inputFieldSize="xs"
					ref={emailInput}
				/>
				{!login && (
					<InputField
						id="name"
						type="text"
						name="name"
						label={_(msg`Name`)}
						required={true}
						inputFieldSize="xs"
						ref={nameInput}
					/>
				)}
				<InputField
					id="password"
					type="password"
					name="password"
					label={_(msg`Password`)}
					required={true}
					minLength={3}
					inputFieldSize="xs"
					onChange={(event) => {
						setPassword(event.target.value);
					}}
					ref={passwordInput}
				/>
				{!login && (
					<InputField
						id="passwordRepeat"
						type="password"
						name="passwordRepeat"
						label={_(msg`Repeat Password`)}
						tooltipMessage={_(msg`Passwords have to match.`)}
						required={true}
						minLength={3}
						pattern={password}
						inputFieldSize="xs"
						ref={repeatPasswordInput}
					/>
				)}

				<FormButton
					id={login ? "login" : "register"}
					value={login ? _(msg`Login`) : _(msg`Register`)}
					size="sm"
					className="w-full"
					onClick={() => {
						if (login) {
							loginLogic()
								.then((result) => {})
								.catch((error) => {});
						} else {
							registerLogic()
								.then((result) => {})
								.catch((error) => {});
						}
					}}
				/>
			</form>
		);
	}

	async function loginLogic() {
		if (!form.current?.checkValidity()) {
			return;
		}
		const response = await fetch("/api/login", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				userEmail: emailInput.current?.value,
				userPassword: passwordInput.current?.value,
			}),
		});
		if (response.ok) {
			toast.info(<Trans>Logged in!</Trans>, { toastId: "accountToast" });
			router.push("/");
		} else {
			toast.error(<Trans>Wrong username or password</Trans>, {
				toastId: "accountToast",
			});
		}
	}

	async function registerLogic() {
		if (
			!form.current?.checkValidity() ||
			passwordInput.current?.value !== repeatPasswordInput.current?.value
		) {
			return;
		}
		const response = await fetch("/api/register", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				userEmail: emailInput.current?.value,
				userName: nameInput.current?.value,
				userPassword: passwordInput.current?.value,
			}),
		});
		if (response.ok) {
			toast.info(<Trans>Account created!</Trans>, {
				toastId: "accountToast",
			});
			setLogin(true);
			emailInput.current?.focus();
		} else {
			toast.error(<Trans>Account already exists!</Trans>, {
				toastId: "accountToast",
			});
		}
	}
}
