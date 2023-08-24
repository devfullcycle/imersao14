# Imersão Full Stack & FullCycle - Sistema de rastreabilidade de veículos

## Descrição

Repositório do front-end feito com Next.js (Front-end)

## Google API Key

Gere uma API Key no painel do Google Console, é necessário ativar o serviço de `Google Maps API`.


## Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):
```
127.0.0.1 host.docker.internal
```
Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.

## Rodar a aplicação

Execute os comandos:

```
docker-compose up
```

Em outro terminal rode o comando para entrar no container do Docker:
```bash
docker compose exec app bash
```

E configure as dependências iniciais (rode somente dentro do container do Docker):
```
npm install
```

Rode a aplicação (rode somente dentro do container do Docker):
```bash
npm run dev
```

Acessar http://localhost:3001.

### Para Windows 

Lembrar de instalar o WSL2 e Docker. Vejo o vídeo: [https://www.youtube.com/watch?v=rpvjVtUPnOc](https://www.youtube.com/watch?v=rpvjVtUPnOc) 

Siga o guia rápido de instalação: [https://github.com/codeedu/wsl2-docker-quickstart](https://github.com/codeedu/wsl2-docker-quickstart) 
