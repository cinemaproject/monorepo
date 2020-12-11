#! /bin/bash

sudo -i
# echo 'gce-user ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers; 
apt-get update
apt-get install -y apt-transport-https ca-certificates \
    curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt-get update
apt-get install -y docker-ce docker-ce-cli containerd.io
docker run -d -p 0.0.0.0:80:80 alibabaih/proxy

apt install -y python3-pip libgl1-mesa-glx libpq-dev 
# chmod -R 0777 /var/www/html
pip3 install --upgrade pip
pip3 install scikit-build
git clone https://github.com/cinemaproject/backend.git
cd backend
pip3 install -r requirements.txt
ln -s /opt/* /var/www/html
python3 app.py&
