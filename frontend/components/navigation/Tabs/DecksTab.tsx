import { useState } from "react";
import { PlusSquare } from "react-feather";

import DeckList from "@/components/deck/DeckList";
import { CreateDeckModal } from "@/components/modal/CreateDeckModal";
import { Group as GroupType } from "@/types/Group";
import { GroupRole } from "@/types/GroupRole";
import { useDecks } from "@/util/swr";

interface DecksTabProps {
	/**
	 * Group entity
	 */
	group: GroupType;
}

/**
 * UI component for the DecksTab
 */
export const DecksTab = ({ group }: DecksTabProps) => {
	const { decks } = useDecks(group.groupID);

	const [showModal, setShowModal] = useState(false);

	const hasWrite =
		group.groupRole && GroupRole[group.groupRole] >= GroupRole.WRITE;

	return (
		<div>
			<CreateDeckModal
				group={group}
				visible={showModal}
				setVisible={setShowModal}
			/>
			<div className="space-y-3">
				<div className="flex w-full items-center justify-end rounded-md bg-neutral-100 px-4 py-3">
					<PlusSquare
						className={`${
							hasWrite
								? "text-kiokuDarkBlue hover:scale-110 hover:cursor-pointer"
								: "text-gray-400 hover:cursor-not-allowed"
						} transition`}
						onClick={() => {
							if (hasWrite) {
								setShowModal(true);
							}
						}}
					/>
				</div>
				<DeckList decks={decks} />
			</div>
		</div>
	);
};
