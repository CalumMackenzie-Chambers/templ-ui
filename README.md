# 🎨 templ-ui

**templ-ui** is a UI component library for [templ](https://templ.guide/) - a powerful templating language for Go. Heavily influenced by [shadcn](https://ui.shadcn.com/) and [radix-ui](https://www.radix-ui.com/), this repository contains a suite of UI components styled using [Tailwind CSS](https://tailwindcss.com/) that are ready to drop into your templ projects.

## 🚀 Getting Started

Before you begin, make sure you're familiar with [Go](https://go.dev/) and [templ](https://github.com/a-h/templ). If you're new to templ, check out their [documentation](https://templ.guide/) to get started.

### Usage

To use **templ-ui** in your project:

1. Choose the component you want from the `/components` directory.
2. Copy the `.templ` file for that component into your application.
3. Integrate it with your templ setup.

It's as simple as that!

> [!NOTE]
> There is a plan to create a CLI tool that will allow you to install components directly into your project. This is currently a work in progress.

### Building from Source

If you'd like to contribute to **templ-ui** or use it as a starting point for your own UI component library, you can build it from source:

1. Clone the repository: `git clone https://github.com/CalumMackenzie-Chambers/templ-ui.git`
2. Install Go dependencies specified in `go.mod`: `go mod download`
3. Install the templ CLI: `go install github.com/a-h/templ/cmd/templ@latest`
4. Install [Air](https://github.com/cosmtrek/air) for live reloading: `go install github.com/cosmtrek/air@latest`
5. Run `npm install` to install the necessary node modules.
6. Start the development server with `npm run dev`. When you add or modify a component, ensure to create or update its example on the website.

## 📜 License

**templ-ui** is open-sourced software licensed under the [MIT license](LICENSE).

## ❤️ Acknowledgements

- [templ](https://templ.guide/): For providing the powerful templating language.
- [shadcn](https://ui.shadcn.com/): For inspiring the UI component structure.
- [Tailwind CSS](https://tailwindcss.com/): For the utility-first CSS framework that makes styling a breeze.
- [Radix UI](https://www.radix-ui.com/): For the excellent UI primitives that inspire accessibility and customization.

## 🙋 Contributing

Interested in contributing to **templ-ui**? We love your enthusiasm! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details on how you can help and our code of conduct.

Happy coding! 🎉
