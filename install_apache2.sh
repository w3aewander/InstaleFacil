echo "INSTALANDO APACHE"
sleep 2
apt-get --force-yes --yes install apache2


clear
echo "CONFIGURANDO VIRTUAL HOST"
sleep 2
a2enmod rewrite
a2dissite 000-default default-ssl

cat << APADEF > /etc/apache2/sites-available/ecidade.conf
<VirtualHost *:80>
	
        LimitRequestLine 16382
        LimitRequestFieldSize 16382
        Timeout 12000
        AddDefaultCharset ISO-8859-1
        SetEnv no-gzip 1
        <Directory /var/www/html>
            Options -Indexes +FollowSymLinks +MultiViews
            AllowOverride All
            Require all granted
        </Directory>

	ServerAdmin webmaster@localhost
	DocumentRoot /var/www/html

	ErrorLog /var/log/apache2/error.log
	CustomLog /var/log/apache2/access.log combined

</VirtualHost>
       
APADEF

a2ensite ecidade
/etc/init.d/apache2 restart