const colors = require('tailwindcss/colors');

module.exports = {
    important: false,
    content: ['./frontend/**/*.{html,js,ts,tsx}'],
    theme: {
        colors: {
            ...colors,
            primary: {
                600: '#3AC78B',
                500: '#62D2A2',
                300: '#B1E9D1',
            },
        },
    },
};
