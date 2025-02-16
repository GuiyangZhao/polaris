/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package boltdb

import (
	"fmt"
	"sync"
	"time"

	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"

	"github.com/polarismesh/polaris/common/eventhub"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/common/utils"
	"github.com/polarismesh/polaris/store"
)

type maintainStore struct {
	handler BoltHandler
	leMap   map[string]bool
	mutex   sync.Mutex
}

// StartLeaderElection
func (m *maintainStore) StartLeaderElection(key string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.leMap[key]
	if ok {
		return nil
	}
	m.leMap[key] = true
	eventhub.Publish(eventhub.LeaderChangeEventTopic, store.LeaderChangeEvent{Key: key, Leader: true})
	return nil
}

// IsLeader
func (m *maintainStore) IsLeader(key string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	v, ok := m.leMap[key]
	if ok {
		return v
	}
	return false
}

// ListLeaderElections
func (m *maintainStore) ListLeaderElections() ([]*model.LeaderElection, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var out []*model.LeaderElection
	for k, v := range m.leMap {
		item := &model.LeaderElection{
			ElectKey: k,
			Host:     utils.LocalHost,
			Ctime:    0,
			Mtime:    time.Now().Unix(),
			Valid:    v,
		}
		item.CreateTime = time.Unix(item.Ctime, 0)
		item.ModifyTime = time.Unix(item.Mtime, 0)

		out = append(out, item)
	}
	return out, nil
}

// ReleaseLeaderElection
func (m *maintainStore) ReleaseLeaderElection(key string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	v, ok := m.leMap[key]
	if !ok {
		return fmt.Errorf("LeaderElection(%s) not started", key)
	}

	if v {
		m.leMap[key] = false
		eventhub.Publish(eventhub.LeaderChangeEventTopic, store.LeaderChangeEvent{Key: key, Leader: false})
	}

	return nil
}

// BatchCleanDeletedInstances
func (m *maintainStore) BatchCleanDeletedInstances(batchSize uint32) (uint32, error) {
	fields := []string{insFieldValid}
	values, err := m.handler.LoadValuesByFilter(tblNameInstance, fields, &model.Instance{},
		func(m map[string]interface{}) bool {
			valid, ok := m[insFieldValid]
			if ok && !valid.(bool) {
				return true
			}
			return false
		})
	if err != nil {
		return 0, err
	}
	if len(values) == 0 {
		return 0, nil
	}

	var count uint32 = 0
	keys := make([]string, 0, batchSize)
	for k := range values {
		keys = append(keys, k)
		count++
		if count >= batchSize {
			break
		}
	}
	err = m.handler.DeleteValues(tblNameInstance, keys)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (m *maintainStore) GetUnHealthyInstances(timeout time.Duration, limit uint32) ([]string, error) {
	return m.getUnHealthyInstancesBefore(time.Now().Add(-timeout), limit)
}

func (m *maintainStore) getUnHealthyInstancesBefore(mtime time.Time, limit uint32) ([]string, error) {
	fields := []string{insFieldProto, insFieldValid}
	instances, err := m.handler.LoadValuesByFilter(tblNameInstance, fields, &model.Instance{},
		func(m map[string]interface{}) bool {

			valid, ok := m[insFieldValid]
			if ok && !valid.(bool) {
				return false
			}

			insProto, ok := m[insFieldProto]
			if !ok {
				return false
			}

			ins := insProto.(*apiservice.Instance)

			insMtime, err := time.Parse("2006-01-02 15:04:05", ins.GetMtime().GetValue())
			if err != nil {
				log.Errorf("[Store][boltdb] parse instance mtime error, %v", err)
				return false
			}

			if insMtime.Before(mtime) {
				return false
			}

			if !ins.GetEnableHealthCheck().GetValue() {
				return false
			}

			if ins.GetHealthy().GetValue() {
				return false
			}

			return true
		})

	if err != nil {
		log.Errorf("[Store][boltdb] load instance from kv error, %v", err)
		return nil, err
	}

	var instanceIds []string
	var count uint32 = 0
	for _, v := range instances {
		instanceIds = append(instanceIds, v.(*model.Instance).ID())
		count += 1
		if count >= limit {
			break
		}
	}

	return instanceIds, nil
}
