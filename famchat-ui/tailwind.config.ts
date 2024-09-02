import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  darkMode: 'selector',
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      colors: {
        'primary-purple': '#3A215F',
        'secondary-purple': '#DEC6FA',
        'soft-pink': '#F6D6E7',
        'light-gray': '#E5E5E5',
      },
    },
    fontFamily: {
      sans: [
        '"Roboto", sans-serif',
      ],
    },
  },
  plugins: [],
};
export default config;
