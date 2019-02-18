#!/bin/bash
# No arguments -> next image on canvas
# First argument -> change type of image (art, photography, movieposters, etc 1 word)

TYPEOFART=$1
if [ -z "$TYPEOFART" ]
  then
    echo "Next Image"
    _="$(curl http://canvas:8080/portrait/updateConfig?nextImage=true 2>&1 > /dev/null)"
    _="$(curl http://canvas:8080/landscape/updateConfig?nextImage=true 2>&1 > /dev/null)"
    exit 0
fi

echo "Change Type of image to: $TYPEOFART"
_="$(curl http://canvas:8080/portrait/updateConfig?typeImage=$TYPEOFART 2>&1 >/dev/null)"
_="$(curl http://canvas:8080/landscape/updateConfig?typeImage=$TYPEOFART 2>&1 >/dev/null)"
exit 0
