# AnkoWeb: Anko Hypertext Processor

![AnkoWeb Build](https://github.com/nthnn/AnkoWeb/actions/workflows/build.yml/badge.svg)
[![License](https://img.shields.io/badge/license-GPL-blue.svg)](https://github.com/nthnn/AnkoWeb/blob/main/LICENSE)
<a href="https://www.buymeacoffee.com/nthnn"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" height="20px"></a>

AnkoWeb is built on top of the Anko virtual machine, which provides a robust execution environment for running AnkoWeb scripts. With AnkoWeb, you can create web applications that are both flexible and performant, thanks to its lightweight, yet powerful scripting capabilities. Originally, Anko is a scriptable interpreter written in Golang and was created by [@mattn](https://github.com/mattn).

## Basic Example

AnkoWeb allows you to embed server-side code directly within your HTML pages, making it a powerful tool for building dynamic web applications.

```aspnet
<!-- index.awp -->
<%
    func say(message) {
        echo("<h1 align=\"center\">" + message + "</h1>");
    }
%>
<!DOCTYPE html>
<head>
    <style>
        h1 {
            margin-top: 24%;
            font-family: Tahoma;
        }
    </style>
</head>
<body>
    <% say("Hello, from AnkoWeb!"); %>
</body>
</html>
```

## Running and Building

AnkoWeb provides several command-line arguments that allow you to customize the behavior of the AnkoWeb server when launching your web applications. These arguments help you control settings such as the host, working directory, and port.

### Options

Here are the available command-line options for AnkoWeb:

1. `-host <string>`
    - **Description**: Specifies the name of the localhost server.
    - **Default Value**: Null string.
    - **Usage**: You can set this option to customize the server's hostname when running your AnkoWeb application. For example, you can use `-host localhost` to set the hostname to "localhost."

2. `-path <string>`
    - **Description**: Defines the working directory for the server.
    - **Default Value**: The current directory where you run the `ankoweb` or `run.sh` command.
    - **Usage**: This option allows you to specify a different working directory for your AnkoWeb application. You can provide a path to a directory where your application's files are located. For example, you can use `-path /path/to/your/app` to set the working directory to "/path/to/your/app."

3. `-port <int>`
    - **Description**: Sets the port of the localhost server.
    - **Default Value**: 1234
    - **Usage**: You can use this option to define the port on which your AnkoWeb application will run. If you want to use a specific port, you can provide it as an argument. For example, you can use `-port 8080` to run your application on port 8080.

Note that the `run.sh` Bash script in the repository can also be used for running the source code on-the-go.

### Building

To build the AnkoWeb on your system, just type the following on your terminal:

```bash
./build.sh
```

This will generate executable AnkoWeb program if nothing went wrong inside the `bin` folder.

## Documentations

As of now, there is no available documentations yet. Instead, refer to the [examples](./examples) folder.

## Contributing

All contributions are welcome to make AnkoWeb even better. Whether you want to report a bug, suggest new features, or contribute code, your contributions are highly appreciated.

### Issue Reporting

If you encounter a bug, have a feature request, or want to suggest improvements, please open an issue on the [GitHub Issue Tracker](https://github.com/nthnn/AnkoWeb/issues). Be sure to provide as much detail as possible, including steps to reproduce the issue if applicable.

### Pull Requests

If you want to contribute code to AnkoWeb, follow these steps:

1. Fork the AnkoWeb repository to your GitHub account. And then clone it to your local machine.

    ```bash
    git clone https://github.com/nthnn/AnkoWeb
    ```

2. Create a new branch for your changes:

    ```bash
    git checkout -b feature/<your feature name>
    ```

3. You can now make changes to the repository.
4. Commit your changes:

    ```bash
    git add -A
    git commit -m "Add your meaningful commit message here"
    ```

5. Push your changes to your forked repository:

    ```bash
    git push origin feature/<your feature name>
    ```

6. Create a pull request (PR) from your branch to the main branch of the AnkoWeb repository.
7. Your PR will be reviewed, and any necessary changes will be discussed and implemented.
8. Once your PR is approved, it will be merged into the main branch, and your contribution will be part of AnkoWeb.

## License

AnkoWeb is open-source and licensed under the [GNU GPL v3 License](LICENSE).
