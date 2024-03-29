import { Trans, msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { hasCookie } from "cookies-next";
import { GetStaticProps } from "next";
import { NextSeo } from "next-seo";

import Cards from "@/components/graphics/Cards";
import { Button } from "@/components/input/Button";
import { loadCatalog } from "@/pages/_app";

export const getStaticProps: GetStaticProps = async (ctx) => {
	const translation = await loadCatalog(ctx.locale!);
	return {
		props: {
			translation,
		},
	};
};

export default function Page() {

	const { _ } = useLingui();

	return (
		<div className="flex flex-1 justify-center overflow-hidden">
			<NextSeo
				canonical="https://kioku.dev"
				title={_(
					msg`Kioku | The free flashcard application that focusses on collaborative content creation!`
				)}
				description={_(
					msg`Kioku is a free flashcard application. You can create or import decks from Anki and edit your flashcards together. Learn your flashcards with our customized spaced repetition algorithm (SRS) and compare your statistics with your friends. Motivate each other and keep learning!`
				)}
				languageAlternates={[
					{ hrefLang: "en", href: "https://kioku.dev" },
					{ hrefLang: "de", href: "https://app.kioku.dev/de/home" },
				]}
				noindex={process.env.NEXT_PUBLIC_SEO != "True"}
				nofollow={process.env.NEXT_PUBLIC_SEO != "True"}
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
								href={
									hasCookie("access_token") ? "/" : "/login"
								}
								buttonStyle="primary"
								buttonTextSize="xs"
							>
								<Trans>Get started</Trans>
							</Button>
							<Button
								id="learnmoreButton"
								href="/features"
								buttonStyle="tertiary"
								buttonTextSize="xs"
								buttonIcon="ArrowRight"
								className="hover:space-x-2"
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
