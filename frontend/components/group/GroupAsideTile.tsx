interface GroupAsideTileProps {
	/**
	 * unique identifier
	 */
	id?: string;
	/**
	 * Group name
	 */
	name: string;
	/**
	 * Count of due cards inside Group
	 */
	count: number;
}

/**
 * UI component for a group in the side bar
 */
export default function GroupAsideTile({
	id,
	name,
	count,
}: GroupAsideTileProps) {
	return (
		<div id={id} className="flex m-4 gap-4 items-center">
			<div className="relative">
				<div className="bg-white w-8 h-8 rounded-md" />
				{count > 0 && (
					<div className="bg-red-600 rounded-sm h-4 px-1 absolute bottom-[-0.5rem] right-[-0.5rem] flex justify-center items-center">
						<span className="text-xs font-bold text-white">
							{count}
						</span>
					</div>
				)}
			</div>
			<span>{name}</span>
		</div>
	);
}
