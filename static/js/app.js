document.addEventListener("DOMContentLoaded", function(event){

    
    var box = document.getElementsByClassName('boxinfo')[0]
    const progresso = document.getElementById("progresso");
    progresso.style.display = "none";


    const btnInstall = document.querySelector("#btnInstall").addEventListener("click", function(event){
       

        box.innerHTML = "Processando a instalação dos pacotes..."
        
        progresso.style.display = "block"

        setTimeout( function(){

            fetch("/checkIfInstalled")
            .then(resp => resp.text())
            .then(resp => {
                if ( resp == "true" ) {
                  console.log(resp)
                  box.innerHTML = "O e-cidade já está instalado em seu servidor." 
                } else {

                    box.innerHTML = "Processo de instalação iniciado no servidor... <br>Por favor, aguarde a conclusão da instalação." 

                    setTimeout(() => {

                        fetch("/processinstall")
                        .then( resp => resp.text())
                        .then(resp =>{
                            box.innerHTML = resp 
                            progresso.style.display = "none"
                        })
                        .catch( err => {
                            box.innerHTML = err 
                        })
                        
                    }, 4000);
                    
                }

               

            })
            
            
        }, 2000)
    })

})

/**
 * 
 * .then(resp => {
                if ( resp.ok ){
                if( resp.text() == "false"){
                    fetch("/processinstall")
                        .then( resp => resp.text())
                        .then(resp =>{
                            box.innerHTML = resp 
                        })
                        .catch( err => {
                            box.innerHTML = err 
                        })
                }
            }})
 */

