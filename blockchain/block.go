package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Data          []byte
	BlockHash     []byte
	PrevBlockHash []byte
	Nonce         int
}

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevBlockHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.BlockHash = hash[:]
// }

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		BlockHash:     []byte{},
		Data:          []byte(data),
		PrevBlockHash: prevHash,
		Nonce:         0,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.BlockHash = hash[:]

	return block
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.BlockHash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
