import { useRouter } from "next/router";
import { Button } from "../components/input/Button";
import Cards from "../components/graphics/Cards";
import { Header } from "../components/navigation/Header";

export default function Page() {
	const router = useRouter();

	return (
		<div className="min-w-screen flex h-screen select-none flex-col bg-eggshell">
			<Header></Header>
			<div className="flex grow flex-row overflow-hidden">
				<div className="flex w-full flex-col justify-center space-y-3 p-5 md:w-2/3 md:space-y-5 md:p-10">
					<p className="text-2xl font-bold text-darkblue sm:text-3xl md:text-4xl lg:text-5xl">
						We&apos;re changing the way people learn.
					</p>
					<p className="text-sm font-semibold text-gray-400 sm:text-base">
						Start your learning journey today with Kioku - the cloud
						native flashcard application that focusses on
						collaborative content creation
					</p>
					<div className="flex flex-row space-x-3 md:space-x-5">
						<Button
							value="Get started"
							id="getstartedButton"
							onClick={() => {
								router.push("/login");
							}}
						></Button>
						<Button
							value="Learn more &#129030;"
							id="lernmoreButton"
							style="secondary"
							onClick={() => {
								router.push("/learn");
							}}
						></Button>
					</div>
				</div>
				<div className="my-auto hidden md:block md:w-1/3">
					<Cards />
				</div>
			</div>
		</div>
	);
}
