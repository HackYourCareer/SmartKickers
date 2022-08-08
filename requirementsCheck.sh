#!/bin/bash

checkPassed=true

echo checking node...
if ! node --version > /dev/null 2>&1; then
    echo Node.js not detected!
    checkPassed=false
fi

echo checking npm...
if ! npm --version > /dev/null 2>&1; then
    echo NPM not detected!
    checkPassed=false
fi

echo checking python3...
if ! python3 --version > /dev/null 2>&1; then
    echo Python not detected!
    exit 1
fi

echo checking Ximea driver...
if ! python3 -c "import ximea" > /dev/null 2>&1; then
    echo Ximea driver not detected!
    checkPassed=false
fi

echo checking opencv-python...
if ! python3 -c "import cv2" > /dev/null 2>&1; then
    echo opencv-python not detected!
    checkPassed=false
fi

echo checking numpy...
if ! python3 -c "import numpy" > /dev/null 2>&1; then
    echo numpy not detected!
    checkPassed=false
fi

echo checking imutils...
if ! python3 -c "import imutils" > /dev/null 2>&1; then
    echo imutils not detected!
    checkPassed=false
fi

if [ "$checkPassed" = true ]; then
    echo Everything looks fine!
    exit 0
fi

exit 1