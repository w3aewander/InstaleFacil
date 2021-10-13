echo "INSTALANDO PHP 5.6"
sleep 2

sudo apt clean
sudo apt  autoremove
sudo apt -y install software-properties-common
add-apt-repository -y ppa:ondrej/php

sudo mkdir /var/log/php 2> /dev/null 

sudo apt --allow --yes install php5.6 php5.6-bcmath php5.6-bz2 php5.6-cli php5.6-common php5.6-curl php5.6-gd php5.6-interbase php5.6-json php5.6-mbstring php5.6-mcrypt php5.6-pgsql php5.6-soap php5.6-sqlite3 php5.6-xml php5.6-xmlrpc php5.6-zip unzip zip php5.6-common php5.6-interbase libapache2-mod-php5.6
#sudo apt --force-yes --yes install php5.6 php5.6-bcmath php5.6-bz2 php5.6-cli php5.6-common php5.6-curl php5.6-gd php5.6-interbase php5.6-json php5.6-mbstring php5.6-mcrypt php5.6-pgsql php5.6-soap php5.6-sqlite3 php5.6-xml php5.6-xmlrpc php5.6-zip unzip zip php5.6-common php5.6-interbase libapache2-mod-php5.6

sudo apt update
sudo apt install -y php5.6

clear
echo "CONFIGURANDO PHP 5.6"
sleep 2
mkdir /var/www/html/tmp 2> /dev/null 
chown -R www-data. /var/www/html/tmp 2>/dev/null
chmod -R 777 /var/www/html/tmp 2> /dev/null
mkdir /var/www/tmp 2>/dev/null
chown -R www-data. /var/www/tmp
chmod -R 777 /var/www/tmp


# Configura o php5
PHPPATH=/etc/php/5.6/apache2/php.ini
sed_configuracao "register_argc_argv = On" "$PHPPATH"
sed_configuracao "post_max_size = 200M" "$PHPPATH"
sed_configuracao "upload_max_filesize = 200M" "$PHPPATH"
sed_configuracao "default_socket_timeout = 60000" "$PHPPATH"
sed_configuracao "max_execution_time = 60000" "$PHPPATH"
sed_configuracao "max_input_time = 60000" "$PHPPATH"
sed_configuracao "memory_limit = 512M" "$PHPPATH"
sed_configuracao "display_errors = Off" "$PHPPATH"
sed_configuracao "log_errors = On" "$PHPPATH"
sed_configuracao "session.gc_maxlifetime = 7200" "$PHPPATH"
sed_configuracao "short_open_tag = On" "$PHPPATH"
sed_configuracao "error_log = /var/log/apache2/php_errors.log" "$PHPPATH"
sed_configuracao 'date.timezone = "America/Sao_Paulo"' "$PHPPATH"
sed_configuracao "ignore_repeated_errors = On" "$PHPPATH"

PHPPATH=/etc/php/5.6/cli/php.ini
sed_configuracao "register_argc_argv = On" "$PHPPATH"
sed_configuracao "post_max_size = 200M" "$PHPPATH"
sed_configuracao "upload_max_filesize = 200M" "$PHPPATH"
sed_configuracao "default_socket_timeout = 60000" "$PHPPATH"
sed_configuracao "max_execution_time = 60000" "$PHPPATH"
sed_configuracao "max_input_time = 60000" "$PHPPATH"
sed_configuracao "memory_limit = 512M" "$PHPPATH"
sed_configuracao "display_errors = Off" "$PHPPATH"
sed_configuracao "log_errors = On" "$PHPPATH"
sed_configuracao "session.gc_maxlifetime = 7200" "$PHPPATH"
sed_configuracao "short_open_tag = On" "$PHPPATH"
sed_configuracao "error_log = /var/log/php/php_errors.log" "$PHPPATH"
sed_configuracao 'date.timezone = "America/Sao_Paulo"' "$PHPPATH"
sed_configuracao "ignore_repeated_errors = On" "$PHPPATH"

PHPPATH=/etc/php/5.6/apache2/php.ini

sed 's,~E_DEPRECATED,~E_NOTICE,g'  /etc/php/5.6/apache2/php.ini > /tmp/php.ini
sed 's,~E_STRICT,~E_DEPRECATED,g'  /tmp/php.ini > /tmp/php2.ini

rm /etc/php/5.6/apache2/php.ini
mv /tmp/php2.ini /etc/php/5.6/apache2/php.ini
chmod 777 /etc/php/5.6/apache2/php.ini


mkdir /var/log/php
touch /var/log/php/php_errors.log
chmod 777 /var/log/php/php_errors.log
touch /var/log/apache2/php_errors.log
chmod 777 /var/log/apache2/php_errors.log
cp /etc/logrotate.d/postgresql-common /etc/logrotate.d/php
sed -i 's/postgresql/php/g' /etc/logrotate.d/php

/etc/init.d/apache2 restart