/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js,templ}",
    "./static/**/*.{html,js}",
    "./static/css/**/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        clifford: '#67e8f9',
        border: '#06b6d4',
      },
      fontFamily: {
        'sans': ['Quicksand'],
      }
    }
  },
  extend: {
    spacing: {
      '8xl': '96rem',
      '9xl': '128rem',
    },
    borderRadius: {
      '4xl': '2rem',
    }
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

