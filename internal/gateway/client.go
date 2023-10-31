import "github.com.br/devfullcycle/fc-ms-wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Get(client *entity.Client) error
}