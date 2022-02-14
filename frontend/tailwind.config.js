module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      minWidth: {
        11: "2.75rem",
      },
      colors: {
        black: {
          20: "rgba(0, 0, 0, 0.2)",
          40: "rgba(0, 0, 0, 0.4)",
          60: "rgba(0, 0, 0, 0.6)",
        },
      },

      spacing: {
        0: "0",
        "1/4": "25%",
        "1/2": "50%",
        "3/4": "75%",
        full: "100%",
        112: "28rem",
      },
      margin: {
        "minus-2": "-2rem",
      },
    },
  },
  plugins: [require("flowbite/plugin")],
};
