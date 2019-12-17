#Send Aries code to the canvas
echo "Uploading everything to the canvas"
scp -r * pi@canvas:/home/pi/go/src/github.com/jugler/aries/
scp -r config/* pi@canvas:/home/pi/go/src/github.com/jugler/aries/

ssh pi@canvas 'sudo systemctl stop aries'
ssh pi@canvas 'GOPATH=/home/pi/go go build -o /home/pi/go/src/github.com/jugler/aries/aries /home/pi/go/src/github.com/jugler/aries/app/aries.go'
ssh pi@canvas 'sudo cp /home/pi/go/src/github.com/jugler/aries/aries.service /etc/systemd/system/aries.service'
ssh pi@canvas 'sudo cp /home/pi/go/src/github.com/jugler/aries/canvas_watchdog.service /etc/systemd/system/canvas_watchdog.service'
ssh pi@canvas 'sudo systemctl daemon-reload'
ssh pi@canvas 'sudo systemctl start aries'