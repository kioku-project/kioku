import Link from "next/link";
import { Menu } from "react-feather";

interface NavbarProps {
	/**
	 * unique identifier
	 */
	id?: string;
	/**
	 * A callback for a press on the burger menu
	 */
	toggleMenuCallback?: () => void;
}

/**
 * UI component for the top navigation bar
 */
export default function Navbar({ toggleMenuCallback }: NavbarProps) {
	return (
		<nav className="flex justify-between items-center bg-[#979797] px-4">
			<div className="flex gap-4 items-center">
				<Menu
					onClick={toggleMenuCallback}
					className="hover:cursor-pointer"
				/>
				<Link className="text-xl" href="/">
					<b>Kioku</b>
				</Link>
			</div>
			{/* Placeholder for profile settings */}
			<div className="bg-black m-2 w-12 h-12 rounded-3xl" />
		</nav>
	);
}
