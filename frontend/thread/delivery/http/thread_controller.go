package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/tsrnd/go-clean-arch/services/crypto"
	"github.com/tsrnd/go-clean-arch/user/usecase"
)

// UserController type
type ThreadController struct {
  Usecase *usecase.UserUsecase,
  Cache *cache.Cache,
}


