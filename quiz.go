package main

import (
    "fmt"
)
       
var sf = fmt.Scanf

local lsp_zero = require('lsp-zero')

lsp_zero.on_attach(function(client, bufnr)
  -- see :help lsp-zero-keybindings
  -- to learn the available actions
  lsp_zero.default_keymaps({buffer = bufnr})
end)

-- here you can setup the language serverlocal lsp_zero = require('lsp-zero')

func main(){
    var name1 string
    var age int
    var country string
    var err error
    
      



    pl("Enter your detail here:")
    pl("What is your name?")
    if _, err = sf("%s", &name1); err != nil{
        pl ("Error Reading name", err)
        return
    }


    pl("What is your age?")
    if _, err = sf("%d", &age); err != nil{
        pl("Error Reading age", err)

    }


    pl("What is your Country Name:?")
   if _, err = sf("%s", &country); err != nil{
        pl("Error Reading country", err)
    }




       for i = 0 ; i < 3; i++ {
         pl("Ok Your name is", name1, "of" , age , "from ",country)

    }
}
