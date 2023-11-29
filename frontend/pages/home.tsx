import { Trans } from "@lingui/macro";
import { hasCookie } from "cookies-next";
import { GetStaticProps } from "next";
import Head from "next/head";
import { useRouter } from "next/router";

import Cards from "../components/graphics/Cards";
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

	return (
		<div className="flex flex-1 justify-center overflow-hidden">
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
				<link rel="alternate" hrefLang="en" href="https://kioku.dev" />
				<link
					rel="alternate"
					hrefLang="de"
					href="https://app.kioku.dev/de/home"
				/>
			</Head>
			<div className="flex min-w-full flex-col bg-eggshell">
				<div className="flex grow flex-row justify-start">
					<div className="flex w-full flex-col justify-center space-y-3 p-5 md:w-2/3 md:space-y-5 md:p-10">
						<p className="text-2xl font-bold text-kiokuDarkBlue sm:text-3xl md:text-4xl lg:text-5xl">
							<Trans>
								We&apos;re changing the way people learn.
							</Trans>
						</p>
						<p className="text-sm font-semibold text-gray-400 sm:text-base">
							<Trans>
								Start your learning journey today with Kioku -
								the cloud native flashcard application that
								focusses on collaborative content creation
							</Trans>
						</p>
						<div className="flex flex-row space-x-3 md:space-x-5">
							<Button
								id="getstartedButton"
								buttonStyle="primary"
								buttonSize="md"
								buttonTextSize="xs"
								onClick={() =>
									hasCookie("access_token")
										? router.push("/")
										: router.push("/login")
								}
							>
								<Trans>Get started</Trans>
							</Button>
							<Button
								id="learnmoreButton"
								buttonStyle="secondary"
								buttonSize="md"
								buttonTextSize="xs"
								buttonIcon="ArrowRight"
								onClick={() => {
									router.push("/features");
								}}
							>
								<Trans>Learn more</Trans>
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
