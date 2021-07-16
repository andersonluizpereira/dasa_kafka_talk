#Inicie o docker com
npm run up

#Entra no terminal e execute - docker exec -it dasa_example_kafka_1 bash

#Como criar um tópico  - com uma partição
kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=1

#Como enviar mensagens simples por dentro do kafka
##No mesmo terminal, que você criou o tópico - execute o comando abaixo.
kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

#Como ver as mensagens do tópico
##Em outro terminal - execute  - docker exec -it dasa_example_kafka_1 bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

#Como ver a anatomia do tópico
##Em outro terminal - execute  - docker exec -it dasa_example_kafka_1 bash
kafka-topics --bootstrap-server localhost:9092 --describe --topic=teste

#Agora vamos para baguncinhas rs - https://www.youtube.com/watch?v=hxLzOdeNIyg
#Como alterar as partições de um tópico
#No mesmo terminal que está aberto acima, execute o comando abaixo
kafka-topics --bootstrap-server localhost:9092 --alter --topic=teste --partitions=3

#Abra 4 terminais - execute  - docker exec -it dasa_example_kafka_1 bash
##Execute o comando abaixo
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste
#Viu que replicou, não ficou tudo igual certo rs.

#Curiosidade 
#Para um consumer ler as mensagens desde o começo digite  --from-beginning
kafka-console-consumer --bootstrap-server localhost:9092 --topic=teste --from-beginning

#Nos mesmos 4 terminais - execute  - docker exec -it dasa_example_kafka_1 bash
##Execute o comando abaixo
kafka-console-consumer --bootstrap-server localhost:9092 --topic=teste --group=x
#Viu foi para cada terminal diferente

#Como ver a distribuição das mensagens
kafka-consumer-groups --bootstrap-server=localhost:9092  --describe --all-groups

#Na APP
#teste do consumer
Nos mesmos 4 terminais abertos, saia do docker entre na pasta '/Users/andersonluizpereira/kafka/dasa_example', vá executando e modificando dasa-consumer para 1,2,3
execute com go run main.go

Mate o consumer default dasa-consumer, espere Consumer group 'dasa-group' fazer o rebalanceamento.

"client.id":"dasa-consumer"