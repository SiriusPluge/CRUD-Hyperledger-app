# crud-hyperled-fabric 
## Application - является сервером для обработки запросов по HTTP.
## В приложении используется "Hyperledger fabric"
## В качестве хранилища используется CouchDB.
_________________

### Для запуска проекта необходимо:

1) Скачать исходники    
    - ```mkdir -p $HOME/go-projects/```
    - ```cd $HOME/go-projects/```
    - ```git clone git@github.com:SiriusPluge/CRUD-Hyperledger-app.git```
    - ```git clone git@github.com:SiriusPluge/CRUD-Hyperledger-chaincode.git```
    - ```curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh```
    - ```./install-fabric.sh```
2) Запустить тестовую сеть
    - ```cd $HOME/go-projects/fabric-samples/test-network```
    - ```./network.sh down```
    - ```./network.sh up createChannel -c mychannel -ca -s couchdb```
    - ```./network.sh deployCC -ccn private -ccp ../../CRUD-Hyperledger-chaincode -ccl go```
3) Запустить сервер
    - ```cd ../../CRUD-Hyperledger-app```
    - ```docker-compose up```

### Для тестирования приложения в корне папки проекта "CRUD-Hyperledger-app" будет расположена postman-коллекция.
Перед отправлением запросов необходимо заполнить данные конфигурации:
1) 


sudo chown -R pluge:pluge ./organizations/peerOrganizations/org1.example.com/peers/

### Если вам встретятся ошибки следующего содержания, то необходимо:
1) 