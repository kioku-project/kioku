import { Header } from "../components/navigation/Header";
import { Card } from "../components/flashcard/Flashcard";

export default function Page() {
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
				cardsleft={16}
				turned={false}
			></Card>
		</div>
	);
}
