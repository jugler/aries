#!/bin/bash

HTTPSTATUS="$(curl -s -o /dev/null -w "%{http_code}" http://canvas:8080/landscape/aries.htm)"
if [ $HTTPSTATUS -ne 200 ]
then
    echo "Aries service is not responding, restarting canvas server"
    sudo reboot 
else
    echo "Aries working correctly"
fi