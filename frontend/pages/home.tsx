import { NextSeo } from "next-seo";
import { useRouter } from "next/router";
import { ArrowRight } from "react-feather";

import Cards from "../components/graphics/Cards";
import { Button } from "../components/input/Button";
import { Navbar } from "../components/navigation/Navbar";
import Head from "next/head";

export default function Page() {
	const router = useRouter();

	return (
		<div>
			<NextSeo noindex={false} nofollow={false} />
			<Head>
				<link rel="canonical" href="https://kioku.dev/" />
			</Head>
			<div className="min-w-screen flex h-screen flex-col bg-eggshell">
				<Navbar
					login={false}
					onClick={() => router.push("/home")}
				></Navbar>
				<div className="flex grow flex-row overflow-hidden">
					<div className="flex w-full flex-col justify-center space-y-3 p-5 md:w-2/3 md:space-y-5 md:p-10">
						<h1 className="text-2xl font-bold text-kiokuDarkBlue sm:text-3xl md:text-4xl lg:text-5xl">
							We&apos;re changing the way people learn.
						</h1>
						<p className="text-sm font-semibold text-gray-400 sm:text-base">
							Start your learning journey today with Kioku - the
							cloud native flashcard application that focusses on
							collaborative content creation
						</p>
						<div className="flex flex-row space-x-3 md:space-x-5">
							<Button
								id="getstartedButton"
								onClick={() => {
									router.push("/");
								}}
							>
								Get started
							</Button>
							<Button
								id="lernmoreButton"
								style="secondary"
								onClick={() => {
									router.push("/features");
								}}
							>
								Learn more
								<ArrowRight className="ml-1 h-2/3"></ArrowRight>
							</Button>
						</div>
					</div>
					<div className="my-auto hidden md:block md:w-1/3">
						<Cards />
					</div>
				</div>
			</div>
		</div>
	);
}
