import { useRef, useState } from "react";
import { useRouter } from "next/router";
import Image from "next/image";
import { Inter } from "next/font/google";
import Head from "next/head";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { InputField } from "../components/form/InputField";
import { FormButton } from "../components/form/FormButton";

const inter = Inter({
	weight: ["200", "400"],
	subsets: ["latin"],
});

export default function Page() {
	const router = useRouter();
	const [login, setLogin] = useState(true); // true = login, false = register
	const emailInput = useRef<HTMLInputElement>(null);
	const nameInput = useRef<HTMLInputElement>(null);
	const passwordInput = useRef<HTMLInputElement>(null);
	const repeatPasswordInput = useRef<HTMLInputElement>(null);

	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<div className="min-w-screen flex min-h-screen select-none items-center justify-center sm:p-5 md:p-10">
				<div className="flex min-h-screen w-full flex-col items-center justify-evenly bg-blue-100 p-10 align-middle sm:min-h-fit sm:rounded-3xl md:flex-row xl:w-5/6">
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
						className={`flex w-full flex-col items-center rounded-2xl bg-kiokuLightBlue p-5 sm:w-5/6 md:w-1/2 lg:w-1/3 ${inter.className}`}
					>
						<h2 className="text-center text-2xl font-bold leading-9 tracking-tight text-kiokuDarkBlue">
							{login
								? "Sign in to your account"
								: "Create an account"}
						</h2>
						{forms()}
						<p className="text-center text-sm text-gray-500">
							{login
								? "Not registered? "
								: "Already registered? "}
							<a
								className="whitespace-nowrap font-semibold text-kiokuDarkBlue transition hover:cursor-pointer hover:text-eggshell"
								onClick={() => setLogin(!login)}
							>
								{login ? "Create an account" : "Sign in"}
							</a>
						</p>
					</div>
				</div>
			</div>
		</div>
	);

	function forms() {
		return (
			<form
				onSubmit={(e) => e.preventDefault()}
				className="my-5 flex w-5/6 flex-col items-center space-y-4"
			>
				<InputField
					id="email"
					type="email"
					name="email"
					label="Email"
					style="primary"
					className="sm:text-sm"
					ref={emailInput}
				/>
				{!login && (
					<InputField
						id="name"
						type="text"
						name="name"
						label="Name"
						style="primary"
						className="sm:text-sm"
						ref={nameInput}
					/>
				)}
				<InputField
					id="password"
					type="password"
					name="password"
					label="Password"
					style="primary"
					className="sm:text-sm"
					ref={passwordInput}
				/>
				{!login && (
					<InputField
						id="passwordRepeat"
						type="password"
						name="passwordRepeat"
						label="Repeat Password"
						style="primary"
						className="sm:text-sm"
						ref={repeatPasswordInput}
					/>
				)}

				<FormButton
					id={login ? "login" : "register"}
					value={login ? "Login" : "Register"}
					style="primary"
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
		if (
			emailInput.current?.value === "" ||
			passwordInput.current?.value === ""
		) {
			return;
		}
		let url = "/api/login";
		const response = await fetch(url, {
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
			emailInput.current?.value === "" ||
			nameInput.current?.value === "" ||
			passwordInput.current?.value === "" ||
			repeatPasswordInput.current?.value === ""
		) {
			return;
		}
		if (
			passwordInput.current?.value === repeatPasswordInput.current?.value
		) {
			let url = "/api/register";
			const response = await fetch(url, {
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
			} else {
				toast.error("Account already exists!", {
					toastId: "accountToast",
				});
			}
		} else {
			toast.error("The passwords do not match!", {
				toastId: "passwordToast",
			});
		}
	}
}
