import { Header } from "../components/navigation/Header";
import { Card } from "../components/flashcard/Flashcard";

export default function Page() {
	return (
		<div className="min-w-screen flex h-screen select-none flex-col bg-eggshell">
			<Header login={true}></Header>
			<Card
				id="flashcardId"
				card={[
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
				cardsleft={16}
			></Card>
		</div>
	);
}
