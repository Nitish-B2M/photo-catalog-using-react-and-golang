/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'main': ["Poppins", "sans-serif"],
        'optional': ["Fredoka", "sans-serif"]
      },
      colors: {
	      'color1':"#c09569",
        'color2':"#DADED4",
	      'color3':"#292929",
	      'color4':"#ffffff",
	      'color5':"#c09569",
	      'color6':"#FF6F61",
      }
    },
  },
  plugins: [],
}

