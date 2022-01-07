<div id="top"></div>

<!-- PROJECT SHIELDS -->

[<div align="center"> ![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url] [![Stargazers][stars-shield]][stars-url]
[![MIT License][license-shield]][license-url]
[![Issues][issues-shield]][issues-url]
[![Issues Closed][issues-closed-shield]</div>][issues-closed-url]

<!-- ![Visitors](https://estruyf-github.azurewebsites.net/api/VisitorHit?user=wst24365888&repo=ez4o/convert-json-cli&countColor=rgb(0,%20126,%20198)) -->

<br />

![convert-json-cli](https://socialify.git.ci/ez4o/convert-json-cli/image?description=1&font=KoHo&name=1&owner=1&pattern=Circuit%20Board&theme=Light)

<!-- PROJECT LOGO -->
<br />
<div align="center">
<p align="center">
    <a href="https://github.com/ez4o/convert-json-cli#usage"><strong>Explore Usage »</strong></a>
    <br />
    <br />
    <a href="https://github.com/ez4o/convert-json-cli/issues">Report Bug</a>
    ·
    <a href="https://github.com/ez4o/convert-json-cli/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#usage">Usage</a></li>
        <li><a href="#parameters">Parameters</a></li>
      </ul>
    </li>
    <li><a href="#set-up-your-own">Set Up Your Own</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

![screenshot][product-screenshot]

**Convert-JSON-CLI** is an excellent tool for converting json files to structs
or classes in any programming language.

It currently only supports command line interface, but a web version will be
developed in the near future!

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

<!-- USAGE EXAMPLES -->

### Usage

`convjson [OPTIONS] [INPUT_FILE_PATH] [TARGET_LANGUAGE]`

### Parameters

| Parameter         | Necessity | Description                                                                                                                                              | Default Value                                           |
| ----------------- | --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- |
| `INPUT_FILE_PATH` | Required  | Path to your JSON file.                                                                                                                                  | None                                                    |
| `TARGET_LANGUAGE` | Required  | Target language. Currently supports: <br /> [go] [php] [python] [c] [swift] [csharp] [protobuf] [rust] [scala] [kotlin] [cpp] [dart] [java] [typescript] | None                                                    |
| `-o string`       | Optional  | Background image. Use url encode tool like <https://www.urlencoder.org/>.                                                                                | **.\\**[INPUT_FILE_DIR]**.**[TARGET_LANGUAGE_EXTENSION] |

### Example

#### Windows

`.\convjson.exe -o .\out\test.dart .\in\test.json dart`

#### Linux

`.\convjson -o .\out\test.dart .\in\test.json dart`

<p align="right">(<a href="#top">back to top</a>)</p>

## Set Up Your Own

These are some instructions on setting up your project locally, just follow
these simple steps.

### Prerequisites

- [GNU Make](https://community.chocolatey.org/packages/make)

- [Go](https://go.dev/doc/install)

### Installation

1. Clone the repo.

   ```sh
   git clone https://github.com/ez4o/convert-json-cli.git
   cd convert-json-cli
   ```

2. Install Go modules.

   ```sh
   go get -u
   ```

   > **Go module** acts really different between versions, if you have any
   > problem installing Go modules, please try:
   > 1. Upgrade Go version to `1.17.x`
   > 2. `go mod tidy -compat="1.17"`
   > 3. `go get -u`

### Generate Executable File

1. Use release command, and it will generate a executable file in root directory.

   ```sh
   make release-windows
   ```

   or

   ```sh
   make release-linux
   ```

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [ ] More Target Languages...

See the [open issues](https://github.com/ez4o/convert-json-cli/issues) for a
full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to
learn, inspire, and create. Any contributions you make are **greatly
appreciated**.

If you have a suggestion that would make this better, please fork the repo and
create a pull request. You can also simply open an issue with the tag
"enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feat/amazing-feature`)
3. Commit your Changes with
   [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
4. Push to the Branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See
[LICENSE](https://github.com/ez4o/convert-json-cli/blob/main/LICENSE) for more
information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

### Author

- HSING-HAN, WU (Xyphuz)
  - Mail me: xyphuzwu@gmail.com
  - About me: <https://about.xyphuz.com>
  - GitHub: <https://github.com/wst24365888>

### Project Link

- <https://github.com/ez4o/convert-json-cli>

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/ez4o/convert-json-cli.svg?style=for-the-badge
[contributors-url]: https://github.com/ez4o/convert-json-cli/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/ez4o/convert-json-cli.svg?style=for-the-badge
[forks-url]: https://github.com/ez4o/convert-json-cli/network/members
[stars-shield]: https://img.shields.io/github/stars/ez4o/convert-json-cli.svg?style=for-the-badge
[stars-url]: https://github.com/ez4o/convert-json-cli/stargazers
[issues-shield]: https://img.shields.io/github/issues/ez4o/convert-json-cli.svg?style=for-the-badge
[issues-url]: https://github.com/ez4o/convert-json-cli/issues
[issues-closed-shield]: https://img.shields.io/github/issues-closed/ez4o/convert-json-cli.svg?style=for-the-badge
[issues-closed-url]: https://github.com/ez4o/convert-json-cli/issues?q=is%3Aissue+is%3Aclosed
[license-shield]: https://img.shields.io/github/license/ez4o/convert-json-cli.svg?style=for-the-badge
[license-url]: https://github.com/ez4o/convert-json-cli/blob/main/LICENSE
[product-screenshot]: https://convert-json-cli.ez4o.com/?username=wst24365888&img_url=https%3A%2F%2Fimages.unsplash.com%2Fphoto-1506744038136-46273834b3fb%3Fixid%3DMnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8%26ixlib%3Drb-1.2.1%26auto%3Dformat%26fit%3Dcrop%26w%3D1000%26q%3D80&fbclid=IwAR1AUDKHzjzBSjKle6J44dYRSrIbvBu8eTxtrfhpPxhBnBsOizgSq63bYbU
