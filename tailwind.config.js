/** @type {import('tailwindcss').Config} */
  module.exports = {
    content: ["./app/views/**/*.templ", "./components/**/*.templ"],
    darkMode: "class",
    theme: {
      colors: {
        white: "hsla(0deg, 0%, 100%, <alpha-value>)",
        black: "hsla(0deg, 0%, 0%, <alpha-value>)",
        transparent: "transparent",
        current: "currentColor",
        base: {
          background: "hsla(var(--background), <alpha-value>)",
          foreground: "hsla(var(--foreground), <alpha-value>)",
        },
        muted: {
          background: "hsla(var(--muted-background), <alpha-value>)",
          foreground: "hsla(var(--muted-foreground), <alpha-value>)",
        },
        primary: {
          DEFAULT: "hsla(var(--primary), <alpha-value>)",
          hover: "hsla(var(--primary-hover), <alpha-value>)",
          muted: "hsla(var(--primary-muted), <alpha-value>)",
          foreground: "hsla(var(--primary-foreground), <alpha-value>)",
        },
        danger: {
          DEFAULT: "hsla(var(--danger), <alpha-value>)",
          hover: "hsla(var(--danger-hover), <alpha-value>)",
          muted: "hsla(var(--danger-muted), <alpha-value>)",
          foreground: "hsla(var(--danger-foreground), <alpha-value>)",
        },
        border: "hsla(var(--border), <alpha-value>)",
      },
      borderRadius: {
        DEFAULT: "0.75rem",
        none: "0px",
        xs: "0.25rem",
        sm: "0.5rem",
        lg: "1rem",
        full: "9999px",
      }
    },
    extend: {
      keyframes: {
        "accordion-down": {
          from: { height: "0" },
          to: { height: "var(--accordion-content-height)" },
        },
        "accordion-up": {
          from: { height: "var(--accordion-content-height)" },
          to: { height: "0" },
        },
      },
      animation: {
        "accordion-down": "accordion-down 0.4s ease-out",
        "accordion-up": "accordion-up 0.4s ease-out",
      },
    },
    plugins: [],
  }
