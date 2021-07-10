# Create an instance

1. Create an AWS Account
   - enter credit card info

2. Sign into console

3. Create a new EC2 instance

    - services / EC2
    - launch instance
    - choose your instance
    - add storage / 30GB free
    - add tags / webserver
    - security / ssh / http
    - launch
    - create new key pair / download

---

# Deploy your binay

- mv \[src\]\[dst\] / sudo chmod 400 your.pem
- Build hello world
  - GOOS=linux GOARCH=amd64 go build -o self-define-name
- Copy your binary into the server
  - scp -i /path/to/\[your\].pem ./main ec2-user@[public-DNS]:
  - say yes to The authenticity of host ... can't be established.

- SSH into your server
  - ssh -i /path/to/\[your\].pen ec2-user@[public-DNS]:
  
- Run your code
  - sudo chmod 700 mybinary
  - sudo ./mybinary
  - check it in a browser at [public-ip]

- Exit
  - ctrl + c
  - exit

---

# Persisting your application

To run application after the terminal session has ended, we must do one of the followings:

## Possible options

1. screen
2. init.d
3. upstart
4. system.d

## System.d

1. create a configuration file
   - cd /etc/systemd/system/
   - sudo nano ```<filename>```.service

```Bash
[Unit]
Description=GO Server

[Service]
ExecStart=/home/<username>/<exepath>
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

2. Add the service to systemd
    - sudo systemctl enable ```<filename>```.sevice
  
3. Activate the service
   - sudo systemctl start ```<filename>```.service

4. Check if systemd started it.
   - sudo systemctl status ```<filename>```

5. Stop systemd if so desired
   - sudo systemctl stop ```<filename>```

---

# Troubleshooting

A possible issue could be that you're cross-compiling for the wrong architecture: AWS might have assigned you a different machine than the one used in this example. To solve this problem, we will install Go on the AWS machine and then run "go env" to see GOOS & GOARCH for that machine.

- Download Go
  - wget https://storage.googlepis.com/golang/go1.7.4.linux-amd64.tar.gz
- unpack go
  - tar -xzf go1.7.4.linux-amd64.tar.gz
- remove the tar file
  - rm -rf go1.7.4.linuz-amd64.tar.gz
- make your go workspace
  - mkdir goworkspace
  - cd goworkspace
  - mkdir bin pkg src
  - cd ../
- add environment variables
  - nano .bashrc
    ```Bash
    export GOROOT=/home/ubuntu/go
    export GOPATH=/home/ubuntu/goworkspace
    export PATH=$PATH:/home/ubuntu/goworkspace/bin
    export PATH=$PATH:/home/ubuntu/go/bin
    ```
- refresh environment variables
  - source ~/.bashrc
- confirm installation
  - go version
- get machine GOOS & GOARCH info
  - go env
  
---

# Troubleshooting

Sometimes students miss setting port openings in security. If you are having issues, check to make sure these settings are correct - and please note, your IP address for SSH will either be 0.0.0.0/0 or something different than mine.
