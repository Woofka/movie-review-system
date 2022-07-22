package storage

import (
	"fmt"
	"sort"
)

var data map[uint]*Review

func init() {
	data = make(map[uint]*Review)
}

func List() []*Review {
	res := make([]*Review, 0, len(data))
	for _, v := range data {
		res = append(res, v)
	}
	sort.Slice(res, func(i, j int) bool { return res[i].id < res[j].id })
	return res
}

func Add(r *Review) error {
	if _, ok := data[r.GetId()]; ok {
		return fmt.Errorf("review with id %d already exists", r.GetId())
	}
	data[r.GetId()] = r
	return nil
}

func Get(id uint) (Review, error) {
	r, ok := data[id]
	if !ok {
		return Review{}, fmt.Errorf("review with id %d does not exists", id)
	}
	return *r, nil
}

func Update(r *Review) error {
	if _, ok := data[r.GetId()]; !ok {
		return fmt.Errorf("review with id %d does not exists", r.GetId())
	}
	data[r.GetId()] = r
	return nil
}

func Delete(id uint) error {
	if _, ok := data[id]; !ok {
		return fmt.Errorf("review with id %d does not exists", id)
	}
	delete(data, id)
	return nil
}
