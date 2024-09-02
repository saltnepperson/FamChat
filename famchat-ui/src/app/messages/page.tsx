import InputTextBox from '@/components/InputTextBox';
import Image from 'next/image';

export default function Page() {
  return (
    <div className="bg-gray-100 h-full w-full flex justify-center items-center">
      <div className="w-full max-w-md p-4">
        <div className="flex items-start mb-4">
          <Image
            src="https://doodleipsum.com/700?i=b87ddbb3499fe9e609e2d0e6aae01b9c" // Replace with actual image path
            alt="Daniel Garcia"
            width={40}
            height={40}
            className="rounded-full"
          />
          <div className="ml-4">
            <div className="font-bold text-sm text-gray-800">Wintson</div>
            <div className="bg-purple-100 text-sm text-gray-700 p-2 rounded-lg">
              Mommy... Can you come here please?
            </div>
            <div className="text-xs text-gray-400 mt-1">12:04</div>
          </div>
        </div>

        {/* Add more messages here as needed */}

        <InputTextBox />
      </div>
    </div>
  );
}
