package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

var rootCmd = &cobra.Command{
	Use:   "gbase",
	Short: "A CLI for managing Go projects",
}

var newCmd = &cobra.Command{
	Use:   "new [project-name or .]",
	Short: "Cria um novo projeto Go com o workspace e configurações necessárias",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		destDir := setupProjectDirectory(projectName)

		// Clonar os repositórios
		repos := []string{
			"https://github.com/Tutuacs/cmd.git",
			"https://github.com/Tutuacs/pkg.git",
			"https://github.com/Tutuacs/internal.git",
		}
		for _, repo := range repos {
			cloneRepo(repo, destDir)
		}
		initGitRepo(destDir)

		// Criar os arquivos de configuração
		createGoWork(destDir)
		createGitIgnoreFile(destDir)
		createEnvFile(destDir)
		createDockerComposeFile(destDir)
		createMakefile(destDir, projectName)
	},
}

// Comando `gbase g h [name]` ou `gbase generate h [name]`
var generateCmd = &cobra.Command{
	Use:   "generate [type] [name]",
	Short: "Gera arquivos baseados no tipo especificado",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fileType := args[0]
		name := args[1]

		if fileType != "h" && fileType != "handler" {
			log.Fatalf(Red+"Tipo não suportado: %s. Use 'h' para handler."+Reset, fileType)
		}

		// Verifica se está dentro de um workspace válido
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf(Red+"Erro ao obter diretório atual: %v"+Reset, err)
		}

		// Verifica se a pasta 'internal' existe
		internalDir := filepath.Join(currentDir, "internal")
		if _, err := os.Stat(internalDir); os.IsNotExist(err) {
			log.Fatalf(Red + "Você precisa estar na pasta do workspace para gerar handlers." + Reset)
		}

		// Criar os arquivos de handler
		handlerDir := filepath.Join(internalDir, name)
		if err := os.MkdirAll(handlerDir, os.ModePerm); err != nil {
			log.Fatalf(Red+"Erro ao criar diretório de handler: %v"+Reset, err)
		}

		createHandlerFiles(handlerDir, name)
	},
}

func init() {
	// Adicionar os comandos `new` e `generate` ao rootCmd
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(generateCmd)

	// Também permitir `gbase g` como alias para `generate`
	rootCmd.AddCommand(&cobra.Command{
		Use:   "g [type] [name]",
		Short: "Alias para `generate`",
		Args:  cobra.ExactArgs(2),
		Run:   generateCmd.Run,
	})
}

// Função auxiliar para configurar o diretório do projeto
func setupProjectDirectory(projectName string) string {
	destDir := ""
	if projectName == "." {
		var err error
		destDir, err = os.Getwd()
		if err != nil {
			log.Fatalf(Red+"Erro ao obter diretório atual: %v"+Reset, err)
		}
		projectName = filepath.Base(destDir)
	} else {
		var err error
		destDir, err = os.Getwd()
		if err != nil {
			log.Fatalf(Red+"Erro ao obter diretório atual: %v"+Reset, err)
		}
		destDir = filepath.Join(destDir, projectName)
	}

	fmt.Printf(Yellow+"Criando workspace para %s em %s\n"+Reset, projectName, destDir)
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		log.Fatalf(Red+"Erro ao criar diretório do projeto: %v"+Reset, err)
	}
	return destDir
}

// Função auxiliar para criar arquivos do handler
func createHandlerFiles(handlerDir, name string) {
	handlerContent := fmt.Sprintf(`package %s

import (
	"net/http"

	"github.com/Tutuacs/pkg/routes"
)

type Handler struct {
	subRoute string
}

func NewHandler() *Handler {
	return &Handler{subRoute: "/%s"}
}

func (h *Handler) BuildRoutes(router routes.Route) {
	// TODO implement the routes call
	router.NewRoute(routes.POST, h.subRoute, h.create)
	router.NewRoute(routes.GET, h.subRoute, h.list)
	router.NewRoute(routes.GET, h.subRoute+"/{id}", h.getById)
	router.NewRoute(routes.PUT, h.subRoute+"/{id}", h.update)
	router.NewRoute(routes.DELETE, h.subRoute+"/{id}", h.delete)
}

// ! Recommended private functions
// * Create stores to get DB data like this 
/*
	store, err := NewStore()
	if err != nil {
		return
	}

	defer store.CloseStore()

	* Use resolver to getParams, getBody and writeResponse

*/

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {

}
`, name, name)
	storeContent := fmt.Sprintf(`package %s

import (
	"database/sql"

	"github.com/Tutuacs/pkg/db"
)

type Store struct {
	db.Store
	db      *sql.DB
	extends bool
	Table   string
}

func NewStore(conn ...*sql.DB) (*Store, error) {
	if len(conn) == 0 {

		con, err := db.NewConnection()

		db.NewConnection()

		return &Store{
			db:      con,
			extends: false,
		}, err
	}

	return &Store{
		db:      conn[0],
		extends: true,
	}, nil
}

func (s *Store) CloseStore() {
	if !s.extends {
		s.db.Close()
	}

	// db.ScanRow()
}

func (s *Store) GetConn() *sql.DB {

	return s.db
}


// TODO: Implement the store consults`, name)

	id := `"id"`
	json := fmt.Sprintf("`json:%s`", id)

	nameTitle := cases.Title(language.Und).String(name)

	typesContent := fmt.Sprintf(`package %s

	// TODO: Create types and dtos for %s

type %s struct {
	ID int64 %s
}
	
type New%sDto struct {

}

type Update%sDto struct {

}
	
	`, name, name, nameTitle, json, nameTitle, nameTitle)

	// Criar arquivos handler.go, store.go e types.go
	files := map[string]string{
		"handler.go": handlerContent,
		"store.go":   storeContent,
		"types.go":   typesContent,
	}

	for fileName, content := range files {
		filePath := filepath.Join(handlerDir, fileName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			log.Fatalf(Green+"Erro ao criar arquivo %s: %v"+Reset, fileName, err)
		}
		fmt.Printf(Green+"Arquivo %s criado em %s\n"+Reset, fileName, handlerDir)
	}
}

func cloneRepo(repo string, dest string) {
	// Comando para clonar o repositório
	cmd := exec.Command("git", "clone", repo)
	cmd.Dir = dest
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf(Red+"Erro ao clonar repositório %s: %v\nOutput: %s"+Reset, repo, err, string(output))
	}
	fmt.Printf(Green+"Clonado %s com sucesso\n"+Reset, repo)

	// Caminho do diretório clonado
	repoDir := filepath.Join(dest, filepath.Base(repo))
	repoDir = repoDir[:len(repoDir)-4] // Remover a extensão ".git" do nome do repositório

	// Remover a pasta .git
	gitDir := filepath.Join(repoDir, ".git")
	err = os.RemoveAll(gitDir)
	if err != nil {
		log.Fatalf(Red+"Erro ao remover a pasta .git de %s: %v"+Reset, repoDir, err)
	}
	fmt.Printf(Yellow+"Removida a pasta .git de %s\n"+Reset, repoDir)

	// Extrair o nome do módulo da URL do repositório
	repoName := filepath.Base(repo)
	repoName = repoName[:len(repoName)-4] // Remover ".git" do nome do repositório

	// Executar go mod init
	cmd = exec.Command("go", "mod", "init", "github.com/Tutuacs/"+repoName)
	cmd.Dir = repoDir
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf(Red+"Erro ao executar 'go mod init' em %s: %v\nOutput: %s"+Reset, repoDir, err, string(output))
	}
	fmt.Printf(Green+"'go mod init' executado com sucesso em %s\n"+Reset, repoDir)

	// Executar go mod tidy
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = repoDir
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf(Red+"Erro ao executar 'go mod tidy' em %s: %v\nOutput: %s"+Reset, repoDir, err, string(output))
	}
	fmt.Printf(Green+"'go mod tidy' executado com sucesso em %s\n"+Reset, repoDir)
}

func initGitRepo(dest string) {
	cmd := exec.Command("git", "init")
	cmd.Dir = dest
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf(Red+"Erro ao inicializar o repositório git no diretório %s: %v\nOutput: %s"+Reset, dest, err, string(output))
	}
	fmt.Println(Green + "Repositório Git inicializado no diretório root do projeto" + Reset)
}

func createGoWork(dest string) {
	// Inicializa o workspace com go work
	cmd := exec.Command("go", "work", "init")
	cmd.Dir = dest
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf(Red+"Erro ao inicializar o workspace: %v\nOutput: %s"+Reset, err, string(output))
	}
	fmt.Println(Green + "Workspace inicializado com sucesso" + Reset)

	// Adiciona os diretórios cmd, pkg e internal ao workspace
	dirs := []string{"./cmd", "./pkg", "./internal"}
	for _, dir := range dirs {
		cmd = exec.Command("go", "work", "use", dir)
		cmd.Dir = dest
		output, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatalf(Red+"Erro ao adicionar %s ao workspace: %v\nOutput: %s"+Reset, dir, err, string(output))
		}
		fmt.Printf(Green+"%s adicionado ao workspace com sucesso\n"+Reset, dir)
	}
}

func createGitIgnoreFile(dest string) {
	gitIgnoreContent := `go*.sum
/bin
`
	gitIgnoreFilePath := filepath.Join(dest, ".gitignore")
	if err := os.WriteFile(gitIgnoreFilePath, []byte(gitIgnoreContent), 0644); err != nil {
		log.Fatalf(Red+"Erro ao criar arquivo .gitignore: %v"+Reset, err)
	}
	fmt.Println(Green + "Arquivo .gitignore criado com sucesso" + Reset)
}

func createEnvFile(dest string) {
	envContent := `# API Configuration
API_PORT=":9000"

# Database Configuration
DB_HOST="127.0.0.1"
DB_PORT="9999"
DB_ADDR="127.0.0.1:9999"
DB_USER="user"
DB_PASS="pass"
DB_NAME="defaultDb"

# JWT Configuration
JWT_EXP=604800  # 3600*24*7 (7 dias em segundos)
JWT_SECRET="secret"
`

	envFilePath := filepath.Join(dest, ".env")
	if err := os.WriteFile(envFilePath, []byte(envContent), 0644); err != nil {
		log.Fatalf(Red+"Erro ao criar arquivo .env: %v"+Reset, err)
	}
	fmt.Println(Green + "Arquivo .env criado com sucesso" + Reset)
}

func createDockerComposeFile(dest string) {
	dockerComposeContent := `services:
  default_postgres:
    container_name: default_postgres
    image: postgres:13.5
    environment:
      POSTGRES_DB: $DB_NAME
      POSTGRES_USER: $DB_USER
      POSTGRES_PASSWORD: $DB_PASS
    ports:
      - "$DB_PORT:5432"
`

	dockerComposeFilePath := filepath.Join(dest, "docker-compose.yaml")
	if err := os.WriteFile(dockerComposeFilePath, []byte(dockerComposeContent), 0644); err != nil {
		log.Fatalf(Red+"Erro ao criar arquivo docker-compose.yaml: %v"+Reset, err)
	}
	fmt.Println(Green + "Arquivo docker-compose.yaml criado com sucesso" + Reset)
}

func createMakefile(dest string, projectName string) {
	makefileContent := `build:
	@go build -o bin/%s cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/%s

migration:
	@go run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
`

	// Replace $(Project_Name) with the actual project name
	makefileContent = fmt.Sprintf(makefileContent, projectName, projectName)

	makefilePath := filepath.Join(dest, "Makefile")
	if err := os.WriteFile(makefilePath, []byte(makefileContent), 0644); err != nil {
		log.Fatalf(Red+"Erro ao criar arquivo Makefile: %v"+Reset, err)
	}
	fmt.Println(Green + "Arquivo Makefile criado com sucesso" + Reset)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
