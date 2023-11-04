import { useRouter } from "next/router";
import { ReactElement } from "react";
import { Award, BarChart2, Cloud, Code, Compass, Users } from "react-feather";

import { Navbar } from "../components/navigation/Navbar";

export default function Page() {
	const router = useRouter();

	return (
		<div>
			<div className="min-w-screen flex h-screen flex-col bg-eggshell">
				<Navbar
					login={false}
					onClick={() => router.push("/home")}
				></Navbar>
				<div className="mx-auto flex flex-col justify-center p-5 text-base leading-7 md:w-2/3 md:p-10 md:text-center">
					<a
						className="text-lg font-semibold text-kiokuLightBlue hover:cursor-pointer"
						onClick={() => router.push("/login")}
					>
						Get started
					</a>
					<div className="mb-7 mt-1 text-3xl font-semibold leading-7 text-kiokuDarkBlue sm:text-4xl">
						Discover Kioku&apos;s awesome Features
					</div>
					<div className="text-lg leading-8 text-gray-600">
						Welcome to Kioku - the cloud native flashcard
						application that focuses on collaborative content
						creation. Our innovative platform allows you to organize
						your knowledge on flashcards and share them with other
						users to take your knowledge to the next level. Sign up
						for Kioku today and experience the future of learning.
					</div>
				</div>
				<div className="flex flex-col p-5 md:p-10">
					<div className="flex flex-col md:flex-row md:space-x-5">
						<FeatureCard
							header="Collaborative"
							description="Collaborate with your friends and fellow students in groups and work on shared decks. Learn together and motivate each other!"
							icon={
								<Users className="text-kiokuDarkBlue"></Users>
							}
						></FeatureCard>
						<FeatureCard
							header="Individual"
							description="Create and customize your own flashcards tailored to your needs and preferences. Set your own pace with our spaced repetition system to maximize your potential!"
							icon={
								<Compass className="text-kiokuDarkBlue"></Compass>
							}
						></FeatureCard>
					</div>
					<div className="flex flex-col md:flex-row md:space-x-5">
						<FeatureCard
							header="Compatible"
							description="Kioku is compatible with Anki, allowing you to import and export your existing decks into our application while taking advantage of Kioku's collaborative features!"
							icon={<Code className="text-kiokuDarkBlue"></Code>}
						></FeatureCard>
						<FeatureCard
							header="Informative"
							description="We provide you with detailed statistics and insights into your study progress. Identify areas of improvement to optimize your strategy for maximum effectiveness!"
							icon={
								<BarChart2 className="text-kiokuDarkBlue"></BarChart2>
							}
						></FeatureCard>
					</div>
					<div className="flex flex-col md:flex-row md:space-x-5">
						<FeatureCard
							header="Available"
							description="Access your flashcards everywhere and at any time. Switch seamlessly between multiple platforms and never miss a learning opportunity again!"
							icon={
								<Cloud className="text-kiokuDarkBlue"></Cloud>
							}
						></FeatureCard>
						<FeatureCard
							header="Entertaining"
							description="Achievements and leaderboards make learning more engaging and motivating. Kioku helps you to achieve better results and stay on track with your personal learning goals!"
							icon={
								<Award className="text-kiokuDarkBlue"></Award>
							}
						></FeatureCard>
					</div>
				</div>
			</div>
		</div>
	);
}

interface FeatureCardProps {
	/**
	 * Text that should be displayed in the Header
	 */
	header: string;
	/**
	 * Text that should be displayed as a description
	 */
	description: string;
	/**
	 * Icon
	 */
	icon: ReactElement;
	/**
	 * Additional classes
	 */
	className?: string;
}

const FeatureCard = ({
	header,
	description,
	icon,
	className = "",
}: FeatureCardProps) => {
	return (
		<div
			className={`flex w-full flex-row space-x-3 py-5 leading-7 md:w-1/2 ${className}`}
		>
			{icon}
			<div className="flex w-full flex-col">
				<p className="font-semibold text-kiokuDarkBlue">{header}</p>
				<p className="text-gray-400">{description}</p>
			</div>
		</div>
	);
};
