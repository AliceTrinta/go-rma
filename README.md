# GO-RMA
> The Go Version of the Resource Maanagement Layer.

[![Generic badge](https://img.shields.io/badge/Version-0.0.4-<COLOR>.svg)](https://shields.io/)
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
* [Release History](#release-history)
* [Contact](#contact)
* [Contributing](#contributing)



<!-- ABOUT THE PROJECT -->
## About The Project
The Resource Management Architecture (RMA) is a way to organize Objects on an IoT network. The RMA is composed 
by the Application Layer, the Resource Management Layer(RML) and the Device's Layer.

### Built With
This work uses:
* [MongoDB](https://www.mongodb.com/golang)


<!-- GETTING STARTED -->
## Getting Started

Follow pass-by-pass to install the project in your machine.

### Prerequisites

Have Installed MongoDB database in your machine.

### Installation

Clone repository
```sh
git clone https://github.com/AliceTrinta/GO-RMA.git
```

### Running

Firste step: create a collection called "example".

Second step: run the main.go file
```sh
GO-RMA
     |
     ---main.go
```

Third step: Enter with a JSON string of a type device, data or action.
```sh
Example of a device: {"_id":"device","UUID":"8e0dae08-b360-4ace-b214-47b3d0e93f1a","gatewayUUID":"","name":"Farmer","description":"Farmer's device'","cycleDelayInMillis":"5000","resourceList":[{"_id":"irrigator","name":"irrigator 1","description":"This is the Irrigator 1","port":"COM3","dataUnit":"","waitTimeInMillis":5000,"commandList":[{"id":"ON","description":"Irrigator is working"},{"id":"OFF","description":"Irrigator is not working"}]},{"_id":"light","name":"light","description":"This is the light sensor","port":"COM3","dataUnit":"","waitTimeInMillis":5000,"commandList":[]},{"_id":"humidity","name":"humidity","description":"This is the humidity sensor","port":"COM3","dataUnit":"","cycleDelayInMillis":5000,"commandList":[]},{"_id":"temperature","name":"temperature","description":"This is the temperature sensor","port":"COM3","dataUnit":"°C","waitTimeInMillis":5000,"commandList":[]},{"_id":"pH","name":"pH","description":"This is the humidity sensor","port":"COM3","dataUnit":"pH","waitTimeInMillis":5000,"commandList":[]}],"lastUpdate":"2020-10-04T23:00:40.843+00:00"}"
```



<!-- USAGE -->
## Usage
The RMA is used to organize object on an IoT network, offering a management tool to turn easier deal with object's resources and user's applications.

<!-- RELEASE HISTORY -->
## Release History

* 0.0.1
    * First version
* 0.0.2
    *  Version after first meeting
* 0.0.4
    * Begining first tests



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
