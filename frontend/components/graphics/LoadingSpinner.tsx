import Image from "next/image";
import React, { useEffect, useState } from "react";

interface SpinnerProps {
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Delay in ms before the spinner is shown
	 */
	delay?: number;
}

/**
 * UI component for displaying the Kioku Loading Spinner
 */
export default function LoadingSpinner({
	className = "",
	delay = 1000,
}: SpinnerProps) {
	const [delayed, setDelayed] = useState(true);
	useEffect(() => {
		const timeout = setTimeout(() => setDelayed(false), delay);
		return () => clearTimeout(timeout);
	}, [delay]);
	return delayed ? (
		<></>
	) : (
		<Image
			src="/loading_spinner.svg"
			width={0}
			height={0}
			alt="Loading spinner"
			className={`animate-spin ${className}`}
			priority={true}
		/>
	);
}
