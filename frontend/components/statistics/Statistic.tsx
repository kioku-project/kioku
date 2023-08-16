import React from "react";
import {
	ArrowDown,
	ArrowDownRight,
	ArrowRight,
	ArrowUp,
	ArrowUpRight,
} from "react-feather";

interface StatisticProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Header
	 */
	header: string;
	/**
	 * current Value
	 */
	value: string;
	/**
	 * seperator displayed between value and reference
	 */
	seperator?: string;
	/**
	 * reference Value
	 */
	reference?: string;
	/**
	 * Change
	 */
	change: number;
	/**
	 * Additional classes
	 */
	className?: string;
}

function getArrow(change: number) {
	const className = "w-4";
	if (change > 10) {
		return <ArrowUp className={className}></ArrowUp>;
	} else if (change > 0) {
		return <ArrowUpRight className={className}></ArrowUpRight>;
	} else if (change == 0) {
		return <ArrowRight className={className}></ArrowRight>;
	} else if (change > -10) {
		return <ArrowDownRight className={className}></ArrowDownRight>;
	} else {
		return <ArrowDown className={className}></ArrowDown>;
	}
}

function getStyle(change: number) {
	if (change > 0) {
		return "bg-green-300 text-green-800";
	} else if (change == 0) {
		return "bg-gray-300 text-gray-900";
	} else {
		return "bg-red-300 text-red-800";
	}
}

/**
 * UI component for basic statistics displaying a value and a trend
 */
export const Statistic = ({
	id,
	header,
	value,
	seperator,
	reference,
	change,
	className = "",
}: StatisticProps) => {
	return (
		<div id={id} className={`flex w-full flex-col px-5 py-3 ${className}`}>
			<div className="font-bold text-kiokuLightBlue">{header}</div>
			<div className="flex flex-row items-end justify-between space-x-5">
				<div className="flex flex-row space-x-1">
					<div className="flex items-end text-3xl font-black text-kiokuDarkBlue">
						{value}
					</div>
					<div className="flex items-end text-kiokuLightBlue">
						{`${seperator ?? ""} ${reference ?? ""}`}
					</div>
				</div>
				<div
					className={`flex h-fit flex-row items-center space-x-1 rounded-2xl  px-2 font-semibold ${getStyle(
						change
					)}`}
				>
					{getArrow(change)}
					<div className="flex items-center">{`${change}%`}</div>
				</div>
			</div>
		</div>
	);
};
