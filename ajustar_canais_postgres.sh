echo "CONFIGURANDO CANAIS DE SOFTWARE DO POSTGRESQL"
sleep 2

#adicionando repositorio do postgresql
touch /etc/apt/sources.list.d/pgdg.list
echo "deb http://apt.postgresql.org/pub/repos/apt/ ${RELEASE}-pgdg main" | tee /etc/apt/sources.list.d/pgdg.list
apt-key adv --keyserver keyserver.ubuntu.com --recv 7FCC7D46ACCC4CF8
