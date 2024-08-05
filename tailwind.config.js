/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'internal/templates/**/*.templ',
  ],
  theme: {
    extend: {
      colors: {
      tLavender: "#7469B6",
      tLilac: "#AD88C6",
      tPink: "#E1AFD1", 
      tPalePink: "#FFE6E6",
      tGothicGreen: "#1A3636", 
      tSage: "#40534C",
      tOlive: "#677D6A",
      tTan: "#D6BD98",
      },
    },
  },
  plugins: [],
}

