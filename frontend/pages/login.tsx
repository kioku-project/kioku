import { Inter } from "next/font/google";
import Image from "next/image";
import { useRouter } from "next/router";
import { useRef, useState } from "react";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

import { Text } from "../components/Text";
import { FormButton } from "../components/form/FormButton";
import { InputField } from "../components/form/InputField";

const inter = Inter({
	weight: ["200", "400"],
	subsets: ["latin"],
});

export default function Page() {
	const router = useRouter();
	const [login, setLogin] = useState(true); // true = login, false = register
	const form = useRef<HTMLFormElement>(null);
	const emailInput = useRef<HTMLInputElement>(null);
	const nameInput = useRef<HTMLInputElement>(null);
	const passwordInput = useRef<HTMLInputElement>(null);
	const repeatPasswordInput = useRef<HTMLInputElement>(null);
	const [password, setPassword] = useState("");

	return (
		<div>
			<div className="min-w-screen flex min-h-screen items-center justify-center sm:p-5 md:p-10">
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
							size="md"
							className="text-center font-bold leading-9 tracking-tight text-kiokuDarkBlue"
						>
							{login
								? "Sign in to your account"
								: "Create an account"}
						</Text>
						{forms()}

						<Text size="3xs" className="text-center text-gray-500">
							{login
								? "Not registered? "
								: "Already registered? "}
							<a
								className="whitespace-nowrap font-semibold text-kiokuDarkBlue transition hover:cursor-pointer hover:text-eggshell"
								onClick={() => {
									emailInput.current?.focus();
									setLogin((prev) => !prev);
								}}
							>
								{login ? "Create an account" : "Sign in"}
							</a>
						</Text>
					</div>
				</div>
			</div>
		</div>
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
					label="Email"
					required={true}
					size="xs"
					ref={emailInput}
				/>
				{!login && (
					<InputField
						id="name"
						type="text"
						name="name"
						label="Name"
						required={true}
						size="xs"
						ref={nameInput}
					/>
				)}
				<InputField
					id="password"
					type="password"
					name="password"
					label="Password"
					required={true}
					minLength={3}
					size="xs"
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
						label="Repeat Password"
						tooltipMessage="Passwords have to match."
						required={true}
						minLength={3}
						pattern={password}
						size="xs"
						ref={repeatPasswordInput}
					/>
				)}

				<FormButton
					id={login ? "login" : "register"}
					value={login ? "Login" : "Register"}
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
			toast.info("Logged in!", { toastId: "accountToast" });
			router.push("/");
		} else {
			toast.error("Wrong username or password", {
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
			toast.info("Account created!", { toastId: "accountToast" });
			setLogin(true);
			emailInput.current?.focus();
		} else {
			toast.error("Account already exists!", {
				toastId: "accountToast",
			});
		}
	}
}
