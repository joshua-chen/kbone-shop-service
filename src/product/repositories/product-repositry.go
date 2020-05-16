package repositories

import (
	"errors"
	"sync"

	"product/datamodels"
)

// Query代表一种“访客”和它的查询动作。
type Query func(datamodels.Product) bool

// ProductRepository会处理一些关于product实例的基本的操作 。
// 这是一个以测试为目的的接口，即是一个内存中的product库
// 或是一个连接到数据库的实例。
type ProductRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (product datamodels.Product, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Product)

	InsertOrUpdate(product datamodels.Product) (updatedProduct datamodels.Product, err error)
	Delete(query Query, limit int) (deleted bool)
}

// NewProductRepository返回一个新的基于内存的product库。
// 库的类型在我们的例子中是唯一的。
func NewProductRepository(source map[int64]datamodels.Product) ProductRepository {
	return &productMemoryRepository{source: source}
}

// productMemoryRepository就是一个"ProductRepository"
// 它负责存储于内存中的实例数据(map)
type productMemoryRepository struct {
	source map[int64]datamodels.Product
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *productMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, product := range r.source {
		ok = query(product)
		if ok {
			if action(product) {
				loops++
				if actionLimit >= loops {
					break // break
				}
			}
		}
	}

	return
}

// Select方法会收到一个查询方法
// 这个方法给出一个单独的product实例
// 直到这个功能返回为true时停止迭代。
//
// 它返回最后一次查询成功所找到的结果的值
// 和最后的product模型
// 以减少caller之间的通信
//
// 这是一个很简单但很聪明的雏形方法
// 我基本在所有会用到的地方使用自从我想到了它
// 也希望你们觉得好用
func (r *productMemoryRepository) Select(query Query) (product datamodels.Product, found bool) {
	found = r.Exec(query, func(m datamodels.Product) bool {
		product = m
		return true
	}, 1, ReadOnlyMode)

	// set an empty datamodels.Product if not found at all.
	if !found {
		product = datamodels.Product{}
	}

	return
}

// SelectMany作用相同于Select但是它返回一个切片
// 切片包含一个或多个实例
// 如果传入的参数limit<=0则返回所有
func (r *productMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.Product) {
	r.Exec(query, func(m datamodels.Product) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

// InsertOrUpdate添加或者更新一个product实例到（内存）储存中。
//
// 返回最新操作成功的实例或抛出错误。
func (r *productMemoryRepository) InsertOrUpdate(product datamodels.Product) (datamodels.Product, error) {
	id := product.ID

	if id == 0 { // 创建一个新的操作
		var lastID int64
		// 找到最大的ID，避免重复。
		// 在实际使用时您可以使用第三方库去生成
		// 一个string类型的UUID
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		product.ID = id

		// map-specific thing
		r.mu.Lock()
		r.source[id] = product
		r.mu.Unlock()

		return product, nil
	}

	// 更新操作是基于product.ID的，
	// 在例子中我们允许了对poster和genre的更新（如果它们非空）。
	// 当然我们可以只是做单纯的数据替换操作:
	// r.source[id] = product
	// 并注释掉下面的代码;
	current, exists := r.Select(func(m datamodels.Product) bool {
		return m.ID == id
	})

	if !exists { // 当ID不存在时抛出一个error
		return datamodels.Product{}, errors.New("failed to update a nonexistent product")
	}

	// map-specific thing
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()

	return product, nil
}

func (r *productMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.Product) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
