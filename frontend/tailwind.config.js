module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        black: {
          20: "rgba(0, 0, 0, 0.2)",
          40: "rgba(0, 0, 0, 0.4)",
          60: "rgba(0, 0, 0, 0.6)",
        },
      },
      maxHeight: {
        0: "0",
        "1/4": "25%",
        "1/2": "50%",
        "3/4": "75%",
        full: "100%",
      },
      margin: {
        "minus-2": "-2rem",
      },
    },
  },
  plugins: [require("flowbite/plugin")],
};
