///////////////////////////////////////////////////////
// This file was auto-generated by go_bindings.zig   //
//              Do not manually modify.              //
///////////////////////////////////////////////////////

package types

/*
#include "../native/tb_client.h"
*/
import "C"
import "strconv"

type AccountFlags struct {
	Linked                     bool
	DebitsMustNotExceedCredits bool
	CreditsMustNotExceedDebits bool
	History                    bool
	Imported                   bool
}

func (f AccountFlags) ToUint16() uint16 {
	var ret uint16 = 0

	if f.Linked {
		ret |= (1 << 0)
	}

	if f.DebitsMustNotExceedCredits {
		ret |= (1 << 1)
	}

	if f.CreditsMustNotExceedDebits {
		ret |= (1 << 2)
	}

	if f.History {
		ret |= (1 << 3)
	}

	if f.Imported {
		ret |= (1 << 4)
	}

	return ret
}

type TransferFlags struct {
	Linked              bool
	Pending             bool
	PostPendingTransfer bool
	VoidPendingTransfer bool
	BalancingDebit      bool
	BalancingCredit     bool
	Imported            bool
}

func (f TransferFlags) ToUint16() uint16 {
	var ret uint16 = 0

	if f.Linked {
		ret |= (1 << 0)
	}

	if f.Pending {
		ret |= (1 << 1)
	}

	if f.PostPendingTransfer {
		ret |= (1 << 2)
	}

	if f.VoidPendingTransfer {
		ret |= (1 << 3)
	}

	if f.BalancingDebit {
		ret |= (1 << 4)
	}

	if f.BalancingCredit {
		ret |= (1 << 5)
	}

	if f.Imported {
		ret |= (1 << 6)
	}

	return ret
}

type AccountFilterFlags struct {
	Debits   bool
	Credits  bool
	Reversed bool
}

func (f AccountFilterFlags) ToUint32() uint32 {
	var ret uint32 = 0

	if f.Debits {
		ret |= (1 << 0)
	}

	if f.Credits {
		ret |= (1 << 1)
	}

	if f.Reversed {
		ret |= (1 << 2)
	}

	return ret
}

type QueryFilterFlags struct {
	Reversed bool
}

func (f QueryFilterFlags) ToUint32() uint32 {
	var ret uint32 = 0

	if f.Reversed {
		ret |= (1 << 0)
	}

	return ret
}

type Account struct {
	ID             Uint128
	DebitsPending  Uint128
	DebitsPosted   Uint128
	CreditsPending Uint128
	CreditsPosted  Uint128
	UserData128    Uint128
	UserData64     uint64
	UserData32     uint32
	Reserved       uint32
	Ledger         uint32
	Code           uint16
	Flags          uint16
	Timestamp      uint64
}

func (o Account) AccountFlags() AccountFlags {
	var f AccountFlags
	f.Linked = ((o.Flags >> 0) & 0x1) == 1
	f.DebitsMustNotExceedCredits = ((o.Flags >> 1) & 0x1) == 1
	f.CreditsMustNotExceedDebits = ((o.Flags >> 2) & 0x1) == 1
	f.History = ((o.Flags >> 3) & 0x1) == 1
	f.Imported = ((o.Flags >> 4) & 0x1) == 1
	return f
}

type Transfer struct {
	ID              Uint128
	DebitAccountID  Uint128
	CreditAccountID Uint128
	Amount          Uint128
	PendingID       Uint128
	UserData128     Uint128
	UserData64      uint64
	UserData32      uint32
	Timeout         uint32
	Ledger          uint32
	Code            uint16
	Flags           uint16
	Timestamp       uint64
}

func (o Transfer) TransferFlags() TransferFlags {
	var f TransferFlags
	f.Linked = ((o.Flags >> 0) & 0x1) == 1
	f.Pending = ((o.Flags >> 1) & 0x1) == 1
	f.PostPendingTransfer = ((o.Flags >> 2) & 0x1) == 1
	f.VoidPendingTransfer = ((o.Flags >> 3) & 0x1) == 1
	f.BalancingDebit = ((o.Flags >> 4) & 0x1) == 1
	f.BalancingCredit = ((o.Flags >> 5) & 0x1) == 1
	f.Imported = ((o.Flags >> 6) & 0x1) == 1
	return f
}

type CreateAccountResult uint32

const (
	AccountOK                                   CreateAccountResult = 0
	AccountLinkedEventFailed                    CreateAccountResult = 1
	AccountLinkedEventChainOpen                 CreateAccountResult = 2
	AccountImportedEventExpected                CreateAccountResult = 22
	AccountImportedEventNotExpected             CreateAccountResult = 23
	AccountTimestampMustBeZero                  CreateAccountResult = 3
	AccountImportedEventTimestampMustNotBeZero  CreateAccountResult = 24
	AccountImportedEventTimestampMustNotAdvance CreateAccountResult = 25
	AccountReservedField                        CreateAccountResult = 4
	AccountReservedFlag                         CreateAccountResult = 5
	AccountIDMustNotBeZero                      CreateAccountResult = 6
	AccountIDMustNotBeIntMax                    CreateAccountResult = 7
	AccountFlagsAreMutuallyExclusive            CreateAccountResult = 8
	AccountDebitsPendingMustBeZero              CreateAccountResult = 9
	AccountDebitsPostedMustBeZero               CreateAccountResult = 10
	AccountCreditsPendingMustBeZero             CreateAccountResult = 11
	AccountCreditsPostedMustBeZero              CreateAccountResult = 12
	AccountLedgerMustNotBeZero                  CreateAccountResult = 13
	AccountCodeMustNotBeZero                    CreateAccountResult = 14
	AccountExistsWithDifferentFlags             CreateAccountResult = 15
	AccountExistsWithDifferentUserData128       CreateAccountResult = 16
	AccountExistsWithDifferentUserData64        CreateAccountResult = 17
	AccountExistsWithDifferentUserData32        CreateAccountResult = 18
	AccountExistsWithDifferentLedger            CreateAccountResult = 19
	AccountExistsWithDifferentCode              CreateAccountResult = 20
	AccountExists                               CreateAccountResult = 21
	AccountImportedEventTimestampMustNotRegress CreateAccountResult = 26
)

func (i CreateAccountResult) String() string {
	switch i {
	case AccountOK:
		return "AccountOK"
	case AccountLinkedEventFailed:
		return "AccountLinkedEventFailed"
	case AccountLinkedEventChainOpen:
		return "AccountLinkedEventChainOpen"
	case AccountImportedEventExpected:
		return "AccountImportedEventExpected"
	case AccountImportedEventNotExpected:
		return "AccountImportedEventNotExpected"
	case AccountTimestampMustBeZero:
		return "AccountTimestampMustBeZero"
	case AccountImportedEventTimestampMustNotBeZero:
		return "AccountImportedEventTimestampMustNotBeZero"
	case AccountImportedEventTimestampMustNotAdvance:
		return "AccountImportedEventTimestampMustNotAdvance"
	case AccountReservedField:
		return "AccountReservedField"
	case AccountReservedFlag:
		return "AccountReservedFlag"
	case AccountIDMustNotBeZero:
		return "AccountIDMustNotBeZero"
	case AccountIDMustNotBeIntMax:
		return "AccountIDMustNotBeIntMax"
	case AccountFlagsAreMutuallyExclusive:
		return "AccountFlagsAreMutuallyExclusive"
	case AccountDebitsPendingMustBeZero:
		return "AccountDebitsPendingMustBeZero"
	case AccountDebitsPostedMustBeZero:
		return "AccountDebitsPostedMustBeZero"
	case AccountCreditsPendingMustBeZero:
		return "AccountCreditsPendingMustBeZero"
	case AccountCreditsPostedMustBeZero:
		return "AccountCreditsPostedMustBeZero"
	case AccountLedgerMustNotBeZero:
		return "AccountLedgerMustNotBeZero"
	case AccountCodeMustNotBeZero:
		return "AccountCodeMustNotBeZero"
	case AccountExistsWithDifferentFlags:
		return "AccountExistsWithDifferentFlags"
	case AccountExistsWithDifferentUserData128:
		return "AccountExistsWithDifferentUserData128"
	case AccountExistsWithDifferentUserData64:
		return "AccountExistsWithDifferentUserData64"
	case AccountExistsWithDifferentUserData32:
		return "AccountExistsWithDifferentUserData32"
	case AccountExistsWithDifferentLedger:
		return "AccountExistsWithDifferentLedger"
	case AccountExistsWithDifferentCode:
		return "AccountExistsWithDifferentCode"
	case AccountExists:
		return "AccountExists"
	case AccountImportedEventTimestampMustNotRegress:
		return "AccountImportedEventTimestampMustNotRegress"
	}
	return "CreateAccountResult(" + strconv.FormatInt(int64(i+1), 10) + ")"
}

type CreateTransferResult uint32

const (
	TransferOK                                              CreateTransferResult = 0
	TransferLinkedEventFailed                               CreateTransferResult = 1
	TransferLinkedEventChainOpen                            CreateTransferResult = 2
	TransferTimestampMustBeZero                             CreateTransferResult = 3
	TransferReservedFlag                                    CreateTransferResult = 4
	TransferIDMustNotBeZero                                 CreateTransferResult = 5
	TransferIDMustNotBeIntMax                               CreateTransferResult = 6
	TransferFlagsAreMutuallyExclusive                       CreateTransferResult = 7
	TransferDebitAccountIDMustNotBeZero                     CreateTransferResult = 8
	TransferDebitAccountIDMustNotBeIntMax                   CreateTransferResult = 9
	TransferCreditAccountIDMustNotBeZero                    CreateTransferResult = 10
	TransferCreditAccountIDMustNotBeIntMax                  CreateTransferResult = 11
	TransferAccountsMustBeDifferent                         CreateTransferResult = 12
	TransferPendingIDMustBeZero                             CreateTransferResult = 13
	TransferPendingIDMustNotBeZero                          CreateTransferResult = 14
	TransferPendingIDMustNotBeIntMax                        CreateTransferResult = 15
	TransferPendingIDMustBeDifferent                        CreateTransferResult = 16
	TransferTimeoutReservedForPendingTransfer               CreateTransferResult = 17
	TransferAmountMustNotBeZero                             CreateTransferResult = 18
	TransferLedgerMustNotBeZero                             CreateTransferResult = 19
	TransferCodeMustNotBeZero                               CreateTransferResult = 20
	TransferDebitAccountNotFound                            CreateTransferResult = 21
	TransferCreditAccountNotFound                           CreateTransferResult = 22
	TransferAccountsMustHaveTheSameLedger                   CreateTransferResult = 23
	TransferTransferMustHaveTheSameLedgerAsAccounts         CreateTransferResult = 24
	TransferPendingTransferNotFound                         CreateTransferResult = 25
	TransferPendingTransferNotPending                       CreateTransferResult = 26
	TransferPendingTransferHasDifferentDebitAccountID       CreateTransferResult = 27
	TransferPendingTransferHasDifferentCreditAccountID      CreateTransferResult = 28
	TransferPendingTransferHasDifferentLedger               CreateTransferResult = 29
	TransferPendingTransferHasDifferentCode                 CreateTransferResult = 30
	TransferExceedsPendingTransferAmount                    CreateTransferResult = 31
	TransferPendingTransferHasDifferentAmount               CreateTransferResult = 32
	TransferPendingTransferAlreadyPosted                    CreateTransferResult = 33
	TransferPendingTransferAlreadyVoided                    CreateTransferResult = 34
	TransferPendingTransferExpired                          CreateTransferResult = 35
	TransferExistsWithDifferentFlags                        CreateTransferResult = 36
	TransferExistsWithDifferentDebitAccountID               CreateTransferResult = 37
	TransferExistsWithDifferentCreditAccountID              CreateTransferResult = 38
	TransferExistsWithDifferentAmount                       CreateTransferResult = 39
	TransferExistsWithDifferentPendingID                    CreateTransferResult = 40
	TransferExistsWithDifferentUserData128                  CreateTransferResult = 41
	TransferExistsWithDifferentUserData64                   CreateTransferResult = 42
	TransferExistsWithDifferentUserData32                   CreateTransferResult = 43
	TransferExistsWithDifferentTimeout                      CreateTransferResult = 44
	TransferExistsWithDifferentCode                         CreateTransferResult = 45
	TransferExists                                          CreateTransferResult = 46
	TransferOverflowsDebitsPending                          CreateTransferResult = 47
	TransferOverflowsCreditsPending                         CreateTransferResult = 48
	TransferOverflowsDebitsPosted                           CreateTransferResult = 49
	TransferOverflowsCreditsPosted                          CreateTransferResult = 50
	TransferOverflowsDebits                                 CreateTransferResult = 51
	TransferOverflowsCredits                                CreateTransferResult = 52
	TransferOverflowsTimeout                                CreateTransferResult = 53
	TransferExceedsCredits                                  CreateTransferResult = 54
	TransferExceedsDebits                                   CreateTransferResult = 55
	TransferImportedEventExpected                           CreateTransferResult = 56
	TransferImportedEventNotExpected                        CreateTransferResult = 57
	TransferImportedEventTimestampMustNotBeZero             CreateTransferResult = 58
	TransferImportedEventTimestampMustNotAdvance            CreateTransferResult = 59
	TransferImportedEventTimestampMustNotRegress            CreateTransferResult = 60
	TransferImportedEventTimestampMustPostdateDebitAccount  CreateTransferResult = 61
	TransferImportedEventTimestampMustPostdateCreditAccount CreateTransferResult = 62
	TransferImportedEventTimeoutMustBeZero                  CreateTransferResult = 63
)

func (i CreateTransferResult) String() string {
	switch i {
	case TransferOK:
		return "TransferOK"
	case TransferLinkedEventFailed:
		return "TransferLinkedEventFailed"
	case TransferLinkedEventChainOpen:
		return "TransferLinkedEventChainOpen"
	case TransferTimestampMustBeZero:
		return "TransferTimestampMustBeZero"
	case TransferReservedFlag:
		return "TransferReservedFlag"
	case TransferIDMustNotBeZero:
		return "TransferIDMustNotBeZero"
	case TransferIDMustNotBeIntMax:
		return "TransferIDMustNotBeIntMax"
	case TransferFlagsAreMutuallyExclusive:
		return "TransferFlagsAreMutuallyExclusive"
	case TransferDebitAccountIDMustNotBeZero:
		return "TransferDebitAccountIDMustNotBeZero"
	case TransferDebitAccountIDMustNotBeIntMax:
		return "TransferDebitAccountIDMustNotBeIntMax"
	case TransferCreditAccountIDMustNotBeZero:
		return "TransferCreditAccountIDMustNotBeZero"
	case TransferCreditAccountIDMustNotBeIntMax:
		return "TransferCreditAccountIDMustNotBeIntMax"
	case TransferAccountsMustBeDifferent:
		return "TransferAccountsMustBeDifferent"
	case TransferPendingIDMustBeZero:
		return "TransferPendingIDMustBeZero"
	case TransferPendingIDMustNotBeZero:
		return "TransferPendingIDMustNotBeZero"
	case TransferPendingIDMustNotBeIntMax:
		return "TransferPendingIDMustNotBeIntMax"
	case TransferPendingIDMustBeDifferent:
		return "TransferPendingIDMustBeDifferent"
	case TransferTimeoutReservedForPendingTransfer:
		return "TransferTimeoutReservedForPendingTransfer"
	case TransferAmountMustNotBeZero:
		return "TransferAmountMustNotBeZero"
	case TransferLedgerMustNotBeZero:
		return "TransferLedgerMustNotBeZero"
	case TransferCodeMustNotBeZero:
		return "TransferCodeMustNotBeZero"
	case TransferDebitAccountNotFound:
		return "TransferDebitAccountNotFound"
	case TransferCreditAccountNotFound:
		return "TransferCreditAccountNotFound"
	case TransferAccountsMustHaveTheSameLedger:
		return "TransferAccountsMustHaveTheSameLedger"
	case TransferTransferMustHaveTheSameLedgerAsAccounts:
		return "TransferTransferMustHaveTheSameLedgerAsAccounts"
	case TransferPendingTransferNotFound:
		return "TransferPendingTransferNotFound"
	case TransferPendingTransferNotPending:
		return "TransferPendingTransferNotPending"
	case TransferPendingTransferHasDifferentDebitAccountID:
		return "TransferPendingTransferHasDifferentDebitAccountID"
	case TransferPendingTransferHasDifferentCreditAccountID:
		return "TransferPendingTransferHasDifferentCreditAccountID"
	case TransferPendingTransferHasDifferentLedger:
		return "TransferPendingTransferHasDifferentLedger"
	case TransferPendingTransferHasDifferentCode:
		return "TransferPendingTransferHasDifferentCode"
	case TransferExceedsPendingTransferAmount:
		return "TransferExceedsPendingTransferAmount"
	case TransferPendingTransferHasDifferentAmount:
		return "TransferPendingTransferHasDifferentAmount"
	case TransferPendingTransferAlreadyPosted:
		return "TransferPendingTransferAlreadyPosted"
	case TransferPendingTransferAlreadyVoided:
		return "TransferPendingTransferAlreadyVoided"
	case TransferPendingTransferExpired:
		return "TransferPendingTransferExpired"
	case TransferExistsWithDifferentFlags:
		return "TransferExistsWithDifferentFlags"
	case TransferExistsWithDifferentDebitAccountID:
		return "TransferExistsWithDifferentDebitAccountID"
	case TransferExistsWithDifferentCreditAccountID:
		return "TransferExistsWithDifferentCreditAccountID"
	case TransferExistsWithDifferentAmount:
		return "TransferExistsWithDifferentAmount"
	case TransferExistsWithDifferentPendingID:
		return "TransferExistsWithDifferentPendingID"
	case TransferExistsWithDifferentUserData128:
		return "TransferExistsWithDifferentUserData128"
	case TransferExistsWithDifferentUserData64:
		return "TransferExistsWithDifferentUserData64"
	case TransferExistsWithDifferentUserData32:
		return "TransferExistsWithDifferentUserData32"
	case TransferExistsWithDifferentTimeout:
		return "TransferExistsWithDifferentTimeout"
	case TransferExistsWithDifferentCode:
		return "TransferExistsWithDifferentCode"
	case TransferExists:
		return "TransferExists"
	case TransferOverflowsDebitsPending:
		return "TransferOverflowsDebitsPending"
	case TransferOverflowsCreditsPending:
		return "TransferOverflowsCreditsPending"
	case TransferOverflowsDebitsPosted:
		return "TransferOverflowsDebitsPosted"
	case TransferOverflowsCreditsPosted:
		return "TransferOverflowsCreditsPosted"
	case TransferOverflowsDebits:
		return "TransferOverflowsDebits"
	case TransferOverflowsCredits:
		return "TransferOverflowsCredits"
	case TransferOverflowsTimeout:
		return "TransferOverflowsTimeout"
	case TransferExceedsCredits:
		return "TransferExceedsCredits"
	case TransferExceedsDebits:
		return "TransferExceedsDebits"
	case TransferImportedEventExpected:
		return "TransferImportedEventExpected"
	case TransferImportedEventNotExpected:
		return "TransferImportedEventNotExpected"
	case TransferImportedEventTimestampMustNotBeZero:
		return "TransferImportedEventTimestampMustNotBeZero"
	case TransferImportedEventTimestampMustNotAdvance:
		return "TransferImportedEventTimestampMustNotAdvance"
	case TransferImportedEventTimestampMustNotRegress:
		return "TransferImportedEventTimestampMustNotRegress"
	case TransferImportedEventTimestampMustPostdateDebitAccount:
		return "TransferImportedEventTimestampMustPostdateDebitAccount"
	case TransferImportedEventTimestampMustPostdateCreditAccount:
		return "TransferImportedEventTimestampMustPostdateCreditAccount"
	case TransferImportedEventTimeoutMustBeZero:
		return "TransferImportedEventTimeoutMustBeZero"
	}
	return "CreateTransferResult(" + strconv.FormatInt(int64(i+1), 10) + ")"
}

type AccountEventResult struct {
	Index  uint32
	Result CreateAccountResult
}

type TransferEventResult struct {
	Index  uint32
	Result CreateTransferResult
}

type AccountFilter struct {
	AccountID    Uint128
	TimestampMin uint64
	TimestampMax uint64
	Limit        uint32
	Flags        uint32
	Reserved     [24]uint8
}

func (o AccountFilter) AccountFilterFlags() AccountFilterFlags {
	var f AccountFilterFlags
	f.Debits = ((o.Flags >> 0) & 0x1) == 1
	f.Credits = ((o.Flags >> 1) & 0x1) == 1
	f.Reversed = ((o.Flags >> 2) & 0x1) == 1
	return f
}

type AccountBalance struct {
	DebitsPending  Uint128
	DebitsPosted   Uint128
	CreditsPending Uint128
	CreditsPosted  Uint128
	Timestamp      uint64
	Reserved       [56]uint8
}

type QueryFilter struct {
	UserData128  Uint128
	UserData64   uint64
	UserData32   uint32
	Ledger       uint32
	Code         uint16
	Reserved     [6]uint8
	TimestampMin uint64
	TimestampMax uint64
	Limit        uint32
	Flags        uint32
}

func (o QueryFilter) QueryFilterFlags() QueryFilterFlags {
	var f QueryFilterFlags
	f.Reversed = ((o.Flags >> 0) & 0x1) == 1
	return f
}

