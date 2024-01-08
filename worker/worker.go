package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hrand1005/ZoneScaler/common"
	"github.com/phuslu/log"
)

const (
	maxRetries    = 30
	retryInterval = 1 * time.Second
)

type Worker struct {
	coordinatorAddr string
	// static data about the worker
	node *common.GameNode
	// partitions assigned to this worker by the coordinator
	partitions []*common.Partition
}

func New(conf Config) *Worker {
	node := &common.GameNode{
		ID:            uuid.NewString(),
		Address:       fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Load:          0,
		Regions:       []string{"region1", "region2"},
		IsActive:      true,
		LastHeartbeat: time.Now(),
	}

	coordinatorAddr := fmt.Sprintf("%v:%v", os.Getenv("COORDINATOR_HOST"), os.Getenv("COORDINATOR_PORT"))
	return &Worker{
		node:            node,
		coordinatorAddr: coordinatorAddr,
	}
}

func (w *Worker) Start() error {
	var err error
	// register yourself with the coordinator
	registered := false
	for try := 0; try < maxRetries; try++ {
		err = w.registerWithCoordinator()
		if err == nil {
			registered = true
			break
		}
		time.Sleep(retryInterval)
	}
	if !registered {
		log.Fatal().Err(err).Msg("Failed to register with the coordinator")
		return err
	}

	// TODO: request game partitions, possibly in a separate goroutine
	// so that the node can continue to take on additional work so long as
	// it has the capacity to do so

	// handle player interactions for our partition
	// TODO: again, we can embed an http server object in the Worker if
	// further configurability is required
	http.HandleFunc("/player", PlayerDataHandler)
	return http.ListenAndServe(w.node.Address, nil)
}

// RegisterWithCoordinator registers the worker with the coordinator
func (w *Worker) registerWithCoordinator() error {
	jsonData, err := json.Marshal(w.node)
	if err != nil {
		log.Error().Msgf("failed to marshal worker info: %v", err)
		return err
	}

	addNodeURL := fmt.Sprintf("http://%v/nodes", w.coordinatorAddr)
	resp, err := http.Post(addNodeURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("failed to register with coordinator")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("registration failed. Coordinator responded with status code: %d", resp.StatusCode)
		log.Error().Int("status_code", resp.StatusCode).Msg(msg)
		return fmt.Errorf(msg)
	}

	log.Info().Msg("Successfully registered with the coordinator")
	return nil
}
