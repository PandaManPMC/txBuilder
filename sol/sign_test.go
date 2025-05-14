package sol

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/associated_token_account"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/mr-tron/base58"
	"log"
	"testing"
)

const clientUrl = "https://api.devnet.solana.com"

func TestSendTransaction(t *testing.T) {
	c := client.NewClient(clientUrl)

	// ✅ 64 字节私钥（base64 示例，你可以用 JSON 数组或 raw byte）
	privateKey := "341f58e5e65307ea627763a2d730deb6da5b95aca75ef5de2a4d916accd2dc3316b887a803bff07bcf4cd46281ba1cb6ada25fbf0419b3a9496cc20631ff058f"
	pk, err := ImportPrivateKeyFromHex(privateKey)
	if nil != err {
		t.Fatal(err)
	}

	// ✅ 构造账户对象
	sender, err := types.AccountFromBytes(pk)
	if err != nil {
		t.Fatal(err)
	}

	// ✅ 获取 recent blockhash
	latestBlockhashResp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// ✅ 构造转账指令（1 SOL = 1_000_000_000 Lamports）
	toAddress := "2qy1811C9mgVquHUrPuuzRrxFrp3Y7wh5QCZzVvQbZda"
	amount := uint64(1_000_000) // 0.001 SOL

	instruction := system.Transfer(system.TransferParam{
		From:   sender.PublicKey,
		To:     common.PublicKeyFromString(toAddress),
		Amount: amount,
	})

	// ✅ 构造并签名交易
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        sender.PublicKey,
			RecentBlockhash: latestBlockhashResp.Blockhash,
			Instructions:    []types.Instruction{instruction},
		}),
		Signers: []types.Account{sender},
	})
	if err != nil {
		t.Fatal(err)
	}

	// ✅ 广播交易
	sig, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("✅ Transaction sent! Signature:", sig)
}

func TestSimulateTransaction(t *testing.T) {
	c := client.NewClient(clientUrl)

	// ✅ 64 字节私钥（base64 示例，你可以用 JSON 数组或 raw byte）
	privateKey := "341f58e5e65307ea627763a2d730deb6da5b95aca75ef5de2a4d916accd2dc3316b887a803bff07bcf4cd46281ba1cb6ada25fbf0419b3a9496cc20631ff058f"
	pk, err := ImportPrivateKeyFromHex(privateKey)
	if nil != err {
		t.Fatal(err)
	}

	sender, err := types.AccountFromBytes(pk)
	if err != nil {
		t.Fatal(err)
	}

	receiver := common.PublicKeyFromString("2qy1811C9mgVquHUrPuuzRrxFrp3Y7wh5QCZzVvQbZda")

	// 获取 blockhash
	blockhashResp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// 构造转账指令
	instruction := system.Transfer(system.TransferParam{
		From:   sender.PublicKey,
		To:     receiver,
		Amount: 10000, // 0.00001 SOL
	})

	// 构建交易
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        sender.PublicKey,
			RecentBlockhash: blockhashResp.Blockhash,
			Instructions:    []types.Instruction{instruction},
		}),
		Signers: []types.Account{sender},
	})
	if err != nil {
		log.Fatal("构建交易失败:", err)
	}

	// 模拟交易
	simResult, err := c.SimulateTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal("模拟交易失败:", err)
	}

	fmt.Println("=== 模拟交易结果 ===")
	fmt.Printf("是否成功: %v\n", simResult.Err == nil)
	fmt.Printf("计算单位消耗 (units): %v\n", simResult.UnitConsumed)
	fmt.Printf("日志: \n")
	for _, logLine := range simResult.Logs {
		fmt.Println(logLine)
	}

	unitPrice := 5000 // lamports per unit
	units := *simResult.UnitConsumed

	lamportsUsed := units * uint64(unitPrice)
	solUsed := float64(lamportsUsed) / 1_000_000_000.0

	fmt.Printf("模拟交易消耗大约: %.6f SOL\n", solUsed)
}

func TestGetFeeForMessage(t *testing.T) {
	c := client.NewClient(clientUrl)

	// ✅ 64 字节私钥（base64 示例，你可以用 JSON 数组或 raw byte）
	privateKey := "341f58e5e65307ea627763a2d730deb6da5b95aca75ef5de2a4d916accd2dc3316b887a803bff07bcf4cd46281ba1cb6ada25fbf0419b3a9496cc20631ff058f"
	pk, err := ImportPrivateKeyFromHex(privateKey)
	if nil != err {
		t.Fatal(err)
	}

	sender, err := types.AccountFromBytes(pk)
	if err != nil {
		t.Fatal(err)
	}

	receiver := common.PublicKeyFromString("2qy1811C9mgVquHUrPuuzRrxFrp3Y7wh5QCZzVvQbZda")

	// 获取最新区块哈希
	latestBlockhash, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 构造转账指令
	instruction := system.Transfer(system.TransferParam{
		From:   sender.PublicKey,
		To:     receiver,
		Amount: 10000, // 0.00001 SOL
	})

	msg := types.NewMessage(types.NewMessageParam{
		FeePayer:        sender.PublicKey,
		RecentBlockhash: latestBlockhash.Blockhash,
		Instructions:    []types.Instruction{instruction},
	})
	// 查询真实手续费
	feeResp, err := c.GetFeeForMessage(context.Background(), msg)
	if err != nil {
		log.Fatal(err)
	}

	t.Log(*feeResp)

	fee := *feeResp
	t.Log(fmt.Sprintf("%.9f", float64(fee)/float64(1_000_000_000)))
}

func TestSendToken(t *testing.T) {
	c := client.NewClient(clientUrl)

	// ✅ 64 字节私钥（base64 示例，你可以用 JSON 数组或 raw byte）
	privateKey := "341f58e5e65307ea627763a2d730deb6da5b95aca75ef5de2a4d916accd2dc3316b887a803bff07bcf4cd46281ba1cb6ada25fbf0419b3a9496cc20631ff058f"
	pk, err := ImportPrivateKeyFromHex(privateKey)
	if nil != err {
		t.Fatal(err)
	}

	sender, err := types.AccountFromBytes(pk)
	if err != nil {
		t.Fatal(err)
	}

	receiver := common.PublicKeyFromString("2qy1811C9mgVquHUrPuuzRrxFrp3Y7wh5QCZzVvQbZda")

	usdtMint := common.PublicKeyFromString("4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU") // token mint 地址

	// 获取 sender 和 receiver 的 Associated Token Account（ATA）
	fromTokenAccount, nonce, _ := common.FindAssociatedTokenAddress(sender.PublicKey, usdtMint)
	t.Log(nonce)
	toTokenAccount, nonce, _ := common.FindAssociatedTokenAddress(receiver, usdtMint)
	t.Log(nonce)
	t.Log("fromTokenAccount=", base58.Encode(fromTokenAccount[:]))
	t.Log("toTokenAccount=", base58.Encode(toTokenAccount[:]))

	// 获取最新 blockhash
	latestBlockhash, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// 判断 toTokenAccount 是否存在，如果不存在，创建
	toATAInfo, err := c.GetAccountInfo(context.Background(), toTokenAccount.ToBase58())
	instruction := make([]types.Instruction, 0)
	t.Log(toATAInfo.Data)
	if err != nil || len(toATAInfo.Data) == 0 {
		ix := associated_token_account.Create(
			associated_token_account.CreateParam{
				Funder:                 sender.PublicKey,
				Owner:                  receiver,
				Mint:                   usdtMint,
				AssociatedTokenAccount: toTokenAccount,
			},
		)
		instruction = append(instruction, ix)
	}

	// 构造 token 转账指令（单位是最小单位，USDT 是 6 位精度）
	transferIx := token.Transfer(token.TransferParam{
		From:   fromTokenAccount,
		To:     toTokenAccount,
		Auth:   sender.PublicKey,
		Amount: 1_000_000, // 即 1 USDT（6位精度）
	})

	// 构造并发送交易
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        sender.PublicKey,
			RecentBlockhash: latestBlockhash.Blockhash,
			Instructions:    append(instruction, transferIx),
		}),
		Signers: []types.Account{sender},
	})
	if err != nil {
		t.Fatal(err)
	}

	txHashPre := base58.Encode(tx.Signatures[0])
	t.Log(txHashPre)

	txHash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("✅ 交易成功，哈希：", txHash)
}
