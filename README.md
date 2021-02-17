# termux-go
A simple use of the `termux-api` to display a Message Bubble (Toast), and speak it out in voice format using gTTS
## Tools/Apps Used
 - termux
 - termux-api
## Setup and Run
```bash
apt update && apt install -y git openssh termux-api
sshd && nmap localhost
echo $(whoami)
```
Open Terminal on PC, and run these commands
```bash
ssh $(whoami)@ip -p 8022
curl -fSSL https://github.com/MahanthMohan/GopherChat/raw/master/battery
./battery | termux-toast && ./battery | termux-tts-speak -l english -p 1.0
```
## Results
A notification bubble should be display on screen, and the current battery status will be spoken out aloud

