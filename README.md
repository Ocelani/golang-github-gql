# golang-github-gql

###### Otávio Celani

### Project

It's a project that collects GitHub data through the GraphQL API.
The data collected is about the 1000 most popular repositories. You can find it inside the ./data directory, in JSON and CSV.

### Running

First of all, set a .env file with your GITHUB_TOKEN variable.
Then...

Got two options:

- Run the binary file: ./main
- Run with Go runtime (needs to install Golang)

#### With Binary

##### On an Unix base system (Linux or MacOS)

1. Go to the project root directory.
2. Run on CLI de following command:

```bash
$ ./main
```

### Prática 2

#### Questões

- [x] **RQ.01**
      Quais as características dos top-100repositórios Java mais populares?
- [x] **RQ.02**
      Quais as características dos top-100 repositórios Python mais populares?
- [x] **RQ.03**
      Repositórios Java e Python populares possuem características de “boa qualidade” semelhantes?
- [x] **RQ.04**
      A popularidade influencia nas características dos repositórios Java e Python?

#### Métricas

Utilizaremos como fatores de qualidade métricas associadas à quatro dimensões:

- **Popularidade:**
- - _número de estrelas_, _número de watchers_, _número de forks_ dos repositórios coletados
- **Tamanho:**
- - linhas de código (LOC e SLOC) e linhas de comentários
- **Atividade:**
- - _número de releases_, _frequência de publicação de releases_ (número de releases / dias)
- **Maturidade:**
- - _idade (em anos) de cada repositório coletado_

### Prática 1

#### Questões

- [x] **RQ.01**
      Sistemas populares são mais maduros/antigos?
      _Métrica: idade do repositório (calculado a partir da data de sua criação)_

- [x] **RQ.02**
      Sistemas populares recebem muita contribuição externa?
      _Métrica: total de pullrequests aceitas_

- [x] **RQ.03**
      Sistemas populares lançam releases com frequência?
      _Métrica: total de releases_

- [x] **RQ.04**
      Sistemas populares são atualizados com frequência?
      _Métrica: tempo até a última atualização (calculado a partir da data de última atualização)_

- [x] **RQ.05**
      Sistemas populares são escritos naslinguagens mais populares?
      _Métrica: linguagem primária de cada um desses repositórios_

- [x] **RQ.06**
      Sistemas populares possuem um alto percentual de issuesfechadas?
      _Métrica: razão entre número de issues fechadas pelo total de issues_
