import { Header } from "../../../components/navigation/Header";
import { Flashcard } from "../../../components/flashcard/Flashcard";
import Head from "next/head";
import Authenticated from "../../../components/accessControl/Authenticated";

export default function Page() {
	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<Authenticated>
				<div className="min-w-screen flex h-screen select-none flex-col bg-eggshell">
					<Header login={true}></Header>
					<Flashcard
						id="flashcardId"
						sides={[
							{
								header: "Front Header",
								description: "Front Description",
							},
							{
								header: "Middle Header",
								description: "Middle Description",
							},
							{
								header: "Back Header",
								description: "Back Description",
							},
						]}
					></Flashcard>
				</div>
			</Authenticated>
		</div>
	);
}
