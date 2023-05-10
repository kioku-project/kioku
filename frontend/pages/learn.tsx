import { useRouter } from "next/router";
import { Header } from "../components/navigation/Header";
import { useState } from "react";
import { Card } from "../components/flashcard/Flashcard";

export default function Page() {
	const router = useRouter();
	const [cardsleft, setcardsleft] = useState(16);

	return (
		<div className="min-w-screen flex h-screen select-none flex-col bg-eggshell">
			<Header></Header>
			<Card
				id="flashcardId"
				card={{
					front: {
						header: "Front Header",
						description: "Front Description",
					},
					back: {
						header: "Back Header",
						description: "Back Description",
					},
				}}
				cardsleft={cardsleft}
				turned={false}
			></Card>
		</div>
	);
}
