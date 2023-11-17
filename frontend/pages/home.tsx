import { Trans, msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { hasCookie } from "cookies-next";
import { GetStaticProps } from "next";
import { NextSeo } from "next-seo";
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

	const { _ } = useLingui();

	return (
		<div className="flex flex-1 justify-center overflow-hidden">
			<NextSeo
				canonical={ router.locale === "en" ? "https://kioku.dev/" : "https://app.kioku.dev/de/home/" }
				title={_(
					msg`Kioku | Learn flashcards together with friends for free online!`
				)}
				description={_(
					msg`Kioku is a free flashcard application where you can create decks and learn together with friends. Import decks from anki or create new ones yourself and learn them together with friends. Compare your learning statistics with each other as motivatation to continue learning!`
				)}
				languageAlternates={[
					{ hrefLang: "en", href: "https://kioku.dev/" },
					{ hrefLang: "de", href: "https://app.kioku.dev/de/home" },
				]}
				noindex={!process.env.NEXT_PUBLIC_SEO}
				nofollow={!process.env.NEXT_PUBLIC_SEO}
				openGraph={{
					url: "https://kioku.dev",
				}}
			/>
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
