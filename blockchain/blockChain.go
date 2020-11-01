package blockchain

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"math/big"
)
var BUCKET_NAME  ="blocks"
var LAST_KEY  ="lasthash"
//
var CHAINDB  = "chain1.db"
var CHAIN BlockChain
/*
  区块链结构体实例定义：用于表示代表一条区块链
  该区块链包含以下功能：
      ①将新产生的区块与已有的区块链接起来，并保存
		②可以查询某个区块的信息
		③可以将所有区块进行遍历，输出区块信息(打印)
 */
type BlockChain struct {
	LastHash []byte  //最新区块的hash
	BoltDb *bolt.DB
}

//查询所有区块信息，并返回，将所有的区块放入到切片中
func (bc BlockChain)QueryAllBlocks()[]*Block  {
	blocks:=make([]*Block,0)
	db:=bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			panic("查询数据出错")

		}
		eachKey:=bc.LastHash
		preHashBig:=new(big.Int)
		zeroBig:=big.NewInt(0)
		for    {
			eachBlockBytes:=bucket.Get(eachKey)
			//反序列化以后得到的每一个区块
			eachBlock,_:=DeSerialize(eachBlockBytes)
			//将遍历到的每一个区块结构体指针放入到【】byte容器中
			blocks=append(blocks,eachBlock)
			preHashBig.SetBytes(eachBlock.PrevHash)

			if preHashBig.Cmp(zeroBig)==0 {
			//判断区块链遍历算法已到创世区块，如已到则跳出循环
				break
			}//否则，继续向前遍历
			eachKey=eachBlock.PrevHash
		}
		return nil
	})

	return blocks
}



//通过区块的高度查询某个具体的区块信息，返回区块实例
func (bc BlockChain) QueryBlockByHeight(height int64) *Block {
	if height<0 {
		//如果目标高度小于零，则说明参数不合法
		return nil

	}
	var block *Block
	db:=bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			panic("查询数据失败！")
		}
		hashKey:=bc.LastHash
		for    {
			lastBlockBytes:=bucket.Get(hashKey)
			eachBlock,_:=DeSerialize(lastBlockBytes)
			if eachBlock.Height<height {
				//查询给定的数字超出区块链中的最新高度
				break
			}
			if eachBlock.Height==height {
				//高度与目标一致，已经找到目标区块
				block=eachBlock
				break
			}
			//当前遍历的区块高度与目标高度不一致，继续遍历
			//以eachBlock.PrevHash为key。使用Get获取上一个区块的数据

			hashKey=eachBlock.PrevHash
		}

		return nil
	})

	return block
}


/*
	创建一条区块链，并返回实例
		解释：由于区块链就是由一个一个的区块组成的，因此
 */
func NewBlockChain() BlockChain  {
	//打开存储区块数据的chain.db文件
	db,err:=bolt.Open(CHAINDB,0600,nil)
	if err!=nil {
		panic(err.Error())
	}

	var bl BlockChain
	//先从区块链中读，判断是否存在创世区块
	db.Update(func(tx *bolt.Tx) error {

		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			bucket,err=tx.CreateBucket([]byte(BUCKET_NAME))
			if err!=nil {
				panic(err.Error())
			}
		}
		lastHash:=bucket.Get([]byte(LAST_KEY))
		if len(lastHash) == 0 {
			//①、创世区块
			genesis:=CreateGenesisBlock()
			//②、创建一个存储区块数据的文件
			fmt.Printf("genesis的Hash值:%x\n",genesis.Hash)
			bl=BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
			genesisBytes,_:=genesis.Serialize()
			bucket.Put(genesis.Hash,genesisBytes)
			bucket.Put([]byte(LAST_KEY),genesis.Hash)
		}else {//有创世区块
			lastHash:=bucket.Get([]byte(LAST_KEY))
			lastBlockBytes:=bucket.Get(lastHash)
			lastBlock,err:=DeSerialize(lastBlockBytes)
			if err != nil {
				panic("读取区块链数据失败")
			}
			bl=BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}
		}
		return nil
	})
	//为全局赋值
		CHAIN=bl
	return bl
}
//该方法用于根据用户传入的认证id查询区块的信息，并返回
func (bc BlockChain) QueryBlockByCertId(cert_id []byte) (*Block,error) {
	var block *Block
	db:=bc.BoltDb
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			err=errors.New("查询区块数据遇到错误！")
			fmt.Println(err)
			return err
		}
		//否则桶存在
		eachHash:=bucket.Get([]byte(LAST_KEY))
		eachBig:=new(big.Int)
		zeroBig:=big.NewInt(0)
		for   {
			//eachHash:=bucket.Get(eachKey)
			eachBlockBytes:=bucket.Get(eachHash)
			eachBlock,_:=DeSerialize(eachBlockBytes)
			if string(eachBlock.Data)==string(cert_id) {
				block=eachBlock
				break
				eachBig.SetBytes(eachBlock.PrevHash)
				if eachBig.Cmp(zeroBig)==0 {
					break
				}
			}
			eachHash=eachBlock.PrevHash
		}
		return nil
	})
	return block,err

}




	//③、把新创建的创世区块存入到chain.db中的一个桶中
	//db.Update(func(tx *bolt.Tx) error {
	//	bucket,err:=tx.CreateBucket([]byte(BUCKET_NAME))
	//	if err!=nil {
	//		panic(err.Error())
	//
	//	}
	//	//将创世区块保存到桶中
	//	serialBlock,err:=genesis.Serialize()
	//	if err!=nil {
	//		panic(err.Error())
	//	}
	//	//把创世区块存入到桶中
	//	bucket.Put(genesis.Hash,serialBlock)
	//	//更新指向最新区块的hash值
	//	bucket.Put([]byte(LAST_KEY),genesis.Hash)
	//	bl.LastHash=genesis.Hash
	//	return nil
	//})


/*
  调用BlockChain的SaveBlock方法，该方法将一个生成的区块保存到chain.db文件中
 */
func (bc BlockChain) SaveData(data []byte) (Block,error) {
	db:=bc.BoltDb
	var e  error
	var lastBlock *Block

	//先查询chain.db中存储的最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket ==nil {
			e=errors.New("boltdb未创建，请重试")
			return e
		}
		//lastHash:=bucket.Get([]byte(LAST_KEY))
		lastBlockBytes:=bucket.Get(bc.LastHash)
		lastBlock,_=DeSerialize(lastBlockBytes)

		return nil
	})
	//先生成一个区块，把data存入到新生成的区块中
	newBlock:=NewBlock(lastBlock.Height+1,data,lastBlock.Hash)

	//更新chain.db,把newblock存入bolddb中
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		//key=hash,value=block的byte
		//区块序列化
		newBlockBytes,_:=newBlock.Serialize()
		//把区块信息保存到boltdb中
		bucket.Put(newBlock.Hash,newBlockBytes)
		//更新代表最后一个区块hash值的记录
		bucket.Put([]byte(LAST_KEY),newBlock.Hash)
        bc.LastHash=newBlock.Hash
		return nil
	})

	return  newBlock,e


	//更新chain.db
	//db.Update(func(tx *bolt.Tx) error {
	//	var tong *bolt.Bucket
	//	tong=tx.Bucket([]byte(BUCKET_NAME))
	//		if tong==nil {
	//
	//		}
	//	tong, err := tx.CreateBucket([]byte(BUCKET_NAME))
	//	if err != nil {
	//		return err
	//	}
		//先查看获取桶中

		//	lastBlock:=tong.Get([]byte(LAST_KEY))
		//	blockHash,err:=block.Serialize()
		//	if err!=nil {
		//		return err
		//	}
		//	if lastBlock==nil {
		//		tong.Put(block.Hash,blockHash)
		//		tong.Put([]byte(LAST_KEY),blockHash)
		//	}
			//tong.Put([]byte("0"),[]byte("a"))
			//return nil


}