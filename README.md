# SmartKickers

## Overview

SmartKickers is a project to make your kickers more immersive. It uses Ximea camera to track the ball.

The following repository contains:

- Golang server for processing the data
- React client to display game information

## Table of contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Launching](#launching)
  - [Launching the server and web application](#launching-the-server-and-web-application)
  - [Launching the ball-tracking software](#launching-the-ball-tracking-software)
- [Contact Information](#contact-information)

## Features

- Keeping track of the score
- Manually adjusting the score

## Prerequisites

Get the following:

1. [Git](https://git-scm.com/downloads)
2. [Python](https://www.python.org/downloads/) in version 3.10.5 or higher
3. [Node.js](https://nodejs.org/en/download/) in version 16.16.0 or higher
4. Access to [SmartKickersAI](https://github.com/HackYourCareer/SmartKickersAI)

   > **NOTE:** To get access to the SmartKickersAI repository, [contact](#contact-information) Team Beavers.

5. [Ximea camera drivers](https://www.ximea.com/support/wiki/apis/)
6. Docker daemon for example Docker Desktop:
   - [MacOS](https://docs.docker.com/desktop/install/mac-install/)
   - [Windows](https://docs.docker.com/desktop/install/windows-install/)
   - [Linux](https://docs.docker.com/desktop/install/linux-install/)
7. [NumPy](https://numpy.org/install/)

   > **TIP:** If you're a macOS user, you can also use `brew` to install NumPy:

   ```bash
   brew install numpy
   ```

8. [OpenCV-Python](https://pypi.org/project/opencv-python/)

9. [Imutils](https://pypi.org/project/imutils/)

> **TIP:** If you're a macOS user, you can run the [`requirementsCheck.sh`](requirementsCheck.sh) script to check for any missing dependencies. Note that the script looks for OpenCV, NumPy, Imutils, and the Ximea driver **only after you have successfully installed Python!** If you don't have Python, it doesn't print missing Python-dependent packages except Python itself.
>
> ```bash
> sudo ./requirementsCheck.sh
> ```

## Installation

1. Create a new directory and navigate to it in a terminal.

   ```bash
   mkdir smartkickers
   cd smartkickers
   ```

2. Clone the repository.

   ```bash
   git clone https://github.com/HackYourCareer/SmartKickersAI.git
   ```

3. Pull the frontend image from GCR.

   ```bash
   docker pull ghcr.io/hackyourcareer/smartkickers-frontend:latest
   ```

4. Pull the backend image from GCR.

   ```bash
   docker pull ghcr.io/hackyourcareer/smartkickers-backend:latest
   ```

## Launching

> **NOTE:** The configuration files for the camera view are provided by the developers in the [`SmartKickersAI`](https://github.com/HackYourCareer/SmartKickersAI/tree/main/LocalServer) repository.

### Launching the server and web application

Run the pulled containers to launch the server and web application.

1. In a terminal run:

   ```bash
   docker run -p 3000:3000 ghcr.io/hackyourcareer/smartkickers-backend:latest
   ```

2. Open a new terminal and run:

   ```bash
   docker run -p 3007:80 ghcr.io/hackyourcareer/smartkickers-frontend:latest
   ```

### Launching the ball-tracking software

1. Plug in the camera to the computer.

2. To launch the application, navigate to the `SmartKickersAI` repo and double click the `start` file.

   On successful connection to the Go server, you see the following output:

   ```bash
    {
      url: 'ws://127.0.0.1:3000',
      wsurl: 'ws://127.0.0.1:3000/shot',
      dispatcherlURL: 'ws://127.0.0.1:3000',
      cfURL: 'ws://127.0.0.1:3000/cf',
      tableID: '10',
      token: 'c3d8c29ec43b9eadb8bc80ad1458ab8',
      messageTypes: {
          BallPosition: '615eec77d6be09356891',
          Shot: '2ca34aee4bb3dec060b1',
          BestShotGuid: 'c01b38461a678f3eefa5'
      }
   }
   init
   conn!
   { type: 'utf8', utf8Data: '{"start":"10"}' }
   { start: '10' }
   game id
   10
   connected to IoT Services via wss
   ```

3. In your Internet browser, go to [`localhost:3007`](http://localhost:3007/) to open the React application.

   You can now play the Smart Kickers game. Enjoy!

   ![React application](assets/reactApp.png "React application")

## Contact Information

Team Beavers members:

- [Piotr Kołodziejski](https://github.com/Pichi00) - piotr.kolodziejski@sap.com
- [Marek Kawalski](https://github.com/marekkawalski) - marek.kawalski@sap.com
- [Michał Kalke](https://github.com/MichalKalke) - michal.kalke@sap.com
- [Kacper Małachowski](https://github.com/KacperMalachowski) - kacper.malachowski@sap.com
- [Marek Michali](https://github.com/MarekMichali) - marek.michali@sap.com
- [Filip Gołyszny](https://github.com/Filip22022) - filip.golyszny@sap.com
