#Send Aries code to the canvas
echo "Uploading everything to the canvas"
scp -r * pi@tframe:/home/pi/go/src/github.com/aries/
scp -r config/* pi@tframe:/home/pi/go/src/github.com/aries/

ssh pi@tframe 'sudo systemctl stop aries'
ssh pi@tframe 'GOPATH=/home/pi/go go build -o /home/pi/go/src/github.com/aries/aries /home/pi/go/src/github.com/aries/app/aries.go'
ssh pi@tframe 'sudo cp /home/pi/go/src/github.com/aries/aries.service /etc/systemd/system/aries.service'
ssh pi@tframe 'sudo systemctl daemon-reload'
ssh pi@tframe 'sudo systemctl start aries'