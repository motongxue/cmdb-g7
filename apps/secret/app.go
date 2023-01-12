package secret

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	request "github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
)

const (
	AppName = "secret"
)

var (
	validate = validator.New()
)

func NewDefaultSecret() *Secret {
	return &Secret{
		Data: &CreateSecretRequest{
			RequestRate: 5,
		},
	}
}

func NewSecretSet() *SecretSet {
	return &SecretSet{
		Items: []*Secret{},
	}
}

func (s *SecretSet) Add(item *Secret) {
	s.Items = append(s.Items, item)
}

func NewSecret(req *CreateSecretRequest) (*Secret, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Secret{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Data:     req,
	}, nil
}

func NewCreateSecretRequest() *CreateSecretRequest {
	return &CreateSecretRequest{
		RequestRate: 5,
	}
}

func (req *CreateSecretRequest) Validate() error {
	if len(req.AllowRegions) == 0 {
		return fmt.Errorf("required less one allow_regions")
	}
	return validate.Struct(req)
}

func NewQuerySecretRequestFromHTTP(r *http.Request) *QuerySecretRequest {
	qs := r.URL.Query()

	return &QuerySecretRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{
		Page:     request.NewDefaultPageRequest(),
		Keywords: "",
	}
}

func NewDescribeSecretRequest(id string) *DescribeSecretRequest {
	return &DescribeSecretRequest{
		Id: id,
	}
}

func NewDeleteSecretRequestWithID(id string) *DeleteSecretRequest {
	return &DeleteSecretRequest{
		Id: id,
	}
}

func (s *CreateSecretRequest) AllowRegionString() string {
	return strings.Join(s.AllowRegions, ",")
}

func (s *CreateSecretRequest) LoadAllowRegionFromString(regions string) {
	if regions != "" {
		s.AllowRegions = strings.Split(regions, ",")
	}
}
