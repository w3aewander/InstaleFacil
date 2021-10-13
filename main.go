package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"text/template"

	"golang.org/x/crypto/ssh"

	"github.com/gorilla/mux"

	"gopkg.in/ini.v1"
)

type InforPage struct {
	PageTitle  string
	Title      string
	Message    string
	Output     string
	Err        string
	AppVersion string
}

type FormContato struct {
	nome    string
	email   string
	assunto string
	message string
}

var data InforPage
var ssh_server string
var ssh_port string
var app_version string

func init() {

	ssh_server = ""
	ssh_port = ""
	data.PageTitle = "Instale Fácil"
	data.Message = "Intalação de Pacotes"

	data.Output = "As informações serão exibidas aqui"

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	ssh_server = cfg.Section("server").Key("ssh_server_ip").String()
	ssh_port = cfg.Section("server").Key("ssh_server_port").String()

	app_version = cfg.Section("").Key("app_version").String()
	data.AppVersion = app_version

}

func readPubKey(file string, keyPass string) ssh.AuthMethod {
	var key ssh.Signer
	var err error
	var b []byte
	b, err = ioutil.ReadFile(file)
	mustExec(err, "failed to read public key")
	if !strings.Contains(string(b), "ENCRYPTED") {
		key, err = ssh.ParsePrivateKey(b)
		mustExec(err, "failed to parse private key")
	} else {
		key, err = ssh.ParsePrivateKeyWithPassphrase(b, []byte(keyPass))
		mustExec(err, "failed to parse password-protected private key")
	}

	return ssh.PublicKeys(key)
}

func mustExec(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:\n  %s", msg, err)
	}

}

// Conexão SSH nativo
// Protótipo do comando go run main.go -command whoami -host 192.168.0.7 -key-path ~/.ssh/id_rsa -port 22 -user ecidade
func connSsh(user string, host string, port string, command string, keyPath string, keyPass string) string {

	conf := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			readPubKey(keyPath, keyPass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // XXX: Security issue
	}
	client, err := ssh.Dial("tcp", strings.Join([]string{host, ":", port}, ""), conf)
	mustExec(err, "failed to dial SSH server")
	session, err := client.NewSession()
	mustExec(err, "failed to create SSH session")
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(command)
	mustExec(err, "failed to run command over SSH")

	//resp := fmt.Sprintf("%s: %s", command, b.String())

	resp := fmt.Sprintf("%s\n", b.String())

	return resp

}

//Rota para a pagina HOME
func homeHandle(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/layout.html", "static/index.html"))

	data.PageTitle = "Home"
	data.Title = "Página Inicial"
	data.Message = ""
	data.AppVersion = app_version

	tmpl.Execute(w, data)
}

//Iniciar a instalação do e-cidade
func handleInstall(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/layout.html", "static/instalar.html"))

	data.PageTitle = "Instalação"
	data.Title = "Procedimentos de instalação"
	data.Output = "Em execução..."

	tmpl.Execute(w, data)

}

//Apresenta um formulário de contato para informar
//e contribuir com críticas e sugestões
func handleContato(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/layout.html", "static/contato.html"))

	data.PageTitle = "Contato"
	data.Title = "Formuláario para Contato"
	data.Output = "Em execução..."

	tmpl.Execute(w, data)

}

//Processar a instalação fácil...
func handleProcessInstall(w http.ResponseWriter, r *http.Request) {

	canal := make(chan string)

	go sendCommand(canal, "Processo de instalação concluído.")
	fmt.Fprintf(w, "%s\n", <-canal)

}

//Veririca se há uma instalação ativa do e-cidade
func handleCheckIfIsInstalled(w http.ResponseWriter, r *http.Request) {

	canal2 := make(chan string)
	go checkIfExist(canal2, "Verificando se existe uma instalação do e-cidade.")

	fmt.Fprintln(w, <-canal2)
}

//Verificar se a pasta do e-cidade existe
func checkIfExist(ch chan string, msg string) {

	resp := connSsh("root", ssh_server, ssh_port, "/root/checkifisinstalled.sh", "/root/.ssh/id_rsa", "")
	ch <- resp
}

//Cria um canal para comunicação entre os processos...
func sendCommand(ch chan string, msg string) {

	resp := connSsh("root", ssh_server, ssh_port, "/root/eii-console", "/root/.ssh/id_rsa", "")

	ch <- resp

}

//Mostra informações sobreo o projeto do instalador fácil
func handleAbout(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/layout.html", "static/sobre.html"))

	data.PageTitle = "Sobre"
	data.Title = "Instalador Fácil"

	tmpl.Execute(w, data)
}


func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandle)
	r.HandleFunc("/processinstall", handleProcessInstall)
	r.HandleFunc("/checkIfInstalled", handleCheckIfIsInstalled)
	r.HandleFunc("/instalar", handleInstall)
	r.HandleFunc("/sobre", handleAbout)
	r.HandleFunc("/contato", handleContato)

	println("Iniciando o instalador fácil ...")

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println("IP do Servidor:", cfg.Section("server").Key("ip_ssh_server").String())
	fmt.Println("Porta do Servidor:", cfg.Section("server").Key("ip_ssh_port").String())

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Section("server").Key("port").String()))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Servidor ouvindo no IP %s\n", cfg.Section("server").Key("ip").String())
	fmt.Println("Usando porta:", listener.Addr().(*net.TCPAddr).Port)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)
	panic(http.Serve(listener, nil))

}
