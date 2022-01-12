# GO-RMA
> The Go Version of the Resource Management Layer.

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/AliceTrinta/GO-RMA)

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Contact](#contact)
* [Contributing](#contributing)



<!-- ABOUT THE PROJECT -->
## About The Project
The Resource Management Architecture (RMA) is a way to organize Objects on an IoT network. The RMA is composed 
by the Application Layer, the Resource Management Layer(RML) and the Device's Layer.

### Built With
This work uses:
* [MongoDB](https://www.mongodb.com/golang)
* [NATS](https://nats.io/)


<!-- GETTING STARTED -->
## Getting Started

Follow pass-by-pass to install the project in your machine.

### Prerequisites

Have Installed MongoDB on your machine.
Have Installed NATS server on your machine.

### Installation

Clone repository
```sh
git clone https://github.com/AliceTrinta/GO-RMA.git
```

### Running

Firste step: create a database called "vcdb" on your MongoDB, with four collections called "device", "data", "environment" and "action".

Second step: run the main.go file
```sh
GO-RMA
     |
     ---main.go
```

Third step: Run the device application.

Fourth step: Run RMA-GO.

Now, you can call the endpoints listed above:
```sh
GET - http://localhost:8080/listDevice
```
```sh
GET - http://localhost:8080/listData
```



<!-- USAGE -->
## Usage
The RMA is used to organize object on an IoT network, offering a management tool to turn easier deal with object's resources and user's applications.

<!-- CONTACT -->
## Contact

Alice Trinta – [@malicetrinta](https://www.instagram.com/malicetrinta/) – maria.trinta@aluno.cefet-rj.br

[https://www.researchgate.net/profile/Maria_Alice_Lima2/publications](https://www.researchgate.net/profile/Maria_Alice_Lima2/publications)
Project Link: [https://github.com/AliceTrinta/GO-RMA](https://github.com/AliceTrinta/GO-RMA)



<!-- CONTRIBUTING -->
## Contributing

1. Fork it (<https://github.com/AliceTrinta/GO-RMA>)
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- Markdown link & img dfn's -->
[npm-image]: https://img.shields.io/npm/v/datadog-metrics.svg?style=flat-square
[npm-url]: https://npmjs.org/package/datadog-metrics
[npm-downloads]: https://img.shields.io/npm/dm/datadog-metrics.svg?style=flat-square
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
