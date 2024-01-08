/** @type {import('tailwindcss').Config} */
module.exports = {
  // mode: 'production',
  content: [
    "./view/**/*.{html,js,templ}",
    "./static/**/*.{html,js,css}"],
  theme: {
    extend: {
      colors: {
        clifford: '#67e8f9',
        border: '#06b6d4',
      },
      fontFamily: {
        'sans': ['Quicksand'],
      },
      dropShadow: {
        '2xl': '0 8px 8px rgba(255, 255, 255, 0.10)',
        '3xl': [
          '0 15px 15px rgba(255, 255, 255, 0.10)'
        ]
      },
    }
  },
  plugins: [],
}

