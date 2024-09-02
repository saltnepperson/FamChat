import Link from "next/link";
import { EnvelopeIcon, HomeIcon } from "@heroicons/react/24/solid";

export default function NavBar() {
    return (
        <nav className="h-dvh p-14 text-lg">
            <span className="text-white/50 font-bold tracking-wide leading-loose">
                <Link href="/" className="hover:text-white flex items-center">
                <HomeIcon className="h-6 w-6 mr-2"></HomeIcon>
                Home
                </Link>
                <br/>
                <Link href="/messages" className="hover:text-white flex items-center">
                    <EnvelopeIcon className="h-6 w-6 mr-2" />
                    Messages
                </Link>
            </span>
        </nav>
    );
};
