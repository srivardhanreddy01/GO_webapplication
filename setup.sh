#!/bin/bash

echo "Script started"

# Update and install packages with sudo
sudo apt update
# sudo apt --assume-yes install mariadb-client mariadb-server -y

# Download the stable version of Go
wget https://golang.org/dl/go1.21.2.linux-amd64.tar.gz

# Extract the downloaded archive
sudo tar -C /usr/local -xzf go1.21.2.linux-amd64.tar.gz

# Set up environment variables
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile

# # Start MariaDB and enable on boot
# sudo systemctl start mariadb
# sudo systemctl enable mariadb

# # Automate MySQL Secure Installation
# SECURE_MYSQL=$(expect -c "
# set timeout 10
# spawn sudo mysql_secure_installation

# expect \"Enter current password for root (enter for none):\"
# send \"\r\"

# expect \"Set root password? \[Y/n\]\"
# send \"Y\r\"

# expect \"New password:\"
# send \"Sripragna\$1\r\"

# expect \"Re-enter new password:\"
# send \"Sripragna\$1\r\"

# expect \"Remove anonymous users? \[Y/n\]\"
# send \"Y\r\"

# expect \"Disallow root login remotely? \[Y/n\]\"
# send \"Y\r\"

# expect \"Remove test database and access to it? \[Y/n\]\"
# send \"Y\r\"

# expect \"Reload privilege tables now? \[Y/n\]\"
# send \"Y\r\"

# expect eof
# ")

# echo "$SECURE_MYSQL"

# # Execute SQL Commands
# sudo mysql -u root -pSripragna\$1 <<EOF
# CREATE DATABASE godatabase;
# GRANT ALL ON godatabase.* TO 'root'@'localhost' IDENTIFIED BY 'Sripragna\$1';
# FLUSH PRIVILEGES;
# SHOW DATABASES;
# EXIT
# EOF

# Verify that Go is installed
go version
