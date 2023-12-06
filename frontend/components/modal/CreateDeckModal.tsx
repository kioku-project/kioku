import { Trans, msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRef } from "react";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";

import { InputField } from "@/components/form/InputField";
import { Button } from "@/components/input/Button";
import { Modal, ModalProps } from "@/components/modal/modal";
import { Group as GroupType } from "@/types/Group";
import { postRequest } from "@/util/api";

interface CreateDeckModalProps {
	/**
	 * Group in which the deck will be created
	 */
	group: GroupType;
}

/**
 * Modal for creating decks
 */
export const CreateDeckModal = ({
	group,
	visible,
	setVisible,
	...props
}: CreateDeckModalProps & ModalProps) => {
	const { mutate } = useSWRConfig();
	const { _ } = useLingui();

	const deckNameInput = useRef<HTMLInputElement>(null);
	const deckDescriptionInput = useRef<HTMLInputElement>(null);

	return (
		<Modal
			header={_(msg`Create new Deck`)}
			visible={visible}
			setVisible={setVisible}
			{...props}
		>
			<form className="space-y-5">
				<div className="space-y-3">
					<InputField
						id="deckNameInputFieldId"
						type="text"
						label={_(msg`Deck name`)}
						inputFieldLabelStyle="text-gray-400"
						required
						placeholder={_(msg`Enter deck name`)}
						className="bg-gray-100 px-2 py-3"
						inputFieldSize="5xs"
						autoFocus
						ref={deckNameInput}
					></InputField>
					<InputField
						id="deckDescriptionInputFieldId"
						type="text"
						label={_(msg`Deck description`)}
						inputFieldLabelStyle="text-gray-400"
						placeholder={_(msg`Enter deck description`)}
						className="bg-gray-100 px-2 py-3"
						inputFieldSize="5xs"
						ref={deckDescriptionInput}
					></InputField>
				</div>
				<div className="flex flex-row justify-end space-x-1">
					<Button
						type="button"
						buttonStyle="cancel"
						buttonTextSize="3xs"
						onClick={() => {
							setVisible(false);
						}}
					>
						<Trans>Cancel</Trans>
					</Button>
					<Button
						type="submit"
						buttonStyle="secondary"
						buttonTextSize="3xs"
						onClick={() => {
							createDeck();
							setVisible(false);
						}}
					>
						<Trans>Create Deck</Trans>
					</Button>
				</div>
			</form>
		</Modal>
	);

	async function createDeck() {
		if (
			!deckNameInput.current ||
			deckNameInput.current?.validity.valueMissing
		) {
			return;
		}
		const response = await postRequest(
			`/api/groups/${group.groupID}/decks`,
			JSON.stringify({
				deckName: deckNameInput.current.value,
				deckDescription: deckDescriptionInput.current?.value,
			})
		);
		if (response?.ok) {
			toast.info(t`Deck created!`, { toastId: "newDeckToast" });
			mutate(`/api/groups/${group.groupID}/decks`);
		} else {
			toast.error("Error!", { toastId: "newDeckToast" });
		}
	}
};
