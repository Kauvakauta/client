// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/gpg_ui.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type SelectKeyRes struct {
	KeyID        string `codec:"keyID" json:"keyID"`
	DoSecretPush bool   `codec:"doSecretPush" json:"doSecretPush"`
}

func (o SelectKeyRes) DeepCopy() SelectKeyRes {
	return SelectKeyRes{
		KeyID:        o.KeyID,
		DoSecretPush: o.DoSecretPush,
	}
}

type WantToAddGPGKeyArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o WantToAddGPGKeyArg) DeepCopy() WantToAddGPGKeyArg {
	return WantToAddGPGKeyArg{
		SessionID: o.SessionID,
	}
}

type ConfirmDuplicateKeyChosenArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o ConfirmDuplicateKeyChosenArg) DeepCopy() ConfirmDuplicateKeyChosenArg {
	return ConfirmDuplicateKeyChosenArg{
		SessionID: o.SessionID,
	}
}

type SelectKeyAndPushOptionArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Keys      []GPGKey `codec:"keys" json:"keys"`
}

func (o SelectKeyAndPushOptionArg) DeepCopy() SelectKeyAndPushOptionArg {
	return SelectKeyAndPushOptionArg{
		SessionID: o.SessionID,
		Keys: (func(x []GPGKey) []GPGKey {
			var ret []GPGKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Keys),
	}
}

type SelectKeyArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Keys      []GPGKey `codec:"keys" json:"keys"`
}

func (o SelectKeyArg) DeepCopy() SelectKeyArg {
	return SelectKeyArg{
		SessionID: o.SessionID,
		Keys: (func(x []GPGKey) []GPGKey {
			var ret []GPGKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Keys),
	}
}

type SignArg struct {
	Msg         []byte `codec:"msg" json:"msg"`
	Fingerprint []byte `codec:"fingerprint" json:"fingerprint"`
}

func (o SignArg) DeepCopy() SignArg {
	return SignArg{
		Msg: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.Msg),
		Fingerprint: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.Fingerprint),
	}
}

type GetTTYArg struct {
}

func (o GetTTYArg) DeepCopy() GetTTYArg {
	return GetTTYArg{}
}

type GpgUiInterface interface {
	WantToAddGPGKey(context.Context, int) (bool, error)
	ConfirmDuplicateKeyChosen(context.Context, int) (bool, error)
	SelectKeyAndPushOption(context.Context, SelectKeyAndPushOptionArg) (SelectKeyRes, error)
	SelectKey(context.Context, SelectKeyArg) (string, error)
	Sign(context.Context, SignArg) (string, error)
	GetTTY(context.Context) (string, error)
}

func GpgUiProtocol(i GpgUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.gpgUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"wantToAddGPGKey": {
				MakeArg: func() interface{} {
					ret := make([]WantToAddGPGKeyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]WantToAddGPGKeyArg)
					if !ok {
						err = rpc.NewTypeError((*[]WantToAddGPGKeyArg)(nil), args)
						return
					}
					ret, err = i.WantToAddGPGKey(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"confirmDuplicateKeyChosen": {
				MakeArg: func() interface{} {
					ret := make([]ConfirmDuplicateKeyChosenArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ConfirmDuplicateKeyChosenArg)
					if !ok {
						err = rpc.NewTypeError((*[]ConfirmDuplicateKeyChosenArg)(nil), args)
						return
					}
					ret, err = i.ConfirmDuplicateKeyChosen(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"selectKeyAndPushOption": {
				MakeArg: func() interface{} {
					ret := make([]SelectKeyAndPushOptionArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SelectKeyAndPushOptionArg)
					if !ok {
						err = rpc.NewTypeError((*[]SelectKeyAndPushOptionArg)(nil), args)
						return
					}
					ret, err = i.SelectKeyAndPushOption(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"selectKey": {
				MakeArg: func() interface{} {
					ret := make([]SelectKeyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SelectKeyArg)
					if !ok {
						err = rpc.NewTypeError((*[]SelectKeyArg)(nil), args)
						return
					}
					ret, err = i.SelectKey(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"sign": {
				MakeArg: func() interface{} {
					ret := make([]SignArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SignArg)
					if !ok {
						err = rpc.NewTypeError((*[]SignArg)(nil), args)
						return
					}
					ret, err = i.Sign(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getTTY": {
				MakeArg: func() interface{} {
					ret := make([]GetTTYArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetTTY(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type GpgUiClient struct {
	Cli rpc.GenericClient
}

func (c GpgUiClient) WantToAddGPGKey(ctx context.Context, sessionID int) (res bool, err error) {
	__arg := WantToAddGPGKeyArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.gpgUi.wantToAddGPGKey", []interface{}{__arg}, &res)
	return
}

func (c GpgUiClient) ConfirmDuplicateKeyChosen(ctx context.Context, sessionID int) (res bool, err error) {
	__arg := ConfirmDuplicateKeyChosenArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.gpgUi.confirmDuplicateKeyChosen", []interface{}{__arg}, &res)
	return
}

func (c GpgUiClient) SelectKeyAndPushOption(ctx context.Context, __arg SelectKeyAndPushOptionArg) (res SelectKeyRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.gpgUi.selectKeyAndPushOption", []interface{}{__arg}, &res)
	return
}

func (c GpgUiClient) SelectKey(ctx context.Context, __arg SelectKeyArg) (res string, err error) {
	err = c.Cli.Call(ctx, "keybase.1.gpgUi.selectKey", []interface{}{__arg}, &res)
	return
}

func (c GpgUiClient) Sign(ctx context.Context, __arg SignArg) (res string, err error) {
	err = c.Cli.Call(ctx, "keybase.1.gpgUi.sign", []interface{}{__arg}, &res)
	return
}

func (c GpgUiClient) GetTTY(ctx context.Context) (res string, err error) {
	err = c.Cli.Call(ctx, "keybase.1.gpgUi.getTTY", []interface{}{GetTTYArg{}}, &res)
	return
}
