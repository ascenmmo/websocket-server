// @tg version=1.0.3
// @tg backend="Asenmmo"
// @tg title=`Ascenmmo Rest API`
// @tg servers=`http://stage.ascenmmo.com;stage cluster`
//
//go:generate go tool github.com/seniorGolang/tg/v2/cmd/tg transport --services . --out ../transport --outSwagger ../swagger.yaml
//go:generate go tool github.com/seniorGolang/tg/v2/cmd/tg client -go --services . --outPath ../../clients/suppliers
//go:generate goimports -l -w ../transport ../../clients/suppliers

package api

import (
	"context"

	"github.com/ascenmmo/websocket-server/pkg/api/types"
)

// @tg http-prefix=api/v1/udp/
// @tg jsonRPC-server log trace
// @tg tagNoOmitempty
type ServerSettings interface {
	// @tg http-headers=token|Token
	// @tg summary=`GetConnectionsNum`
	GetConnectionsNum(ctx context.Context, token string) (countConn int, exists bool, err error)
	// @tg http-headers=token|Token
	// @tg summary=`HealthCheck`
	HealthCheck(ctx context.Context, token string) (exists bool, err error)
	// @tg http-headers=token|Token
	// @tg summary=`GetServerSettings`
	GetServerSettings(ctx context.Context, token string) (settings types.Settings, err error)
	// @tg http-headers=token|Token
	// @tg summary=`CreateRoom`
	CreateRoom(ctx context.Context, token string, createRoom types.CreateRoomRequest) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`GetDeletedRooms`
	GetDeletedRooms(ctx context.Context, token string, ids []types.GetDeletedRooms) (deletedIds []types.GetDeletedRooms, err error)
}
