import { useContext, useState } from "react";
import { Router, useRouter } from "next/router";
import Image from "next/image";
import { UserProvider, UserContext } from "../contexts/user";
import { Inter } from "next/font/google";
import Head from "next/head";

import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

import { FormInput } from "../components/form/FormInput";
import { FormButton } from "../components/form/FormButton";

const inter = Inter({
	weight: ["200", "400"],
	subsets: ["latin"],
});

export default function Page() {
	const [login, setLogin] = useState(true); // true = login, false = register
	const router = useRouter();
	const { username, setUsername } = useContext(UserContext);
	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<div className="min-w-screen flex min-h-screen select-none items-center p-5 md:p-10">
				<div className="flex h-fit w-full flex-col items-center rounded-3xl bg-blue-50 md:flex-row">
					<div className="m-5 mb-0 flex w-2/3 flex-col items-center rounded-l md:m-10 md:w-1/2 md:justify-center">
						<div className="relative my-5 h-[120px] w-full">
							<Image
								src="/kioku-logo.svg"
								alt="Kioku Logo"
								className="object-contain"
								fill
							/>
						</div>
						<p
							className={`${inter.className} text-clip indent-[0.5em] text-5xl font-extralight tracking-[0.5em] md:text-6xl`}
						>
							kioku
						</p>
					</div>
					{formView()}
				</div>
			</div>
			<ToastContainer
				position="bottom-center"
				autoClose={3000}
				hideProgressBar
				pauseOnFocusLoss
			/>
		</div>
	);

	function formView() {
		return (
			<div
				className={`flex h-fit w-5/6 flex-col items-center rounded-3xl bg-[#9EADC8] text-black ${inter.className} m-5 mt-10 md:mt-5 md:w-1/2 md:justify-center`}
			>
				<h1 className="mb-4 mt-5 text-2xl">
					{login ? "Login" : "Register"}
				</h1>
				{forms()}
			</div>
		);
	}

	function loginButton() {
		return (
			<>
				<FormButton
					id="login"
					value="Login"
					onClick={() => {
						if (login) {
							// TODO: login logic
							loginLogic().then();
						} else {
							registerLogic().then();
						}
					}}
				/>
				<span
					className="hover:cursor-pointer"
					onClick={() => setLogin(!login)}
				>
					or register
				</span>
			</>
		);
	}

	function registerButton() {
		return (
			<>
				<FormButton
					id="register"
					value="Register"
					onClick={() => {
						if (login) {
							// TODO: login logic
							loginLogic().then();
						} else {
							registerLogic().then();
						}
					}}
				/>
				<span
					className="hover:cursor-pointer"
					onClick={() => setLogin(!login)}
				>
					or login
				</span>
			</>
		);
	}

	function forms() {
		return (
			<form
				onSubmit={(e) => e.preventDefault()}
				className="flex w-full flex-col items-center"
			>
				<FormInput
					id="email"
					type="email"
					name="email"
					label="Email"
					additionalClasses="w-5/6 md:w-2/3"
				/>
				{!login && (
					<>
						<FormInput
							id="name"
							type="text"
							name="name"
							label="Name"
							additionalClasses="w-5/6 md:w-2/3"
						/>
					</>
				)}
				<FormInput
					id="password"
					type="password"
					name="password"
					label="Password"
					additionalClasses="w-5/6 md:w-2/3"
				/>
				{!login && (
					<>
						<FormInput
							id="passwordRepeat"
							type="password"
							name="passwordRepeat"
							label="Repeat Password"
							additionalClasses="w-5/6 md:w-2/3"
						/>
					</>
				)}
				<div className="mb-5 mt-5 flex items-center gap-2">
					{login ? loginButton() : registerButton()}
				</div>
			</form>
		);
	}

	async function loginLogic() {
		let email = document.querySelector("#email") as HTMLInputElement | null;
		let password = document.querySelector(
			"#password"
		) as HTMLInputElement | null;
		if (email?.value === "" || password?.value === "") {
			return;
		}
		let url =
			process.env.NEXT_PUBLIC_ENVIRONMENT !== "develop"
				? "/api/login"
				: "http://localhost:3002/api/login";
		const response = await fetch(url, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				email: email?.value,
				password: password?.value,
			}),
		});
		if (response.ok) {
			toast.info("Logged in!", { toastId: "accountToast" });
			const text = await response.text();
			setUsername(text);
			router.push("/");
		} else {
			toast.error("Wrong username or password", {
				toastId: "accountToast",
			});
		}
	}

	async function registerLogic() {
		let email = document.querySelector("#email") as HTMLInputElement | null;
		let name = document.querySelector("#name") as HTMLInputElement | null;
		let password = document.querySelector(
			"#password"
		) as HTMLInputElement | null;
		let passwordRepeat = document.querySelector(
			"#passwordRepeat"
		) as HTMLInputElement | null;
		if (
			email?.value === "" ||
			name?.value === "" ||
			password?.value === "" ||
			passwordRepeat?.value === ""
		) {
			return;
		}
		if (password?.value === passwordRepeat?.value) {
			let url =
				process.env.NEXT_PUBLIC_ENVIRONMENT !== "develop"
					? "/api/register"
					: "http://localhost:3001/api/register";
			const response = await fetch(url, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					email: email?.value,
					name: name?.value,
					password: password?.value,
				}),
			});
			if (response.ok) {
				toast.info("Account was created!", { toastId: "accountToast" });
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
