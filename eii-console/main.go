package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

//Detecta o nome do sistema operacional
func DetectOSName() string {

	exec.Command("sh", "c", "cat /etc/os-release|tee os_release")

	out, err := exec.Command("sh", "-c", "cat os_info|grep -i ^NAME|cut -d= -f2").Output()

	if err != nil {
		fmt.Println("Erro ao tentar abrir arquivo de informações do sistema.")
		os.Exit(0)
	}

	return string(out)

}

//Detecta  o versão da distribuição (codename) sistema operacional
func DetectOSCodeName() string {

	exec.Command("sh", "c", "cat /etc/os-release|tee os_release")

	out, err := exec.Command("sh", "-c", "cat os_info|grep -i ^VERSION_CODENAME|cut -d= -f2").Output()

	if err != nil {
		fmt.Println("Erro ao tentar abrir arquivo de informações do sistema.")
		os.Exit(0)
	}

	return string(out)

}

//Detecta a versão da distribuição do ubuntu
func DetectOSVersion() string {

	exec.Command("sh", "c", "cat /etc/os-release|tee os_release")

	out, err := exec.Command("sh", "-c", "cat os_info|grep -i ^VERSION_ID|cut -d= -f2").Output()

	if err != nil {
		fmt.Println("Erro ao tentar abrir arquivo de informações do sistema.")
		os.Exit(0)
	}

	return string(out)

}

//REtorna informações sobre o sistema
func showOsInfo() string {

	out := fmt.Sprintf("\nInformações sobre o sistema:\nNome:%s\nCodename: %s\nVersão: %s\n",
		DetectOSName(), DetectOSCodeName(), DetectOSVersion())

	return out
}

//Instala pacotes do sistema operacional via apt
func installPackage(pkg string) {

	cmd, err := exec.Command("sh", "-c", fmt.Sprintf(" sudo apt-get install  -y %s ", pkg)).Output()

	if err != nil {

		panic(err.Error())

	}

	fmt.Println(string(cmd))
	fmt.Println("Pacote instalado com sucesso")

}

func execScript() {

	cmd := exec.Command("sh", "-c", "./install_ecidade.sh|tee instalador_facil.log")

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%v\n%v\n", stdoutBuf.String(), stderrBuf.String())

}

func main() {
	fmt.Printf("%s\n", "EII - E-Cidade Instalação Inteligente")

	fmt.Println("Verificando se há uma instalação do e-cidade")

	if _, err := os.Stat("/var/www/html/e-cidade/index.php"); err != nil {
		if os.IsNotExist(err) {
			execScript()
			return
		} else {

		}
	}
	fmt.Println("Já existe uma instalação existente do e-cidade.")
	fmt.Println("Se ainda assim quiser re-instalar, exclua manualmente a instalação exisstente no sevidor")
	fmt.Println("e execute o instalador novamente novamente.")
}
