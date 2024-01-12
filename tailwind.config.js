/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './view/**/*.html',
    './view/**/*.templ',
    './view/**/*.go',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}

