# Golang

***Install go to run programs:***
     
    sudo apt-get install go(for debian,ubuntu)
    sudo yum install go(for red-hat,fedora,suse)

__________________________________________________________________________________________________________________________________________


**Basic program for aes encryption with golang crypto library.**

**aes_encryption.go**

Steps:

    1)Write the word you want to encrypt and save your key and your ciphertext in file.
    2)For decryption you need to have the key,and the encrypted text.You will get a hex code,which wil contain  the plain text!!
  
  Try this site for hex decoding:
        
      (https://www.online-toolz.com/tools/decode-hex.php)
      
      
___________________________________________________________________________________________________________________________________________


**REST-API with gorilla mux router.** 

**api.go**

Steps:

    1)go get -u github.com/gorilla/mux
    2)Start server with go run api.go .
    3)Download postman-client.
    4)Open postman and make get,post,put,and delete requests.

 
___________________________________________________________________________________________________________________________________________

**Simple api with http package.**

**simple_api.go**

Steps:
  
    1)Start server.go with go run and the programs name.
    2)Open postman-client for requests.


___________________________________________________________________________________________________________________________________________

**Extremely fast scanner with go channels**

**scan.go**



Features:
          
     Scans all tcp/udp 65535 ports and finds for open ports ...
     Scans for spelcific tcp port ...
     Finds ip of a domain name ...
     Lists all ports for known services ...
 
Arguments:

     --help(Shows help menu)
     --get_ip(Get ip of host)
     --port(Scan for specific tcp port)
     --get_port()
 

Run:
    
    go run scan.go(executable wiil be stored at /tmp)
               
Build and Run:
      
    go build scan.go && go run scan.go




