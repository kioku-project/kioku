import { msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRouter } from "next/router";
import { useRef } from "react";
import { PlusSquare } from "react-feather";
import { toast } from "react-toastify";
import { mutate } from "swr";

import DeckList from "@/components/deck/DeckList";
import { InputField } from "@/components/form/InputField";
import { Group as GroupType } from "@/types/Group";
import { postRequest } from "@/util/api";

interface GroupsTabProps {
	/**
	 * groups
	 */
	groups: GroupType[];
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the GroupsTab
 */
export const GroupsTab = ({ groups, className = "" }: GroupsTabProps) => {
	const router = useRouter();
	const { _ } = useLingui();

	const groupNameInput = useRef<HTMLInputElement>(null);

	return (
		<div className={`space-y-3 ${className}`}>
			<div className="flex w-full items-center justify-between rounded-md bg-neutral-100 px-4 py-3">
				<InputField
					id={`groupNameInput`}
					placeholder={_(msg`Create new Group`)}
					inputFieldSize="xs"
					className="w-full bg-transparent font-medium text-kiokuDarkBlue outline-none"
					onKeyUp={(event) => {
						if (event.key === "Enter") {
							createGroup()
								.then((result) => {})
								.catch((error) => {});
						}
					}}
					ref={groupNameInput}
				/>
				<PlusSquare
					className="text-kiokuDarkBlue transition hover:scale-110 hover:cursor-pointer"
					onClick={createGroup}
				/>
			</div>
			<div className={`space-y-3 ${className}`}>
				{groups
					?.filter((group: GroupType) => !group.isDefault)
					.map((group: GroupType) => {
						return (
							<div
								key={group.groupID}
								className="hover:cursor-pointer"
								onClick={() =>
									router.push(`/group/${group.groupID}`)
								}
								onKeyUp={(event) => {
									if (event.key === "Enter") {
										event.target.dispatchEvent(
											new Event("click", {
												bubbles: true,
											})
										);
									}
								}}
							>
								<DeckList
									header={group.groupName}
									key={group.groupID}
								/>
							</div>
						);
					})}
				<DeckList />
			</div>
		</div>
	);

	async function createGroup() {
		if (!groupNameInput.current?.value) {
			groupNameInput.current?.focus();
			return;
		}
		const response = await postRequest(
			`/api/groups`,
			JSON.stringify({ groupName: groupNameInput.current.value })
		);
		if (response?.ok) {
			groupNameInput.current.value = "";
			toast.info(t`Group created!`, { toastId: "newGroupToast" });
		} else {
			toast.error("Error!", { toastId: "newGroupToast" });
		}
		mutate(`/api/groups`);
	}
};
