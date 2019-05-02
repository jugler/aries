#Send Aries code to the canvas
echo "Uploading everything to the canvas"
scp * pi@canvas:/home/pi/aries
scp -r config/* pi@canvas:/home/pi/aries/config/

ssh pi@canvas 'sudo systemctl stop aries'
ssh pi@canvas 'go build -o aries/aries aries/aries.go '
ssh pi@canvas 'sudo cp aries/aries.service /etc/systemd/system/aries.service'
ssh pi@canvas
ssh pi@canvas 'sudo systemctl start aries'