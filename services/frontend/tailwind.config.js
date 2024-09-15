/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{js,jsx,ts,tsx}"],
    theme: {
        extend: {
            fontFamily: {
                vazir: "Vazir",
                openSans: "OpenSans",
            },
        },
    },
    plugins: [require("rippleui")],
    /** @type {import('rippleui').Config} */
    rippleui: {
        themes: [
            {
                themeName: "light",
                colorScheme: "light",
                colors: {
                    primary: "#0398fc",
                    backgroundPrimary: "#f7f7f7",
                },
            },
            {
                themeName: "dark",
                colorScheme: "dark",
                colors: {
                    primary: "#0398fc",
                    backgroundPrimary: "#141414",
                },
            },
        ],
    },
};
