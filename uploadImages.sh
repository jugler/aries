#Send Aries code to the canvas
echo "Uploading images to the canvas"
scp -r /Users/jugler/OneDrive/art/portrait/* pi@canvas:/home/pi/aries_deprecated/imgs/portrait

scp -r /Users/jugler/OneDrive/art/landscape/* pi@canvas:/home/pi/aries_deprecated/imgs/landscape

