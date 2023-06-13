import Image from "next/image";

interface SpinnerProps {
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying the Kioku Loading Spinner
 */
export default function LoadingSpinner({ className }: SpinnerProps) {
	return (
		<Image
			src="/loading_spinner.svg"
			width={0}
			height={0}
			alt="Loading spinner"
			className={`animate-spin ${className ?? ""}`}
			priority={true}
		></Image>
	);
}
