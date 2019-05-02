#Send Aries code to the canvas
echo "Uploading everything to the canvas"
scp -r * pi@canvas:/home/pi/go/src/github.com/aries/
scp -r config/* pi@canvas:/home/pi/go/src/github.com/aries/

ssh pi@canvas 'sudo systemctl stop aries'
ssh pi@canvas 'GOPATH=/home/pi/go go build -o /home/pi/go/src/github.com/aries/aries /home/pi/go/src/github.com/aries/app/aries.go'
ssh pi@canvas 'sudo cp /home/pi/go/src/github.com/aries/aries.service /etc/systemd/system/aries.service'
ssh pi@canvas 'sudo systemctl daemon-reload'
ssh pi@canvas 'sudo systemctl start aries'