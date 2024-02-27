import { Trans, msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRef } from "react";

import { InputField } from "@/components/form/InputField";
import { Button } from "@/components/input/Button";
import { Modal, ModalProps } from "@/components/modal/modal";
import { Deck as DeckType } from "@/types/Deck";
import { createCard } from "@/util/api";

interface CreateCardModalProps {
	/**
	 * Deck in which the card will be created
	 */
	deck: DeckType;
}

/**
 * Modal for creating cards
 */
export const CreateCardModal = ({
	deck,
	visible,
	setVisible,
	...props
}: CreateCardModalProps & ModalProps) => {
	const { _ } = useLingui();

	const cardFrontInput = useRef<HTMLInputElement>(null);
	const cardBackInput = useRef<HTMLInputElement>(null);

	return (
		<Modal
			header={_(msg`Create new Card`)}
			visible={visible}
			setVisible={setVisible}
			{...props}
		>
			<form className="space-y-5">
				<div className="space-y-3">
					<InputField
						id="cardFrontInputFieldId"
						type="text"
						label={_(msg`Card front`)}
						inputFieldLabelStyle="text-gray-400"
						required
						placeholder={_(msg`Enter card front`)}
						className="bg-gray-100 px-2 py-3"
						inputFieldSize="5xs"
						autoFocus
						ref={cardFrontInput}
					></InputField>
					<InputField
						id="cardBackInputFieldId"
						type="text"
						label={_(msg`Card back`)}
						inputFieldLabelStyle="text-gray-400"
						placeholder={_(msg`Enter card back`)}
						className="bg-gray-100 px-2 py-3"
						inputFieldSize="5xs"
						ref={cardBackInput}
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
							cardFrontInput.current &&
								cardBackInput.current &&
								createCard(
									deck.deckID,
									cardFrontInput.current,
									cardBackInput.current
								);
							setVisible(false);
						}}
					>
						<Trans>Create Card</Trans>
					</Button>
				</div>
			</form>
		</Modal>
	);
};
