#Send Aries code to the canvas
echo "Uploading images to the canvas"
scp -r /Users/jugler/OneDrive/art/newportrait/* pi@canvas:/media/pi/CONFIG/imgs/portrait
mv /Users/jugler/OneDrive/art/newportrait/* /Users/jugler/OneDrive/art/portrait/

scp -r /Users/jugler/OneDrive/art/newlandscape/* pi@canvas:/media/pi/CONFIG/imgs/landscape
mv /Users/jugler/OneDrive/art/newlandscape/* /Users/jugler/OneDrive/art/landscape/


