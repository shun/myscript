#!/bin/bash

if [ $1 = "on" ]; then
  xset -display :0.0 dpms force on
  echo "on"

elif [ $1 = "off" ]; then
  xset -display :0.0 dpms force off
  echo "off"

fi
