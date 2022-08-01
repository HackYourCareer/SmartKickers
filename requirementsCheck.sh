#!/bin/bash

checkPassed=true

if ! python3 --version > /dev/null 2>&1; then
    echo Python 3.x not detected!
    checkPassed=false
fi

if ! node --version > /dev/null 2>&1; then
    echo Node.js not detected!
    checkPassed=false
fi

if ! npm --version > /dev/null 2>&1; then
    echo NPM not detected!
    checkPassed=false
fi

if ! python3 -c "import ximea" > /dev/null 2>&1; then
    echo Ximea driver not detected!
    checkPassed=false
fi

if ! python3 -c "import cv2" > /dev/null 2>&1; then
    echo opencv-python not detected!
    checkPassed=false
fi

if ! python3 -c "import numpy" > /dev/null 2>&1; then
    echo numpy not detected!
    checkPassed=false
fi

if ! python3 -c "import imutils" > /dev/null 2>&1; then
    echo imutils not detected!
    checkPassed=false
fi

if [ "$checkPassed" = true ]; then
    echo Everything looks fine!
    exit 0
fi

exit 1