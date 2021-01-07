package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
)

const loadersKey = "dataloaders"

//Loaders Loaders
type Loaders struct {
	FileUploadByID FileUploadLoader
}

//Dataloader Dataloader Middleware
func Dataloader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			FileUploadByID: FileUploadLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([][]*model.FileUpload, []error) {

					resp, err := service.FileUploadGetByArrayUserID(context.Background(), ids, nil)

					if err != nil {
						panic(err)
					}

					fileUploadByID := map[int][]*model.FileUpload{}

					for _, val := range resp {
						fileUploadByID[val.UserID] = append(fileUploadByID[val.UserID], val)
					}

					fileUploads := make([][]*model.FileUpload, len(ids))
					for i, id := range ids {
						fileUploads[i] = fileUploadByID[id]
					}

					return fileUploads, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

//For Get Data
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
