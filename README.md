<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
<h3 align="center">GoLang RabbitMQ integration</h3>

  <p align="center">
    Simple sender and receiver integration in go
    <br />
    <a href="https://github.com/proapi/go-rabbitmq"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/proapi/go-rabbitmq">View Demo</a>
    ·
    <a href="https://github.com/proapi/go-rabbitmq/issues">Report Bug</a>
    ·
    <a href="https://github.com/proapi/go-rabbitmq/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Simple go application to check the capabilities of RabbitMQ queues.

* Dockerfiles for both consumer and sender
* docker-compose for easier project development
* official amqp go library


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][Go]][Go-url]
* [![RabbitMQ][RabbitMQ]][RabbitMQ-url]
* [![Docker][Docker]][Docker-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

All you need to have is a working docker environment with docker-compose
* docker
  ```sh
  docker compose build
  docker compose up
  ```

### Installation

No further installation needed.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

With working docker and all containers up we can do a simple http call to sender container

```shell
http GET http://localhost:3000/send\?msg\=test
```

You should be able to get the message on the consumer container in the logs.

_For more examples, please refer to the [Documentation](https://github.com/proapi/go-rabbitmq)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/awesome`)
3. Commit your Changes (`git commit -m 'feat: add awesome feature'`)
4. Push to the Branch (`git push origin feature/awesome`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

PROAPI Michał Pawelski - [@_proapi_](https://twitter.com/_proapi_) - kontakt@proapi.eu

LinkedIn: [https://www.linkedin.com/in/michal-pawelski](https://www.linkedin.com/in/michal-pawelski/)

Project Link: [https://github.com/proapi/go-rabbitmq](https://github.com/proapi/go-rabbitmq)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/proapi/go-rabbitmq.svg?style=for-the-badge
[contributors-url]: https://github.com/proapi/go-rabbitmq/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/proapi/go-rabbitmq.svg?style=for-the-badge
[forks-url]: https://github.com/proapi/go-rabbitmq/network/members
[stars-shield]: https://img.shields.io/github/stars/proapi/go-rabbitmq.svg?style=for-the-badge
[stars-url]: https://github.com/proapi/go-rabbitmq/stargazers
[issues-shield]: https://img.shields.io/github/issues/proapi/go-rabbitmq.svg?style=for-the-badge
[issues-url]: https://github.com/proapi/go-rabbitmq/issues
[license-shield]: https://img.shields.io/github/license/proapi/go-rabbitmq.svg?style=for-the-badge
[license-url]: https://github.com/proapi/go-rabbitmq/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/michal-pawelski/
[Go]: https://img.shields.io/badge/go-000000?style=for-the-badge&logo=go
[Go-url]: https://go.dev/
[RabbitMQ]: https://img.shields.io/badge/RabbitMQ-000000?style=for-the-badge&logo=rabbitmq
[RabbitMQ-url]: https://www.rabbitmq.com
[Docker]: https://img.shields.io/badge/Docker-000000?style=for-the-badge&logo=docker
[Docker-url]: https://www.docker.com/
