# SmartKickers

## Overview

SmartKickers is a project to make your kickers more immersive. It uses Ximea camera to track the ball. 

The following repository contains: 
* Golang server for processing the data.
* React client to display game information.

You will also need SmartKickersAI provided by the developers per request. It contains:
* Ball tracking software.
* Node server responsible for sending data to the Golang server.

## Features

* Keeping track of the score.

## Requirements

* python
* node
* npm
* ximea camera drivers
* macOS

#### Supported versions:

* python 3.10.5
* node 16.16.0
* npm 8.11.0
* macOS 12.4

## Installation

We are using macOS to build and run everything locally for now.

> **NOTE:** Ensure you have also SmartKickersAI source code.

1. Clone the repository.

    ```bash 
    git clone git@github.com:HackYourCareer/SmartKickers.git
    ```

2. Run the requirementsCheck.sh script to check if your system fulfils all requirements.

    ```bash 
    ./requirementsCheck.sh
    ```

    > **NOTE:** Be sure to give permission for the script to execute.

3. Install missing dependencies if any missing.

4. Switch to **SmartKickers/frontend/smart-kickers-game** and run the following command.

    ```bash 
    npm install
    ```

## Launching

> **NOTE:** Before the first launch SmartKickersAI needs to be configured fot the table where it will be used. More details in SmartKickersAI README

1. Switch to **SmartKickers/backend** and launch the go server.

    ```bash 
    go run main.go
    ```

2. Switch to **SmartKickersAI/LocalServer/server** and launch the node server.

    ```bash 
    node server.js
    ```

3. Switch to **SmartKickers/frontend/smart-kickers-game** and launch the react client.

    ```bash 
    npm start
    ```