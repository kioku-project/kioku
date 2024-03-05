import { Trans, msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRef } from "react";

import { InputField } from "@/components/form/InputField";
import { Button } from "@/components/input/Button";
import { Modal, ModalProps } from "@/components/modal/modal";
import { createGroup } from "@/util/api";

/**
 * Modal for creating groups
 */
export const CreateGroupModal = ({
	visible,
	setVisible,
	...props
}: ModalProps) => {
	const { _ } = useLingui();

	const groupNameInput = useRef<HTMLInputElement>(null);
	const groupDescriptionInput = useRef<HTMLInputElement>(null);

	return (
		<Modal
			header={_(msg`Create new Group`)}
			visible={visible}
			setVisible={setVisible}
			{...props}
		>
			<form
				onSubmit={(event) => {
					event.preventDefault();
					event.currentTarget.checkValidity() &&
						groupNameInput.current &&
						groupDescriptionInput.current &&
						createGroup([
							groupNameInput.current,
							groupDescriptionInput.current,
						]);
					setVisible(false);
				}}
				className="space-y-5"
			>
				<div className="space-y-3">
					<InputField
						id="groupNameInputFieldId"
						type="text"
						name="groupName"
						label={_(msg`Group Name`)}
						inputFieldLabelStyle="text-gray-400"
						required
						placeholder={_(msg`Enter group name`)}
						className="bg-gray-100 px-2 py-3"
						inputFieldSize="5xs"
						autoFocus
						ref={groupNameInput}
					></InputField>
					<InputField
						id="groupDescriptionInputFieldId"
						type="text"
						name="groupDescription"
						label={_(msg`Group Description`)}
						inputFieldLabelStyle="text-gray-400"
						placeholder={_(msg`Enter group description`)}
						className="bg-gray-100 px-2 py-3"
						inputFieldSize="5xs"
						ref={groupDescriptionInput}
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
					>
						<Trans>Create Group</Trans>
					</Button>
				</div>
			</form>
		</Modal>
	);
};
