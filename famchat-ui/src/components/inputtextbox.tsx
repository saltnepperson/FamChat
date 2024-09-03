import { PaperAirplaneIcon, LinkIcon } from '@heroicons/react/24/solid';

export default function InputTextBox() {
    return (
        <div className="flex items-center bg-white rounded-lg px-6 py-4">
            <button className="ml-2 text-purple-500">
                <LinkIcon className="h-5 w-5 mr-5" />
            </button>
            <input
                type="text"
                className="flex-1 bg-transparent border-none focus:outline-none text-sm placeholder-white-500"
                placeholder="Type A Message..."
            />
            <button className="ml-2 text-purple-500">
                <PaperAirplaneIcon className="h-5 w-5" />
            </button>
        </div>
    );
}
