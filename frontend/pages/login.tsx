import { Trans, msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { GetStaticProps } from "next";
import Head from "next/head";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect, useRef, useState } from "react";
import { ArrowRight, Check } from "react-feather";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

import { Text } from "../components/Text";
import { InputField } from "../components/form/InputField";
import { Logo } from "../components/graphics/Logo";
import { Button } from "../components/input/Button";
import { loadCatalog } from "./_app";

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
	const { _ } = useLingui();

	const [login, setLogin] = useState(true); // true = login, false = register
	const form = useRef<HTMLFormElement>(null);
	const emailInput = useRef<HTMLInputElement>(null);
	const nameInput = useRef<HTMLInputElement>(null);
	const passwordInput = useRef<HTMLInputElement>(null);
	const repeatPasswordInput = useRef<HTMLInputElement>(null);
	const [password, setPassword] = useState("");
	const [passwordsMatching, setPasswordsMatching] = useState(false);
	const passwordMinLength = 3;

	useEffect(() => {
		(async () => {
			const response = await fetch("/api/reauth");
			if (response.status === 200) {
				router.replace("/");
			}
		})();
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
			<div className="min-w-screen flex flex-1 bg-[#F8F8F8]">
				<div className="h-full w-full bg-gradient-to-bl from-[#FF83FA]/20 to-50%">
					<div className="flex h-full w-full items-center justify-center bg-gradient-to-tr from-[#83DAFF]/20 p-3 sm:p-5">
						<div className="flex w-80 flex-col items-center space-y-3 rounded-md bg-white p-5 shadow-[0_35px_60px_-15px_rgba(0,0,0,0.2)] md:px-7">
							<Link href="/home">
								<Logo
									text={false}
									logoSize="sm"
									className="p-3"
								/>
							</Link>
							<form
								className="w-full space-y-3"
								onSubmit={(event) => {
									event.preventDefault();
								}}
								ref={form}
							>
								<InputField
									id="emailInputFieldId"
									type="email"
									placeholder={_(msg`Email`)}
									required
									inputFieldSize="5xs"
									className="bg-[#ECECEC] p-3 text-[#A4A4A4]"
									ref={emailInput}
								/>
								{!login && (
									<InputField
										id="userNameInputFieldId"
										type="text"
										placeholder={_(msg`Username`)}
										required
										inputFieldSize="5xs"
										className="bg-[#ECECEC] p-3 text-[#A4A4A4]"
										ref={nameInput}
									/>
								)}
								<InputField
									id="passwordInputFieldId"
									type={"password"}
									placeholder={_(msg`Password`)}
									required
									minLength={passwordMinLength}
									inputFieldSize="5xs"
									className="bg-[#ECECEC] p-3 text-[#A4A4A4]"
									onChange={(event) => {
										event.target.setCustomValidity("");
										setPassword(event.target.value);
										setPasswordsMatching(
											repeatPasswordInput.current
												?.value === event.target.value
										);
										if (
											event.target.validity.tooShort ||
											event.target.validity.valueMissing
										) {
											event.target.setCustomValidity(
												t`Password must have at least ${passwordMinLength} characters`
											);
										}
									}}
									ref={passwordInput}
								/>
								{!login && (
									<>
										<InputField
											id="repeatPasswordInputFieldId"
											type={"password"}
											placeholder={_(
												msg`Repeat Password`
											)}
											required
											pattern={password}
											inputFieldSize="5xs"
											className="bg-[#ECECEC] p-3 text-[#A4A4A4]"
											ref={repeatPasswordInput}
											onChange={(event) => {
												event.target.setCustomValidity(
													""
												);
												if (
													passwordInput.current
														?.value !==
													event.target.value
												) {
													event.target.setCustomValidity(
														t`Passwords have to match`
													);
												}
												setPasswordsMatching(
													passwordInput.current
														?.value ===
														event.target.value
												);
											}}
										/>

										<div className="space-y-1 py-1 font-light text-[#676767]">
											<PasswordCheck
												text={_(
													msg`Minimum ${passwordMinLength} characters`
												)}
												valid={
													!passwordInput.current
														?.validity.tooShort &&
													!passwordInput.current
														?.validity.valueMissing
												}
											/>
											<PasswordCheck
												text={_(
													msg`Passwords have to match`
												)}
												valid={passwordsMatching}
											></PasswordCheck>
										</div>
									</>
								)}
								<Button
									id="loginSubmitButton"
									buttonIcon={
										<ArrowRight
											size={16}
											className="flex-none"
										/>
									}
									buttonTextSize="5xs"
									className="w-full justify-between bg-black p-3 text-white hover:scale-[1.02] hover:cursor-pointer hover:bg-neutral-900"
									onClick={() => {
										if (login) {
											loginLogic();
										} else {
											registerLogic();
										}
									}}
								>
									{login ? (
										<Trans>Sign in</Trans>
									) : (
										<Trans>Sign up</Trans>
									)}
								</Button>
							</form>
							<Text
								textSize="5xs"
								className="flex flex-row flex-wrap justify-center space-x-1 p-3 text-[#8E8E8E] md:p-5"
							>
								<span className="whitespace-nowrap">
									{login ? (
										<Trans>
											Don&apos;t have an account?
										</Trans>
									) : (
										<Trans>Already have an account?</Trans>
									)}
								</span>
								<button
									className="whitespace-nowrap text-black underline"
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
									{login ? (
										<Trans>Sign up now!</Trans>
									) : (
										<Trans>Sign in now!</Trans>
									)}
								</button>
							</Text>
						</div>
					</div>
				</div>
			</div>
		</>
	);

	async function loginLogic() {
		if (!form.current?.checkValidity()) {
			return;
		}
		const response = await fetch(`/api/login`, {
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
		const response = await fetch(`/api/register`, {
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

const PasswordCheck = ({ text, valid }: { text: string; valid: boolean }) => {
	return (
		<div className="flex flex-row items-center space-x-1">
			<Check
				size={12}
				className={valid ? "text-[#2DE100]" : "text-[#C2C2C2]"}
			/>
			<Text textSize="5xs">{text}</Text>
		</div>
	);
};
