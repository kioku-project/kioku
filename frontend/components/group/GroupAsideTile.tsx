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
		<div id={id} className="m-4 flex items-center gap-4">
			<div className="relative">
				<div className="h-8 w-8 rounded-md bg-white" />
				{count > 0 && (
					<div className="absolute bottom-[-0.5rem] right-[-0.5rem] flex h-4 items-center justify-center rounded-sm bg-red px-1">
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
