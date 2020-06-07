package main

import (
	"errors"
	"log"
)

type WithdrawCommand interface {
	Execute() error
	Undo() error
	Redo() error
}

type CommandManager struct {
	UndoCommand []WithdrawCommand
	RedoCommand []WithdrawCommand
}

func NewCommandManager() CommandManager{
	return CommandManager{
		UndoCommand:make([]WithdrawCommand,0),
		RedoCommand:make([]WithdrawCommand,0),
	}
}

func (cmdManager *CommandManager)ExecuteCmd(cmd WithdrawCommand)error{
	if err:=cmd.Execute();err!=nil{
		return err
	}
	cmdManager.UndoCommand = append(cmdManager.UndoCommand,cmd)
	if len(cmdManager.RedoCommand)!=0{
		cmdManager.clearRedoCmd()
	}
	return nil
}

func (cmdManager *CommandManager)Undo()error{
	if err:=cmdManager.UndoCommand[len(cmdManager.UndoCommand)-1].Undo();err!=nil{
		return err
	}
	cmdManager.UndoCommand = cmdManager.UndoCommand[:len(cmdManager.UndoCommand)-1]
	return nil
}

func (cmdmanager *CommandManager)UndoAll()error{
	for !cmdmanager.isUndoCmdEmpty(){
		err:=cmdmanager.Undo()
		if err!=nil{
			return err
		}
	}
	return nil
}

func (cmdManager *CommandManager)Redo()error{
	return nil
}
func (cmdManager *CommandManager)isUndoCmdEmpty()bool{
	return len(cmdManager.UndoCommand) == 0
}
func (cmdManager *CommandManager)clearRedoCmd(){
	cmdManager.RedoCommand = cmdManager.RedoCommand[0:0]
}

type WithdrawArgs struct {
	WithdrawAmount int64
	TraeNo string
	Uid int64
}

type AddWithdrawRecordCmd struct {
	Args WithdrawArgs
	WantErr bool
}

func (command *AddWithdrawRecordCmd)Execute()error{
	command.Args.TraeNo = "123456"
	log.Printf("添加提现记录,用户ID：%d,提现金额:%d 交易单号:%v\n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	if command.WantErr{
		return errors.New("模拟[添加提现记录]错误\n")
	}
	return nil
}

func (command *AddWithdrawRecordCmd)Undo()error{
	log.Printf("回滚提现记录,用户ID：%v,提现金额:%d,交易单号:%v\n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	return nil
}

func (command *AddWithdrawRecordCmd)Redo()error{
	log.Printf("redo 提现记录,用户ID：%v，提现金额：%v，交易单号：%v\n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	return nil
}

type DeductAmount struct {
	Args WithdrawArgs
	WantErr bool
}

func (command *DeductAmount)Execute()error{
	log.Printf("扣款，用户ID：%v 扣款金额:%v 交易单号：%v\n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	if command.WantErr{
		return errors.New("模拟扣款失败\n")
	}
	return nil
}

func(command *DeductAmount)Undo()error{
	log.Printf("回滚扣款，用户ID：%v,扣款金额：%v,交易单号：%v\n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	return nil
}

func (command *DeductAmount)Redo()error{
	log.Printf("redo 扣款,用户ID：%v,提现金额：%v,交易单号:%v\n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	return nil
}

type DoWithdrawCommand struct {
	Args WithdrawArgs
	WantErr bool
}

func(command *DoWithdrawCommand)Execute()error{
	log.Printf("微信大钱，用户ID：%v，金额:%v ,交易单号:%v \n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	if command.WantErr{
		return errors.New("模拟扣款失败")
	}
	return nil
}

func (command *DoWithdrawCommand)Undo()error{
	log.Printf("undo 微信大钱,用户ID：%v 金额：%v,交易单号:%v\n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	return nil
}

func (command *DoWithdrawCommand)Redo()error{
	log.Printf("redo 微信打钱,用户ID：%v，金额：%v ,交易单号:%v\n",
		command.Args.Uid,command.Args.WithdrawAmount,command.Args.TraeNo)
	return nil
}

func main(){
	tests:=[]struct{
		Name string
		wantAddrecordErr bool
		wantDeductErr bool
		wantDoWithdraw bool
		args WithdrawArgs
	}{
		{
			"正常提现",
			false,
			false,
			false,
			WithdrawArgs{
				WithdrawAmount:100,
				Uid:1,
			},
		},{
			"添加提现记录失败",
			true,
			false,
			false,
			WithdrawArgs{
				WithdrawAmount:100,
				Uid:1,
			},
		},
		{
			"扣款失败",
			false,
			true,
			false,
			WithdrawArgs{
				WithdrawAmount:100,
				Uid:1,
			},
		},{
			"提现打款失败",
			false,
			false,
			true,
			WithdrawArgs{
				WithdrawAmount:100,
				Uid:1,
			},
		},
	}
	for _,test:=range tests{
		tt:=test
		cmdManager:=NewCommandManager()
		var err error
		//添加记录
		addWithdrawRecordCmd:=&AddWithdrawRecordCmd{
			Args:tt.args,
			WantErr:tt.wantAddrecordErr,
		}
		err=cmdManager.ExecuteCmd(addWithdrawRecordCmd)
		tt.args.TraeNo = addWithdrawRecordCmd.Args.TraeNo

		if tt.wantAddrecordErr&&err!=nil{
			_=cmdManager.UndoAll()
			return
		}
		//扣款
		err=cmdManager.ExecuteCmd(&DeductAmount{
			Args:tt.args,
			WantErr:tt.wantDeductErr,
		})
		if tt.wantDeductErr&&err!=nil{
			_=cmdManager.UndoAll()
			return
		}

		//打钱
		err=cmdManager.ExecuteCmd(&DoWithdrawCommand{
			Args:tt.args,
			WantErr:tt.wantDoWithdraw,
		})
		if tt.wantDoWithdraw&&err!=nil{
			_=cmdManager.UndoAll()
			return
		}
	}
}
