import type { Metadata } from "next";
import { primary_font } from "@/app/ui/fonts";
import NavBar from "@/components/navbar";
import NavBarIcon from "@/components/navbaricon";
import "@/app/ui/global.css";

export const metadata: Metadata = {
  title: "FamChat",
  description: "A simple, secure family-friendly chat app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="bg-white">
      <body className={`${primary_font.className} antialiased dark:bg-black h-full`}>
        <div className="flex h-full">
          <div className="basis-1/6 drop-shadow-md bg-primary-purple flex flex-col items-center py-6">
            {/* Ensure the NavBarIcon and NavBar are vertically aligned */}
            <NavBarIcon />
            <NavBar />
            {/* Current user profile goes here */}
          </div>
          <div className="basis-5/6 p-4">
            {children}
          </div>
        </div>
      </body>
    </html>
  );
}
