#Send Aries code to the canvas
echo "Uploading everything to the canvas"
scp * pi@canvas:/home/pi/aries
scp -r config/* pi@canvas:/home/pi/aries/config/

