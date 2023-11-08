import { Trans } from "@lingui/macro";
import { hasCookie } from "cookies-next";
import { GetStaticProps } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import { ArrowRight } from "react-feather";

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
								onClick={() =>
									hasCookie("access_token")
										? router.push("/")
										: router.push("/login")
								}
							>
								<Trans>Get started</Trans>
							</Button>
							<Button
								id="lernmoreButton"
								buttonStyle="secondary"
								onClick={() => {
									router.push("/features");
								}}
							>
								<Trans>Learn more</Trans>
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
