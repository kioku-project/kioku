import { Trans, msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { GetStaticProps } from "next";
import { NextSeo } from "next-seo";
import { useRouter } from "next/router";
import { useEffect, useRef, useState } from "react";
import { ArrowRight, Check } from "react-feather";
import { toast } from "react-hot-toast";

import { Text } from "@/components/Text";
import { InputField } from "@/components/form/InputField";
import { Logo } from "@/components/graphics/Logo";
import { Button } from "@/components/input/Button";
import { loadCatalog } from "@/pages/_app";
import { submitForm } from "@/util/api";
import { loginRoute, reauthRoute, registerRoute } from "@/util/endpoints";

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
			const response = await fetch(reauthRoute);
			if (response.status === 200) {
				router.replace("/");
			}
		})();
	}, [router]);

	return (
		<>
			<NextSeo
				title={_(msg`Kioku | Login or register for Kioku!`)}
				description={_(
					msg`Register today and start using the free flashcard application together with your friends. Simply create new decks or import existing decks from Anki and collaborate in groups.`,
				)}
				languageAlternates={[
					{ hrefLang: "en", href: "https://app.kioku.dev/login" },
					{ hrefLang: "de", href: "https://app.kioku.dev/de/login" },
				]}
				noindex={process.env.NEXT_PUBLIC_SEO != "True"}
				nofollow={process.env.NEXT_PUBLIC_SEO != "True"}
				openGraph={{
					url: "https://app.kioku.dev/login",
				}}
			/>
			<div className="min-w-screen flex flex-1 bg-neutral-50">
				<div className="h-full w-full bg-gradient-to-bl from-[#FF83FA]/20 to-50%">
					<div className="flex h-full w-full items-center justify-center bg-gradient-to-tr from-[#83DAFF]/20 sm:p-5">
						<div className="flex h-full w-full flex-col items-center justify-center space-y-3 rounded-md bg-white p-8 shadow-[0_35px_60px_-15px_rgba(0,0,0,0.2)] sm:h-fit sm:w-80 md:px-7">
							<Logo
								href={"/home"}
								text={false}
								logoSize="w-32 sm:w-16 md:w-20"
								className="m-10 sm:m-3"
							/>
							<form
								className="w-full space-y-3 text-black"
								onSubmit={(event) => {
									event.preventDefault();
								}}
								ref={form}
							>
								<InputField
									id="emailInputFieldId"
									type="email"
									name="userEmail"
									placeholder={_(msg`Email`)}
									required
									className="bg-neutral-100 p-3 text-base sm:text-xs"
									ref={emailInput}
								/>
								{!login && (
									<InputField
										id="usernameInputFieldId"
										type="text"
										name="userName"
										placeholder={_(msg`Username`)}
										required
										minLength={3}
										className="bg-neutral-100 p-3 text-base sm:text-xs"
										ref={nameInput}
									/>
								)}
								<InputField
									id="passwordInputFieldId"
									type={"password"}
									name="userPassword"
									placeholder={_(msg`Password`)}
									inputFieldIconStyle="text-neutral-400"
									required
									minLength={passwordMinLength}
									className="bg-neutral-100 p-3 text-base sm:text-xs"
									onChange={(event) => {
										event.target.setCustomValidity("");
										setPassword(event.target.value);
										setPasswordsMatching(
											repeatPasswordInput.current
												?.value === event.target.value,
										);
										if (
											event.target.validity.tooShort ||
											event.target.validity.valueMissing
										) {
											event.target.setCustomValidity(
												t`Password must have at least ${passwordMinLength} characters`,
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
												msg`Repeat Password`,
											)}
											inputFieldIconStyle="text-neutral-400"
											required
											pattern={password}
											className="bg-neutral-100 p-3 text-base sm:text-xs"
											ref={repeatPasswordInput}
											onChange={(event) => {
												event.target.setCustomValidity(
													"",
												);
												if (
													passwordInput.current
														?.value !==
													event.target.value
												) {
													event.target.setCustomValidity(
														t`Passwords have to match`,
													);
												}
												setPasswordsMatching(
													passwordInput.current
														?.value ===
														event.target.value,
												);
											}}
										/>

										<div className="space-y-1 py-1 font-light text-neutral-500">
											<PasswordCheck
												text={_(
													msg`Minimum ${passwordMinLength} characters`,
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
													msg`Passwords have to match`,
												)}
												valid={passwordsMatching}
											></PasswordCheck>
										</div>
									</>
								)}
								<Button
									id="loginSubmitButtonId"
									buttonStyle="secondary"
									buttonSize="p-3"
									buttonIcon={
										<ArrowRight
											size={16}
											className="flex-none"
										/>
									}
									className="w-full justify-between text-base sm:text-xs"
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
								className="flex flex-row flex-wrap justify-center space-x-1 p-3 text-neutral-400 md:p-5"
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
									id="switchLoginButtonId"
									className="whitespace-nowrap text-black underline"
									onClick={() => {
										emailInput.current?.focus();
										setLogin((prev) => !prev);
									}}
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
		if (
			!form.current?.checkValidity() ||
			!emailInput.current ||
			!passwordInput.current
		) {
			return;
		}
		const toastID = toast.loading(t`Logging in...`, { id: "loginToast" });
		const response = await submitForm(loginRoute, [
			emailInput.current,
			passwordInput.current,
		]);
		if (response.ok) {
			toast.success(t`Logged in!`, { id: toastID });
			router.push("/");
		} else {
			const error = await response.text();
			toast.error(error, { id: toastID });
		}
	}

	async function registerLogic() {
		if (
			!form.current?.checkValidity() ||
			passwordInput.current?.value !==
				repeatPasswordInput.current?.value ||
			!emailInput.current ||
			!nameInput.current ||
			!passwordInput.current
		) {
			return;
		}
		const toastID = toast.loading(t`Registering account...`, {
			id: "registerToast",
		});
		const response = await submitForm(registerRoute, [
			emailInput.current,
			nameInput.current,
			passwordInput.current,
		]);
		if (response.ok) {
			toast.success(t`Account created!`, { id: toastID });
			setLogin(true);
			emailInput.current?.focus();
		} else {
			const error = await response.text();
			toast.error(error, { id: toastID });
		}
	}
}

const PasswordCheck = ({ text, valid }: { text: string; valid: boolean }) => {
	return (
		<div className="flex flex-row items-center space-x-1">
			<Check
				size={12}
				className={valid ? "text-[#2DE100]" : "text-neutral-400"}
			/>
			<Text textSize="5xs">{text}</Text>
		</div>
	);
};
